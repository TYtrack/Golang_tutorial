gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:./

clean:
	rm pb/*

server1:
	go run cmd/server/main.go -port 50051 

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server:
	go run cmd/server/main.go -port 9999 

client:
	go run cmd/client/main.go -address 0.0.0.0:9999 

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:9999 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh ; cd ..


.PHONY: gen server clean test client cert client-tls