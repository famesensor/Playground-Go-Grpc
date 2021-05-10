package main

import (
	v1 "famesensor/go-grpc-learn/grpc_call/client_stream/proto/message"
	"fmt"
	"io"
	"net"

	log "github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type server struct {
	v1.UnimplementedMessageServiceServer
}

func main() {
	logrus.Info("Starting server..")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// ValidateBooks function
func (srv *server) SendMessage(stream v1.MessageService_SendMessageServer) error {
	logrus.Info("Send Message Function")

	// Initialize the message reponse
	errors := []*v1.MessageError{}
	index := 1
	for {

		// Start receiving stream messages from the client
		req, err := stream.Recv()

		// Check if the stream has finished
		if err == io.EOF {
			// Close the connection and return the response to the client
			return stream.SendAndClose(&v1.MessageResponse{
				MessageError: errors,
			})
		}

		// Handle any possible errors while streaming requests
		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}

		message := req.Message
		if message == "" {
			err := &v1.MessageError{
				Message: fmt.Sprintf("%s %d %s", "message", index, "is send failed"),
			}

			errors = append(errors, err)
		}

		index += 1
	}
}
