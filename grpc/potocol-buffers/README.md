# cheetsheet for Protocol buffer


## Official documentation
- [proto 3](https://protobuf.dev/programming-guides/proto3/)
- [Updating A Message Type](https://protobuf.dev/programming-guides/proto3/#updating)
- [Protocol Buffers Well-Known Types](https://protobuf.dev/reference/protobuf/google.protobuf/#index)
- [Protocol Buffer all options](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto)
- [Protocol buffer style guide](https://protobuf.dev/programming-guides/style/)
- [FAQ](https://protobuf.dev/reference/go/faq/)
- some examples
  - [Clement-Jean/proto-course](https://github.com/Clement-Jean/proto-course)
  - [Clement-Jean/proto-go-course](https://github.com/Clement-Jean/proto-go-course/tree/main)
  - [Protocol Buffers - Code Example](https://github.com/protocolbuffers/protobuf/tree/main/examples)
  - [Google Common Types](https://github.com/googleapis/googleapis/tree/master/google/type)
  - [protobuf](https://github.com/protocolbuffers/protobuf/tree/main/src/google/protobuf)
- **videos**
  - [This is why gRPC was invented](https://www.youtube.com/watch?v=u4LWEXDP7_M)
  - [Protobuf - How Google Changed Data Serialization FOREVER](https://www.youtube.com/watch?v=FR754e5xIwg)



In order to perform code generation, you will need to install `protoc`  on your computer.</br>
It is actually very easy, open a command line interface and type `brew install protobuf` </br>
Find the correct protocol buffers version based on your [Linux Distro](https://github.com/google/protobuf/releases) </br>
Example with x64: </br>
```
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

Download the [windows archive](https://github.com/google/protobuf/releases) </br>
Extract all to `C:\proto3` </br>  
Your directory structure should now be: </br>
```
C:\proto3\bin 
C:\proto3\include 
```
Finally, add `C:\proto3\bin` to your PATH:
- From the desktop, right-click My Computer and click Properties.
- In the System Properties window, click on the Advanced tab
- In the Advanced section, click the Environment Variables button.
- Finally, in the Environment Variables window (as shown below), highlight the Path variable in the Systems Variable section and click the Edit button. Add or modify the path lines with the paths you wish the computer to access. Each different directory is separated with a semicolon as shown below.  `C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin`

you need to add ; `C:\proto3\bin`  at the end </br>

#### How to run protoc in your local developement project
Setup your protoc with the following steps
```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
source ~/.zshrc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go mod tidy
```
For windows please check the documentation


#### Sample protoc commands to compile your proto files
```
# we can compile the entire folder
protoc --go_out=./src/go proto/*.proto
protoc --java_out=./src/java proto/*.proto
protoc --python_out=./src/python proto/*.proto

# we can also compile by the file
protoc --go_out=./src/go proto/address.proto
protoc --java_out=./src/java proto/address.proto
protoc --python_out=./src/python proto/address.proto


# if we specify these options in the proto files itself then we will
# no longer need to specify the packages for the protoc compiler
option go_package = "/src/go";
option java_package = "/src/java";
option py_package = "/src/python";

protoc --go_out=. proto/*.proto
protoc --java_out=. proto/*.proto
protoc --python_out=. proto/*.proto

# Make sure to to import the other proto files
# if we are using them in the current proto file
# use message with the proper . convention if that proto file have any package declairation
```



#### This is a sample proto file
first we have to mention the syntax then we will create a proto messge, which is similar to struct </br>
a message contains fields, which each field consists of [ field type , field name = tag; ] </br>
```
syntax = "proto3";

message Account{
    uint32 id = 1;
    string name = 2;
    bool is_verified = 3;
}
```

when we do not set any field then the field will not be serialized </br>
we will not get any payload overhead for any empty data, though we can use some default value </br>