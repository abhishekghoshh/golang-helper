package main

import (
	"grpc-basics/proto"
	"grpc-basics/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	proto.RegisterGreetServiceServer(grpcServer, &server.GreetController{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
