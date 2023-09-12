package main

import (
	"context"
	"log"
	"net"

	"github.com/wudtichaikarun/grpc/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	// https://github.com/parthw/fun-coding/blob/main/golang/understanding-grpc-change/main.go
	calculator.CalculatorServer
}

func (*server) mustEmbedUnimplementedCalculatorServer() {
}

func (s *server) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.A + req.B
	return &calculator.AddResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, req *calculator.SubtractRequest) (*calculator.SubtractResponse, error) {
	result := req.A - req.B
	return &calculator.SubtractResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculator.RegisterCalculatorServer(s, &server{})

	// Enable reflection for tools like grpcurl
	reflection.Register(s)

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
