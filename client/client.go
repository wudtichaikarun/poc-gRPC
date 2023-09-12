package main

import (
	"context"
	"log"

	"github.com/wudtichaikarun/grpc/calculator"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	// Calling the Add method
	addRequest := &calculator.AddRequest{A: 10, B: 5}
	addResponse, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	log.Printf("Add result: %d", addResponse.Result)

	// Calling the Subtract method
	subtractRequest := &calculator.SubtractRequest{A: 10, B: 5}
	subtractResponse, err := client.Subtract(context.Background(), subtractRequest)
	if err != nil {
		log.Fatalf("Subtract failed: %v", err)
	}
	log.Printf("Subtract result: %d", subtractResponse.Result)
}
