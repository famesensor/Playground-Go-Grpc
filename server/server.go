package main

import (
	"context"
	"go-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReponse, error) {
	name := request.Name
	reponse := &proto.HelloReponse{
		Greeting: "Hello " + name,
	}
	return reponse, nil
}

func main() {
	// register server and it listen on port 3000 for grpc cpnnection
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	// auto generate file protobuf file
	proto.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
