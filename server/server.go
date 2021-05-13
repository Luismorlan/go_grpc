package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	"example.com/grpc/service"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type helloWorldServer struct {
	service.UnimplementedHelloWorldServer
}

func newServer() *helloWorldServer {
	s := &helloWorldServer{}
	return s
}

// Say a simple hello world.
func (s *helloWorldServer) Say(ctx context.Context, req *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	return &service.HelloWorldResponse{
		Text: "hello world",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterHelloWorldServer(grpcServer, newServer())
	log.Println("Starting to serve at port:", port)
	grpcServer.Serve(lis)
}
