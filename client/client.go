package main

import (
	"context"
	"famesensor/go-grpc-learn/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello client...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := proto.NewHelloServiceClient(cc)
	request := &proto.HelloRequest{Name: "Fame"}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response : ", res.Greeting)
}
