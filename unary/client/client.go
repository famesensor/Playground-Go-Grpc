package main

import (
	"context"
	v1 "famesensor/go-grpc-learn/v1"
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

	client := v1.NewHelloServiceClient(cc)
	request := &v1.HelloRequest{Name: "Fame"}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response : ", res.Greeting)
}
