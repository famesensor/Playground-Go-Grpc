package main

import (
	"context"
	"errors"
	"famesensor/go-grpc-learn/proto"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *server) Password(ctx context.Context, request *proto.PasswordRequest) (*proto.PasswordResponse, error) {
	if request.Password.Password != request.Password.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, errors.New("Error password").Error())
	}

	return &proto.PasswordResponse{
		Status:  "success",
		Message: "Password",
	}, nil
}

func (s *server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReponse, error) {
	name := request.Name
	reponse := &proto.HelloReponse{
		Greeting: "Hello " + name,
	}
	return reponse, nil
}

func main() {
	flag.Parse()
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
