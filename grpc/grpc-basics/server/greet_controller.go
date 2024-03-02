package server

import (
	"context"
	pb "grpc-basics/proto"
	"io"
	"strings"
)

type GreetController struct {
	pb.GreetServiceServer
}

func (*GreetController) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
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
