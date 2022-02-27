/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-12 01:07:31
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-24 00:07:11
 * @FilePath: /grpc_demo/cmd/client/main.go
 */

package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"pcbook/client"
	"pcbook/pb"
	"pcbook/sample"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func testCreatelaptop(laptopClient client.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

func testRateLaptop(laptopClient client.LaptopClient) {
	n := 3
	laptopIDs := make([]string, n)
	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, n)

	for {
		fmt.Print("rating laptop [y/n]")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}
		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()

		}
		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func testUploadLaptopImage(laptopClient client.LaptopClient) {
	laotop := sample.NewLaptop()
	laptopClient.CreateLaptop(laotop)
	laptopClient.UploadImage(laotop.GetId(), "tmp/back.jpeg")
}

func testSearchLaptop(laptopClient client.LaptopClient) {

	i := 0
	for i < 10 {
		laptopClient.CreateLaptop(sample.NewLaptop())
		i++
	}
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Unit: pb.Memory_GIGABYTE, Value: 8},
	}
	laptopClient.SearchLaptop(filter)
}

const (
	username        = "zplus"
	password        = "12345"
	refreshDuration = time.Second * 30
)

func authMethods() map[string]bool {
	const laptopServicePath = "/zplus.pcbook.LaptopService/"
	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA certification")
	}

	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	fmt.Println("Hello World from Client")
	address := flag.String("address", "0.0.0.0:9999", "the server address")
	enableTLS := flag.Bool("tls", false, "enable tls/ssl")

	log.Printf("dial server on address = %s , TLS = %v", *address, *enableTLS)
	flag.Parse()

	transportOption := grpc.WithInsecure()
	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal(err)
		}
		transportOption = grpc.WithTransportCredentials(tlsCredentials)

	}

	conn1, err := grpc.Dial(*address, transportOption)
	if err != nil {
		log.Fatal("cannot conn the server")
	}

	authClient := client.NewAuthClient(conn1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)

	if err != nil {
		log.Fatal("canont dial the auth server: ", err)
	}

	conn2, err := grpc.Dial(
		*address,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot conn the server")
	}
	laptopClient := client.NewLaptopClient(conn2)
	//laptop := sample.NewLaptop()
	// createLaptop(laptopClient, laptop)
	//testUploadLaptopImage(laptopClient, laptop.Id, "tmp/back.jpeg")
	testRateLaptop(*laptopClient)
}
