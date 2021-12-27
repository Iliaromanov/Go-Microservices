/* Simple server that can generate the  sum or product of two integers */
package main

import (
	"context"
	"net"
	"github.com/Iliaromanov/Go-Microservices/gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":8000");
	if err != nil {
		panic(err);
	}

	srv := grpc.NewServer();  // create new gRPC server (from google.golang.org/grpc)
	proto.RegisterAddServiceServer(srv, &server{}); // register our service on the created server
	reflection.Register(srv); // for serializing and de-serializing

	// call serve on created server and have it listen on tcp 8000
	if e := srv.Serve(listener); e != nil {
		panic(e);
	}
}

func (s *server) Add(ctxt context.Context, request *proto.Request) (*proto.Response, error) {
	// the getter methods (GetA and GetB) are defined in the generated Go file obtained by
	//   compiling our defined service.proto file
	a, b := request.GetA(), request.GetB();
	result := a + b;
	return &proto.Response{Result: result}, nil; // return pointer to the proto Resposne struct
}

func (s *server) Multiply(ctxt context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB();
	result := a * b;
	return &proto.Response{Result:result}, nil;
}