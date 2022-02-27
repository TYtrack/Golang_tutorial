/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-12 01:07:31
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-24 00:04:41
 * @FilePath: /grpc_demo/cmd/server/main.go
 */

package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey     = "secret"
	tokenDuration = time.Minute * 15
)

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "zplus", "123456", "admin")
	if err != nil {
		return err
	}

	err = createUser(userStore, "ctx", "123456", "user")
	if err != nil {
		return err
	}
	return nil
}

func createUser(userStore service.UserStore, username string, password string, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return errors.New("cannot create new user")
	}
	return userStore.Save(user)
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/zplus.pcbook.LaptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemClientCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add server CA certification")
	}

	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	return credentials.NewTLS(config), nil
}

func main() {
	port := flag.Int("port", 0, "the port")
	enableTLS := flag.Bool("tls", false, "enable tls/ssl")

	flag.Parse()
	log.Printf("start server on port = %d , TLS = %v", *port, *enableTLS)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot execute the seedUser")
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)
	laptop_server := service.NewLaotopServer(laptopStore, imageStore, ratingStore)

	authInterceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOption := []grpc.ServerOption{
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	}

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal(err)
		}
		serverOption = append(serverOption, grpc.Creds(tlsCredentials))

	}
	grpc_server := grpc.NewServer(serverOption...)

	pb.RegisterAuthServiceServer(grpc_server, authServer)
	pb.RegisterLaptopServiceServer(grpc_server, laptop_server)

	reflection.Register(grpc_server)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot start the server")
	}

	err = grpc_server.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start the server")
	}
}
