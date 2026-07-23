# gRPC

[gRPC](https://grpc.io/) is a high-performance, open-source universal RPC framework.

## Sub-topics

- [Protocol Buffers](./protocol-buffers.md) — Protobuf schema, scalar types, enums, oneof, nesting, imports
- [gRPC Basics](./grpc-basics.md) — Server, client, 4 streaming types, TLS, interceptors, error handling

## Setup

```bash
export GO111MODULE=on
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Tools for Testing gRPC

- [grpcui](https://github.com/fullstorydev/grpcui) — Web UI for gRPC
- [grpcurl](https://github.com/fullstorydev/grpcurl) — curl for gRPC
- [evans-cli](https://github.com/ktr0731/evans) — Interactive gRPC client

## Reverse Proxy

- [envoyproxy](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/other_protocols/grpc) — gRPC reverse proxy
- [grpc-web](https://github.com/grpc/grpc-web) — gRPC for browser clients

## Reference Repositories

- [dreamsofcode-io/grpc](https://github.com/dreamsofcode-io/grpc)
- [Clement-Jean/proto-course](https://github.com/Clement-Jean/proto-course)
- [Clement-Jean/proto-go-course](https://github.com/Clement-Jean/proto-go-course)

## Key External Resources

- [Protocol Buffers Documentation](https://protobuf.dev/programming-guides/proto3/)
- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [LinkedIn Adopts Protobufs — 60% Latency Reduction](https://www.infoq.com/news/2023/07/linkedin-protocol-buffers-restli/)
