package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/credentials"
	"log"

	"google.golang.org/grpc"
	pb "grpc-examples/helloworld/pb"
)

var (
	address = flag.String("addr", "localhost:8972", "address")
	name    = flag.String("n", "world", "name")
)

func main() {
	flag.Parse()
	c, err := credentials.NewClientTLSFromFile("../conf/server.pem", "luojie")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

	// 连接服务器
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("faild to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
