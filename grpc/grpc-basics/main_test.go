package main

import (
	"context"
	"grpc-basics/proto"
	"io"
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

func Test_helloUnary(t *testing.T) {
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

func Test_helloServerStreaming(t *testing.T) {
	greetResponses, err := grpcConnection(func(conn *grpc.ClientConn) ([]*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		// now we invoke the greet method of the greet service
		serverStreamingClient, err := greetServiceClient.GreetServerStreaming(
			context.Background(),
			&proto.GreetRequest{
				Name: "world",
			},
		)
		if nil != err {
			return nil, err
		}
		responses := []*proto.GreetResponse{}
		for {
			mssg, err := serverStreamingClient.Recv()
			if io.EOF == err {
				break
			} else if nil != err {
				return nil, err
			}
			responses = append(responses, mssg)
		}
		return responses, nil
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greetings length: %d\n", len(greetResponses))
	log.Printf("Greetings are: %s\n", greetResponses)
}

func Test_helloClientStreaming(t *testing.T) {
	greetResponse, err := grpcConnection(func(conn *grpc.ClientConn) (*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		// now we invoke the greet method of the greet service
		clientStreaming, err := greetServiceClient.GrretClientStreaming(context.Background())
		if nil != err {
			return nil, err
		}
		allNames := []string{
			"abhishek",
			"nasim",
			"sayan",
		}
		for _, name := range allNames {
			clientStreaming.Send(
				&proto.GreetRequest{
					Name: name,
				})
		}
		return clientStreaming.CloseAndRecv()
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", greetResponse)
}

func Test_helloBidirectionalStreaming(t *testing.T) {
	greetResponses, err := grpcConnection(func(conn *grpc.ClientConn) ([]*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		allNames := []string{
			"abhishek",
			"nasim",
			"sayan",
		}

		// now we invoke the greet method of the greet service
		bidirectionalClient, err := greetServiceClient.GreetBidirectionalStreaming(
			context.Background(),
		)
		if nil != err {
			return nil, err
		}
		responses := []*proto.GreetResponse{}
		for _, name := range allNames {
			err = bidirectionalClient.Send(
				&proto.GreetRequest{
					Name: name,
				},
			)
			if nil != err {
				return nil, err
			}
			mssg, err := bidirectionalClient.Recv()
			if io.EOF == err {
				break
			} else if nil != err {
				return nil, err
			}
			responses = append(responses, mssg)
		}
		return responses, nil
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greetings length: %d\n", len(greetResponses))
	log.Printf("Greetings are: %s\n", greetResponses)
}
