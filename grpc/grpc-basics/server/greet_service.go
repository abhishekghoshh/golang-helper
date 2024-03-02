package server

import (
	"context"
	pb "grpc-basics/proto"
)

func Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return NewGreetResponse(req.Name), nil
}
func NewGreetResponse(name string) *pb.GreetResponse {
	return &pb.GreetResponse{
		Message: "Hi," + name,
	}
}
