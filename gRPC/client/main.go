package main

import (
	"github.com/Iliaromanov/Go-practice/gRPC/proto"
	"github.com/gin-gonic/gin" // using gin for creating api endpoints
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure()); // connect to server on port 8000 as defined in server
	// .WithInsecure because server is not using HTTPs
	if err != nil {
		panic(err);
	}

	// add connection using the generated NewAddServiceClient method
	client := proto.NewAddServiceClient(conn);

	g := gin.Default();
}