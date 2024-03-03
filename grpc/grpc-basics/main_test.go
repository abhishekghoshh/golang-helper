package main

import (
	"context"
	"grpc-basics/proto"
	"io"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func addHeaderInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		send, _ := metadata.FromOutgoingContext(ctx)
		newMD := metadata.Pairs("authorization", "aDummyToken")
		ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
func logInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("%s was invoked with %v\n", method, req)
		headers, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			log.Printf("Sending headers: %v\n", headers)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
func addClientInterceptors(clientOpts []grpc.DialOption) []grpc.DialOption {
	if interceptorEnbaled {
		clientOpts = append(clientOpts, grpc.WithChainUnaryInterceptor(addHeaderInterceptor()))
		clientOpts = append(clientOpts, grpc.WithChainUnaryInterceptor(logInterceptor()))
	}
	return clientOpts
}

func clientCredOptions(opts []grpc.DialOption) []grpc.DialOption {
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}
	return opts
}
func grpcConnection[T any](runnable func(*grpc.ClientConn) (T, error)) (T, error) {
	opts := []grpc.DialOption{}
	opts = clientCredOptions(opts)
	opts = addClientInterceptors(opts)
	// this is grpc client
	conn, err := grpc.Dial(clientAddress, opts...)
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

func Test_helloUnaryWithError(t *testing.T) {
	_, err := grpcConnection(func(conn *grpc.ClientConn) (*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		// now we invoke the greet method of the greet service
		return greetServiceClient.Greet(
			context.Background(),
			&proto.GreetRequest{
				Name: "    ",
			},
		)
	})
	if err != nil {
		grpcError, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("Error message from server: %v\n", grpcError.Message())
			log.Println(grpcError.Code())
			log.Printf("Is invalid argument error: %t", (codes.InvalidArgument == grpcError.Code()))
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	} else {
		log.Fatalln("Expecting an error here")
	}
}

func Test_helloUnaryWithTimeout(t *testing.T) {

	timeout := 2 * time.Second

	_, err := grpcConnection(func(conn *grpc.ClientConn) (*proto.GreetResponse, error) {
		// then we create a greet service client
		greetServiceClient := proto.NewGreetServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// now we invoke the greet method of the greet service
		return greetServiceClient.GreetWithTimeout(
			ctx,
			&proto.GreetRequest{
				Name: "Name",
			},
		)
	})
	if err != nil {
		grpcError, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("Error message from server: %v\n", grpcError.Message())
			log.Println(grpcError.Code())
			log.Printf("Is Deadline exceeded error: %t", (codes.DeadlineExceeded == grpcError.Code()))
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	} else {
		log.Fatalln("Expecting an error here")
	}
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
		serverStreamingClient.CloseSend()
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
		// we can send and receive in different channels
		// but here for the simplicity we are doing this synchronusly
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
		// we will close the client
		bidirectionalClient.CloseSend()
		return responses, nil
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greetings length: %d\n", len(greetResponses))
	log.Printf("Greetings are: %s\n", greetResponses)
}
