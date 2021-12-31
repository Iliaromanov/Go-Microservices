/* Server */
package main

import (
	"os"
	"context"
	"github.com/Iliaromanov/Go-Microservices/gRPC_Docker-ChatApp/proto"

	glog "google.golang.org/grpc/grpclog"
	
)

var grpcLog glog.LoggerV2

// Executed when first running the server
func init() {
	// Define what io writer the logger will use for info, warning, and error messages
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout); // use stdout for all outputs
}

type Connection struct {
	stream proto.Broadcast_CreateStreamServer // allows streaming messages between server and client
	id 	   string
	active bool
	// channel for errors, channels are used for sending and recieving data
	// useful info on channels in go: https://www.geeksforgeeks.org/channel-in-golang/ 
	error  chan error // channel because using go routines
}

func main() {

}

func (s * server) CreateStream(ctxt context.Context, conn *proto.Connect) (*proto.Message, error) {

}

func (s *server) BroadcastMessage(ctxt context.Context, msg *proto.Message) (*proto.Close, error) {

}