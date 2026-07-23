# gRPC Basics

gRPC is a high-performance RPC framework using Protocol Buffers for service definition and HTTP/2 for transport.

## Setup

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate code from proto
protoc -Iproto --go_opt=module=grpc-basics --go_out=. \
  --go-grpc_opt=module=grpc-basics --go-grpc_out=. proto/*.proto
```

Dependencies: `google.golang.org/grpc v1.62.0`, `google.golang.org/protobuf v1.32.0`

## Service Definition (Proto)

```protobuf
syntax = "proto3";
package greet;
option go_package = "grpc-basics/proto";

message GreetRequest  { string name = 1; }
message GreetResponse { string message = 1; }

service GreetService {
    rpc Greet(GreetRequest) returns (GreetResponse);                                   // Unary
    rpc GreetServerStreaming(GreetRequest) returns (stream GreetResponse);              // Server streaming
    rpc GrretClientStreaming(stream GreetRequest) returns (GreetResponse);              // Client streaming
    rpc GreetBidirectionalStreaming(stream GreetRequest) returns (stream GreetResponse); // Bidirectional
    rpc GreetWithTimeout(GreetRequest) returns (GreetResponse);                         // Unary with deadline
}
```

## Server Implementation

### Main Server Setup

```go
const addr = "0.0.0.0:50051"
const tls = true
const interceptorEnabled = true

func main() {
    tcpListener, _ := net.Listen("tcp", addr)
    defer tcpListener.Close()

    opts := []grpc.ServerOption{}
    opts = serverCredOptions(opts)      // TLS if enabled
    opts = addServerInterceptors(opts)  // Logging + auth if enabled

    grpcServer := grpc.NewServer(opts...)
    defer grpcServer.Stop()

    proto.RegisterGreetServiceServer(grpcServer, &server.GreetController{})
    reflection.Register(grpcServer)     // for dev tooling (evans-cli, grpcui)

    grpcServer.Serve(tcpListener)
}
```

### TLS Configuration

```go
func serverCredOptions(opts []grpc.ServerOption) []grpc.ServerOption {
    if tls {
        certFile := "ssl/server.crt"
        keyFile := "ssl/server.pem"
        creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
        opts = append(opts, grpc.Creds(creds))
    }
    return opts
}
```

SSL certificates generated via OpenSSL (see `ssl/ssl.sh` or `ssl/ssl.ps1`).

### Controller (Service Implementation)

```go
type GreetController struct {
    pb.GreetServiceServer  // embeds unimplemented server for forward compatibility
}

// Unary RPC with error handling
func (*GreetController) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
    if strings.TrimSpace(req.Name) == "" {
        return nil, status.Errorf(codes.InvalidArgument, "Empty field : name")
    }
    return &pb.GreetResponse{Message: "Hi," + req.Name}, nil
}

// Server-side streaming
func (*GreetController) GreetServerStreaming(req *pb.GreetRequest, stream pb.GreetService_GreetServerStreamingServer) error {
    for i := 0; i < 10; i++ {
        stream.Send(&pb.GreetResponse{Message: "Hi," + req.Name})
    }
    return nil
}

// Client-side streaming
func (*GreetController) GrretClientStreaming(stream pb.GreetService_GrretClientStreamingServer) error {
    var allNames strings.Builder
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.GreetResponse{Message: allNames.String()})
        }
        if err != nil { return err }
        allNames.WriteString(msg.Name + " ")
    }
}

// Bidirectional streaming
func (*GreetController) GreetBidirectionalStreaming(stream pb.GreetService_GreetBidirectionalStreamingServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF { return nil }
        if err != nil { return err }
        stream.Send(&pb.GreetResponse{Message: msg.Name + ", you're welcome here!"})
    }
}

// Unary with deadline support
func (*GreetController) GreetWithTimeout(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
    for i := 0; i < 3; i++ {
        if ctx.Err() == context.DeadlineExceeded {
            return nil, status.Error(codes.Canceled, "The client canceled the request")
        }
        time.Sleep(1 * time.Second)
    }
    return &pb.GreetResponse{Message: "Hi," + req.Name}, nil
}
```

### Server Interceptors

```go
// Logs incoming requests and metadata
func LogInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        log.Printf("Received a request: %v\n", req)
        headers, ok := metadata.FromIncomingContext(ctx)
        if ok { log.Printf("Received headers: %v\n", headers) }
        return handler(ctx, req)
    }
}

// Validates authorization header
func AuthorizationHeaderInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        headers, ok := metadata.FromIncomingContext(ctx)
        if !ok { return nil, status.Error(codes.Internal, "Error while reading the context") }
        if len(headers.Get("authorization")) == 0 {
            return nil, status.Error(codes.Unauthenticated, "Expected authorization header")
        }
        return handler(ctx, req)
    }
}

// Chain interceptors
opts = append(opts, grpc.ChainUnaryInterceptor(server.LogInterceptor()))
opts = append(opts, grpc.ChainUnaryInterceptor(server.AuthorizationHeaderInterceptor()))
```

## Client Implementation

### Client Connection Setup

```go
const clientAddress = "localhost:50051"

func clientCredOptions(opts []grpc.DialOption) []grpc.DialOption {
    if tls {
        certFile := "ssl/ca.crt"
        creds, _ := credentials.NewClientTLSFromFile(certFile, "")
        opts = append(opts, grpc.WithTransportCredentials(creds))
    } else {
        opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
    }
    return opts
}
```

### Unary Call

```go
func Test_helloUnary(t *testing.T) {
    conn, _ := grpc.Dial(clientAddress, opts...)
    defer conn.Close()

    client := proto.NewGreetServiceClient(conn)
    resp, err := client.Greet(context.Background(), &proto.GreetRequest{Name: "world"})
    log.Printf("Greeting: %s\n", resp.Message)
}
```

### Error Handling (gRPC Status Codes)

```go
func Test_helloUnaryWithError(t *testing.T) {
    _, err := client.Greet(ctx, &proto.GreetRequest{Name: "   "})
    if err != nil {
        grpcError, isGrpcError := status.FromError(err)
        if isGrpcError {
            log.Printf("Error message: %v\n", grpcError.Message())
            log.Printf("Code: %v\n", grpcError.Code())
            log.Printf("Is InvalidArgument: %t", codes.InvalidArgument == grpcError.Code())
        }
    }
}
```

### Unary with Timeout

```go
func Test_helloUnaryWithTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    _, err := client.GreetWithTimeout(ctx, &proto.GreetRequest{Name: "Name"})
    // Check for codes.DeadlineExceeded
}
```

### Server-Side Streaming

```go
func Test_helloServerStreaming(t *testing.T) {
    stream, _ := client.GreetServerStreaming(ctx, &proto.GreetRequest{Name: "world"})
    for {
        msg, err := stream.Recv()
        if err == io.EOF { break }
        responses = append(responses, msg)
    }
}
```

### Client-Side Streaming

```go
func Test_helloClientStreaming(t *testing.T) {
    stream, _ := client.GrretClientStreaming(ctx)
    for _, name := range []string{"abhishek", "nasim", "sayan"} {
        stream.Send(&proto.GreetRequest{Name: name})
    }
    resp, _ := stream.CloseAndRecv()
}
```

### Bidirectional Streaming

```go
func Test_helloBidirectionalStreaming(t *testing.T) {
    stream, _ := client.GreetBidirectionalStreaming(ctx)
    for _, name := range names {
        stream.Send(&proto.GreetRequest{Name: name})
        msg, _ := stream.Recv()
        responses = append(responses, msg)
    }
    stream.CloseSend()
}
```

### Client Interceptors

```go
// Add authorization header to every outgoing request
func addHeaderInterceptor() grpc.UnaryClientInterceptor {
    return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
        newMD := metadata.Pairs("authorization", "aDummyToken")
        ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))
        return invoker(ctx, method, req, reply, cc, opts...)
    }
}

// Chain client interceptors
clientOpts = append(clientOpts, grpc.WithChainUnaryInterceptor(addHeaderInterceptor()))
clientOpts = append(clientOpts, grpc.WithChainUnaryInterceptor(logInterceptor()))
```

## Streaming Patterns Summary

| Type | Client | Server |
|---|---|---|
| **Unary** | Send 1 | Receive 1, return 1 |
| **Server Streaming** | Send 1 | Receive 1, return stream |
| **Client Streaming** | Send stream | Receive stream, return 1 |
| **Bidirectional** | Send stream | Receive stream, return stream |

## Testing gRPC

Use [evans-cli](https://github.com/ktr0731/evans) or [grpcui](https://github.com/fullstorydev/grpcui) for interactive testing.

Enable server reflection for these tools: `reflection.Register(grpcServer)`

## References

- [gRPC Introduction](https://grpc.io/docs/what-is-grpc/introduction/)
- [Protocol Buffers Documentation](https://protobuf.dev/programming-guides/proto3/)
- [gRPC vs REST comparison](https://docs.microsoft.com/en-us/aspnet/core/grpc/comparison)
