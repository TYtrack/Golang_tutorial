/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-15 13:00:45
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-15 13:16:00
 * @FilePath: /grpc_demo/client/laptop_client.go
 */

package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"pcbook/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopClient struct {
	service pb.LaptopServiceClient
}

func NewLaptopClient(conn *grpc.ClientConn) *LaptopClient {
	service := pb.NewLaptopServiceClient(conn)
	return &LaptopClient{
		service: service,
	}
}

func (laptopClient *LaptopClient) CreateLaptop(laptop *pb.Laptop) {

	req := pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	//设置超时取消请求
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response, err := laptopClient.service.CreateLaptop(ctx, &req)
	if err != nil {
		sta, ok := status.FromError(err)
		if ok && sta.Code() == codes.AlreadyExists {
			log.Println("The id already exists")
		} else {
			log.Fatalln("server internel error")
		}
		return

	}
	log.Printf("created the laptop with id [%v]", response.Id)

}

func (laptopClient *LaptopClient) SearchLaptop(filter *pb.Filter) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.service.SearchLaptop(ctx, req)

	if err != nil {
		log.Fatalf("cannot start the search :%v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("cannot receive the response : %v", res)
		}
		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Print("    + brand: ", laptop.GetBrand())
		log.Print("    + name: ", laptop.GetName())
		log.Print("    + cpu_cores: ", laptop.GetCpu().GetNumberCores())
		log.Print("    + cpu_min_ghz: ", laptop.GetCpu().GetMinGhz())
		log.Print("    + ram : ", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
	}

}

func (laptopClient *LaptopClient) UploadImage(laotopID string, imagePath string) error {

	file, err := os.Open(imagePath)

	if err != nil {
		log.Fatal("cannot open image file", err)
	}

	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := laptopClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image ilfe", err)
	}

	req := pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laotopID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(&req)
	if err != nil {
		log.Fatal("cannot send image ilfe", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer:", err)
		}
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server:", err, stream.RecvMsg(nil))

		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response:", err)
	}

	log.Printf("image uploaded with id :%s ,size%v", res.GetId(), res.GetSize())
	return nil

}

func (laptopClient *LaptopClient) RateLaptop(laotopIDs []string, scores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	stream, err := laptopClient.service.RateLaptop(ctx)
	if err != nil {
		return err
	}
	waitRespnse := make(chan error)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("no more responses")
				waitRespnse <- nil
				return
			}
			if err != nil {
				waitRespnse <- fmt.Errorf("cannot receive stream response:%v ", err)
				return
			}

			log.Print("receive response :", res)
		}
	}()

	for i, laptopID := range laotopIDs {
		req := &pb.RateLaptopRequest{
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request :%v-%v", err, stream.RecvMsg(nil))
		}
		log.Print("sent request :", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send stream")
	}
	err = <-waitRespnse
	return err
}
