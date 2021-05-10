package main

import (
	"context"
	"famesensor/go-grpc-learn/grpc_call/server_stream/proto/temperature"
	"io"
	"log"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.Println("Client run...")

	con, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}
	defer con.Close()

	clinet := temperature.NewTemperatureServiceClient(con)

	stream, err := clinet.GetTemperature(context.Background(), &temperature.TemperatureRequest{
		Message: "please temperature...",
	})

	for {
		// Start receiving streaming messages
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error when receiving server response stream: %v", err)
		}
		log.Printf("Response from GetDocuments: %v", res.GetTemperature())
	}
}
