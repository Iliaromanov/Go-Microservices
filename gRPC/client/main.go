package main

import (
	"net/http"
	"strconv"

	"github.com/Iliaromanov/Go-Microservices/gRPC/proto"
	"github.com/gin-gonic/gin" // using gin for creating api endpoints
	"google.golang.org/grpc"
)

func main() {
	// connect to server on port 8000 as defined in the main.go in server
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure()); // .WithInsecure because server not using HTTPS
	if err != nil {
		panic(err);
	}

	// create a client object using the generated NewAddServiceClient method
	client := proto.NewAddServiceClient(conn);

	// create a gin object which can be used for creating endpoint methods
	g := gin.Default();

	// add .GET method endpoints for our created server methods
	//	passing their arguments through URL parameters
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		// ParseInt(str, base, bitSize)
		// ctx.Param("a") gets the string in place of :a in the url
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64);
		if err != nil {
			// gin.H adds a new key-value pair to the JSON response
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"});
			return;
		}
		
		// ctx.Param("b") gets the string in place of :b in the url
		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64);
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"});
			return;
		}

		request := &proto.Request{A: int64(a), B: int64(b)};

	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64);
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"});
			return;
		}
		
		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64);
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"});
			return;
		}

		request := &proto.Request{A: int64(a), B: int64(b)};
	})
}