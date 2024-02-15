# protocol-buffers
Google Protocol Buffers


```
option go_package = "/src/go";
option java_package = "/src/java";
// option py_package = "/src/python";

export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
source ~/.zshrc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --go_out=. proto/*.proto
protoc --java_out=. proto/*.proto
protoc --python_out=./src/python proto/*.proto
```