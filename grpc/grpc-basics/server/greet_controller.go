package server

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	pb "grpc-basics/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreetController struct {
	pb.GreetServiceServer
}

func (*GreetController) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	// we can also handle error and return to the client
	if strings.TrimSpace(req.Name) == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Empty field : name",
		)
	}
	return Greet(ctx, req)
}

func (*GreetController) GreetServerStreaming(req *pb.GreetRequest, streamingServer pb.GreetService_GreetServerStreamingServer) error {
	for i := 0; i < 10; i++ {
		streamingServer.Send(
			NewGreetResponse(req.Name),
		)
	}
	return nil
}

func (*GreetController) GrretClientStreaming(clientStreaming pb.GreetService_GrretClientStreamingServer) error {
	var allNames strings.Builder
	for {
		mssg, err := clientStreaming.Recv()

		if io.EOF == err {
			return clientStreaming.SendAndClose(
				NewGreetResponse(allNames.String()),
			)
		} else if nil != err {
			return err
		}
		allNames.WriteString(mssg.Name + " ")
	}
}

func (*GreetController) GreetBidirectionalStreaming(bidirectionalStreaming pb.GreetService_GreetBidirectionalStreamingServer) error {
	for {
		mssg, err := bidirectionalStreaming.Recv()
		if io.EOF == err {
			return nil
		} else if nil != err {
			return err
		}
		if err = bidirectionalStreaming.Send(NewGreetResponse(mssg.Name + ", you're welcome here!")); err != nil {
			return err
		}
	}
}

func (*GreetController) GreetWithTimeout(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	return Greet(ctx, req)
}
