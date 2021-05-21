package main

import (
	"context"
	v1 "famesensor/go-grpc-learn/grpc_call/bi_directional_stream/proto/product"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.Println("Client Stream...")

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	client := v1.NewProductServiceClient(con)

	stream, err := client.CreateProducts(context.Background())
	if err != nil {
		log.Fatalf("Error when getting stream object: %v", err)
		return
	}

	// data...
	data := []*v1.CreateProductsRequest{
		{
			Name:  "A",
			Desc:  "this is A",
			Total: 10,
			Price: 200,
		},
		{
			Name:  "B",
			Desc:  "this is B",
			Total: 100,
			Price: 20,
		},
		{
			Name:  "C",
			Desc:  "this is C",
			Total: 15,
			Price: 100,
		},
	}

	// Create a new channel
	waitResponse := make(chan struct{})

	// Use a go routine to send request messages to the server
	go func() {
		// Iterate over the requests slice
		for index, req := range data {
			log.Println("send to service: ", index+1)
			// Send request message
			stream.Send(req)

			// Sleep for a little bit..
			time.Sleep(1000 * time.Millisecond)
		}
		// Close stream
		stream.CloseSend()
	}()

	// Use a go routine to receive response messages from the server
	go func() {
		for {
			// Get response and possible error message from the stream
			res, err := stream.Recv()

			// Break for loop if there are no more response messages
			if err == io.EOF {
				break
			}

			// Handle a possible error
			if err != nil {
				log.Fatalf("Error when receiving response: %v", err)
			}

			// Log the response
			fmt.Println("Server Response: ", res)
		}

		// Close channel
		close(waitResponse)
	}()
	<-waitResponse
}
