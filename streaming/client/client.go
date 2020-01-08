package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"google.golang.org/grpc"
	pb "grpc-examples/streaming/pb"
)

var (
	address = flag.String("addr", "localhost:8972", "address")
	name    = flag.String("n", "world", "name")
	example = flag.Int("e", 0, "0 for replay streaming, 1 for request streaming, 2 for bi-streaming")
)

func main() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair("./conf/client.pem", "./conf/client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "go-grpc-example",
		RootCAs:      certPool,
	})
	// 连接服务器
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("faild to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	switch *example {
	case 0:
		log.Println("Test Reply streaming")
		sayHello1(client)
	case 1:
		log.Println("\n\n\nTest Request streaming")
		sayHello2(client)
	case 2:
		log.Println("\n\n\nTest Bidirectional streaming")
		sayHello3(client)
	}
}

func sayHello1(c pb.GreeterClient) {
	stream, err := c.SayHello1(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
		}

		log.Printf("Greeting: %s", reply.Message)
	}
}

func sayHello2(c pb.GreeterClient) {
	var err error
	stream, err := c.SayHello2(context.Background())

	for i := 0; i < 100; i++ {
		if err != nil {
			log.Printf("failed to call: %v", err)
			break
		}
		stream.Send(&pb.HelloRequest{Name: *name + strconv.Itoa(i)})
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to recv: %v", err)
	}

	log.Printf("Greeting: %s", reply.Message)
}

func sayHello3(c pb.GreeterClient) {
	var err error

	stream, err := c.SayHello3(context.Background())
	if err != nil {
		log.Printf("failed to call: %v", err)
		return
	}

	var i int64
	for {
		stream.Send(&pb.HelloRequest{Name: *name + strconv.FormatInt(i, 10)})
		if err != nil {
			log.Printf("failed to send: %v", err)
			break
		}
		reply, err := stream.Recv()
		if err != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("Greeting: %s", reply.Message)
		i++
	}
}
