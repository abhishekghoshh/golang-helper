package main

import (
	"context"
	"grpc-basics/proto"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func grpcConnection[T any](runnable func(*grpc.ClientConn) (T, error)) (T, error) {
	// these are the grpc options, we can add as many as we want
	opts := []grpc.DialOption{}
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, creds)

	// this is grpc client
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	return runnable(conn)
}

func Test_hello_unary(t *testing.T) {
	greetResponse, err := grpcConnection(func(conn *grpc.ClientConn) (*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		// now we invoke the greet method of the greet service
		return greetServiceClient.Greet(
			context.Background(),
			&proto.GreetRequest{
				Name: "world",
			},
		)
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", greetResponse)
}
