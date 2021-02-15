package main

import (
	"context"
	"go-grpc/proto"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *proto.Request)
