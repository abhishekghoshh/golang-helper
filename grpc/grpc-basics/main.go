package main

import (
	"grpc-basics/proto"
	"grpc-basics/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const addr = "0.0.0.0:50051"
const clientAddress = "localhost:50051"
const tls = true
const interceptorEnbaled = true

func addServerInterceptors(opts []grpc.ServerOption) []grpc.ServerOption {
	if interceptorEnbaled {
		opts = append(opts, grpc.ChainUnaryInterceptor(server.LogInterceptor()))
		opts = append(opts, grpc.ChainUnaryInterceptor(server.AuthorizationHeaderInterceptor()))
	}
	return opts
}
func serverCredOptions(opts []grpc.ServerOption) []grpc.ServerOption {
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	return opts
}

func main() {
	// first we create a tcp listener
	tcpListener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer tcpListener.Close()
	log.Printf("Listening at %s\n", addr)

	// then we create a gRPC server
	opts := []grpc.ServerOption{}
	opts = serverCredOptions(opts)
	opts = addServerInterceptors(opts)

	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.Stop()

	// then we register our grpc controller with the grpc server
	proto.RegisterGreetServiceServer(grpcServer, &server.GreetController{})

	// this is only for the application running in dev
	reflection.Register(grpcServer)

	// now we attach the grpc server to that tcp listener
	if err := grpcServer.Serve(tcpListener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
