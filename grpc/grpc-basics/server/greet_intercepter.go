package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Received a request: %v\n", req)
		headers, ok := metadata.FromIncomingContext(ctx)
		if ok {
			log.Printf("Received headers: %v\n", headers)
		}
		return handler(ctx, req)
	}
}

func AuthorizationHeaderInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "Error while reading the context")
		}
		if len(headers.Get("authorization")) == 0 {
			return nil, status.Error(codes.Unauthenticated, "Expected authorization header")
		}
		return handler(ctx, req)
	}
}
