package main

import (
	"context"
	"errors"
	v1 "famesensor/go-grpc-learn/unary/proto/v1"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	v1.UnimplementedHelloServiceServer
}

func (s *server) Password(ctx context.Context, request *v1.PasswordRequest) (*v1.PasswordResponse, error) {
	if request.Password.Password != request.Password.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, errors.New("Error password").Error())
	}

	return &v1.PasswordResponse{
		Status:  "success",
		Message: "Password",
	}, nil
}

func (s *server) Hello(ctx context.Context, request *v1.HelloRequest) (*v1.HelloReponse, error) {
	name := request.Name
	reponse := &v1.HelloReponse{
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
	v1.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
