package server

import (
	"context"
	pb "grpc-basics/proto"
)

func Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: "Hi," + req.Name,
	}, nil
}
