package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"example.com/grpc/service"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type helloWorldServer struct {
	service.UnimplementedHelloWorldServer

	// internal string to say
	str string
}

func newServer() *helloWorldServer {
	s := &helloWorldServer{
		str: "0",
	}
	return s
}

func (sev *helloWorldServer) SetStr(s string) {
	sev.str = s
}

func (sev *helloWorldServer) Increment() {
	i := 0
	for {
		i++
		// Sleep 5 seconds before setting.
		time.Sleep(1 * time.Second)
		sev.SetStr(fmt.Sprintf("%d", i))
	}
}

// Say a simple hello world.
func (s *helloWorldServer) Say(ctx context.Context, req *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	return &service.HelloWorldResponse{
		Data: &service.Data{Text: s.str},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := newServer()
	// Keeps interating on the transactions.
	go server.Increment()

	grpcServer := grpc.NewServer()
	service.RegisterHelloWorldServer(grpcServer, server)
	log.Println("Starting to serve at port:", port)
	grpcServer.Serve(lis)
}
