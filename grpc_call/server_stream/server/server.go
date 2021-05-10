package main

import (
	"famesensor/go-grpc-learn/grpc_call/server_stream/proto/temperature"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	temperature.UnimplementedTemperatureServiceServer
}

type temp struct {
	temperatures []*temperature.Temperature
}

func main() {
	logrus.Info("Server stream...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("Error run server: ", err)
	}

	srv := grpc.NewServer()
	temperature.RegisterTemperatureServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}

func (*server) GetTemperature(in *temperature.TemperatureRequest, stream temperature.TemperatureService_GetTemperatureServer) error {
	logrus.Info("request from client: ", in)

	temps := temp{}.randomTemplates()

	for _, ele := range temps {
		res := temperature.TemperatureResponse{
			Temperature: ele,
		}

		// send message to client
		stream.Send(&res)

		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (t temp) randomTemplates() []*temperature.Temperature {
	index := 0
	for {
		if index == 10 {
			break
		}

		min := 40.0
		max := 32.0
		t.temperatures = append(t.temperatures, t.getTemperature(min+rand.Float64()*(max-min)))
		index += 1
	}
	return t.temperatures
}

func (t temp) getTemperature(temp float64) *temperature.Temperature {
	return &temperature.Temperature{
		Temperature: float32(temp),
	}
}
