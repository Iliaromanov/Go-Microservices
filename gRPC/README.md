## Running the API Server and Client Server

- Run API GRPC Server:
> `go run server/main.go`
 
- Run Client REST Server:
> `go run client/main.go`

URL for addition method:

http://localhost:4000/add/A/B
where A and B are some integers

URL for multiplication method:

http://localhost:4000/mult/A/B
where A and B are some integers

## General Notes

https://grpc.io/ 

https://developers.google.com/protocol-buffers 

gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

- gRPC APIs are language neutral (like REST)
- protocol buffers serve a similar purpose to json or xml files (serialization), but are much smaller and much more efficient

### Setup

- Install protocol buffer compiler and add it to PATH
- Install the main grpc package
> `go get -u google.golang.org/grpc`
- Install protocol buffer package for go
> `go get -u github.com/golang/protobuf/protoc-gen-go`

### Proto File

- proto file tells Go how it should encode/decode various pieces of data
- the package defined in proto file will be the name of the Go module we work with. eg. `package proto;`
- message is used to define the structure of Requests/Responses when serialized. eg:
```Go
message Request {
  int64 a = 1;
  int64 b = 2;
}
```
  The numbers that a,b are used for serialization. Rhs being 1-15 means the keys to the lhs values will take up 1 byte in the buffer when the request is serialized. 16-2047 means the keys take up 2 bytes.
- define a service using the 'service' keyword. Eg:
```Go
service NameOfService {
  rpc NameOfFunc1(Request) returns (Response);
  rpc NameOfFunc2(Request) returns (Response);
  ...
}
```
  The Request and Response in the above eg. should be defined with the 'message' keyword.
- Proto files are compiled into Go files
> `protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto`

  use `protoc.exe` if on windows
  
- The third_party folder is required to compile the .proto file and its path must be provided in the --proto_path flag in the command above. (its location relative to other files does not have to be the same as it is in this repo, it just has to be accessable when compiling)
