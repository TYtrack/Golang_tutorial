/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-14 13:22:05
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-14 21:27:49
 * @FilePath: /grpc_demo/service/auth_interceptor.go
 */

package service

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		log.Println("---> unary interceptor", info.FullMethod)

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		log.Println("---> stream interceptor", info.FullMethod)
		return handler(srv, ss)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		return nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]

	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "authorization token is invalid")
	}
	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}

	}
	return status.Errorf(codes.PermissionDenied, "no permission to access")

}