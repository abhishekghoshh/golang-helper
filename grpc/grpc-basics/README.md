# gRPC basics


#### build your project manually
```
GO111MODULE=1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
protoc -Iproto --go_opt=module=grpc-basics --go_out=. --go-grpc_opt=module=grpc-basics --go-grpc_out=. proto/*.proto
```


#### test your gRPC from cli
use [evans-cli](https://github.com/ktr0731/evans) </br>
Follow the instruction section and install it if necessary </br>