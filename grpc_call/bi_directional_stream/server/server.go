package main

import (
	v1 "famesensor/go-grpc-learn/grpc_call/bi_directional_stream/proto/product"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	v1.UnimplementedProductServiceServer
}

func main() {
	logrus.Info("Starting server..")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterProductServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (srv *server) CreateProducts(stream v1.ProductService_CreateProductsServer) error {
	logrus.Info("Create Product Function")

	for {
		req, err := stream.Recv()

		// If there are no more requests, we return
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}
		log.Printf("Product-Name: %s, Product-Description: %s, Product-Total: %d, Product-Price: %f \n", req.Name, req.Desc, req.Total, req.Price)

		res := stream.Send(&v1.CreateProductsResponse{
			ProductId:   fmt.Sprintf(`product_%s`, uuid.New()),
			ProductName: req.Name,
		})

		if res != nil {
			log.Fatalf("Error when response was sent to the client: %v", res)
		}
	}
}
