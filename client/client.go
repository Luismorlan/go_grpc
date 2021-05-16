package main

import (
	"context"
	"flag"
	"log"
	"time"

	"example.com/grpc/service"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := service.NewHelloWorldClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Say(ctx, &service.HelloWorldRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Received text: ", res.Data.Text)
}
