package server

import (
	"context"
	pb "grpc-basics/proto"
)

type GreetController struct {
	pb.GreetServiceServer
}

func (*GreetController) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return Greet(ctx, req)
}
