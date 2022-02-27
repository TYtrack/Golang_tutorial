/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-14 21:32:03
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-14 21:36:45
 * @FilePath: /grpc_demo/client/auth_client.go
 */

package client

import (
	"context"
	"pcbook/pb"
	"time"

	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	return &AuthClient{
		service:  pb.NewAuthServiceClient(cc),
		username: username,
		password: password,
	}
}

func (authClient *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.LoginRequest{
		Username: authClient.username,
		Password: authClient.password,
	}

	res, err := authClient.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil

}
