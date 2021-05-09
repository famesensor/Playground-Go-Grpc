package main

import (
	"context"
	v1 "famesensor/go-grpc-learn/grpc_call/client_stream/proto/message"
	"fmt"
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
	client := v1.NewMessageServiceClient(con)

	messages := []string{"hello service", "", "message from client", "message : you must send data to me...", ""}
	// Get the stream and err
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error on Send Message: %v", err)
	}

	for _, message := range messages {
		// Start making streaming requests by sending
		// each book object inside the request message
		fmt.Println("Client streaming request: \n", message)
		stream.Send(&v1.MessageRequest{
			Message: message,
		})
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing the stream and receiving the response: %v", err)
	}

	// Print the response errors message
	fmt.Printf("Message errors: %v \n", res.MessageError)
}
