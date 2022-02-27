/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-14 12:28:02
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-14 12:42:26
 * @FilePath: /grpc_demo/service/auth_server.go
 */
package service

import (
	"context"
	"pcbook/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

func (authServer *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := authServer.userStore.Find(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot access the userStore")
	}
	if user == nil || user.IsCorrectPassword(req.Password) {
		return nil, status.Errorf(codes.NotFound, "username doen't exists / password incorrect")
	}
	accessToken, err := authServer.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate the Token")
	}

	res := &pb.LoginResponse{
		AccessToken: accessToken,
	}

	return res, nil
}
