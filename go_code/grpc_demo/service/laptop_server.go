/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-11 13:57:57
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-13 16:06:53
 * @FilePath: /grpc_demo/service/laptop_server.go
 */

package service

import (
	"bytes"
	"context"
	"io"
	"log"
	"pcbook/pb"

	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxImageSize = 1 << 20

type LaptopServer struct {
	laptopStore LaptopStore
	imageStore  ImageStore
	rateStore   RatingStore
}

func NewLaotopServer(laptopStore LaptopStore, imageStore ImageStore, ratingStore RatingStore) *LaptopServer {
	return &LaptopServer{laptopStore, imageStore, ratingStore}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context,
	req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()
	log.Printf("the sample laptop id is %v\n", laptop.Id)

	if len(laptop.Id) > 0 {
		uuid_1, err := uuid.ParseUUID(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop uuid is not correct: %v\n", uuid_1)
		}
	} else {
		id, err := uuid.GenerateUUID()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new uuid\n")
		}
		laptop.Id = id
	}

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	if ctx.Err() == context.Canceled {
		log.Println("the request canceled")
		return nil, status.Error(codes.Canceled, "the request canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("the request timeout")
		return nil, status.Error(codes.DeadlineExceeded, "the request timeout")
	}

	err := server.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if err == ErrAlreadyExists {
			code = codes.AlreadyExists

		}
		return nil, status.Errorf(code, "cannot save the laptop[%v] tp the store", laptop)
	}

	log.Printf("saved the laptop with id:[%v]\n", laptop.Id)

	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}

func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer) error {

	filter := req.GetFilter()
	log.Printf("receive a search-laptop request for %v\n", filter)

	err := server.laptopStore.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{
				Laptop: laptop,
			}
			//发送流消息
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("send a laptop response for %v\n", laptop.Id)
			return nil
		},
	)
	if err != nil {
		return status.Error(codes.Internal, "unexpected error")
	}

	return nil
}

func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()

	log.Printf("receive an upload-image request for laptop %s with iamge type %s", laptopID, imageType)

	laptop, err := server.laptopStore.Find(laptopID)

	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop:%v", laptopID))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.Internal, "doen't exists the laptop:%v", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		if err := contextError(stream.Context()); err != nil {
			return nil
		}

		log.Println("waiting to receive more date")
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot receive chunk data :%v", err))
		}
		chunk := req.GetChunkData()
		size := len(chunk)

		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is to large "))

		}

		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data to the store :%v", err))
		}

	}

	imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save data to the store :%v", err))
	}

	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response :%v", err))
	}
	log.Printf("saved image with is:%s,size: %d", imageID, imageSize)

	return nil
}

func (server *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return logError(status.Error(codes.Unknown, "cannot receive a rating request"))
		}

		laptopID := req.GetLaptopId()
		score := req.GetScore()
		log.Printf("received a rate-laptop request: id=%s, score=%.2f", laptopID, score)

		found, err := server.laptopStore.Find(laptopID)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot find laptop :%v", err))
		}
		if found == nil {
			return logError(status.Errorf(codes.Internal, "laptopID %s is not found", laptopID))
		}

		rating, err := server.rateStore.Add(laptopID, score)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot add rating to the score :%v", err))
		}
		res := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}

		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response:%v", err))
		}

	}
	return nil
}

func logError(err error) error {
	log.Println(err)
	return err
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		log.Println("the request canceled")
		return status.Error(codes.Canceled, "the request canceled")
	case context.DeadlineExceeded:
		log.Printf("the request timeout")
		return status.Error(codes.DeadlineExceeded, "the request timeout")
	default:
		return nil
	}

}
