package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/wudtichaikarun/grpc/calculator"
	"google.golang.org/grpc"
)

type number struct {
	Value int32
}

func randomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	min := 1
	max := 100

	return random.Intn(max-min+1) + min
}

func setInterval(interval time.Duration, f func(number, number)) chan bool {
	stop := make(chan bool)

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a := number{Value: int32(randomNumber())}
				b := number{Value: int32(randomNumber())}
				f(a, b)
			case <-stop:
				return
			}
		}
	}()

	return stop
}

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

	// Define a function to call the Subtract method
	callSubtract := func(A number, B number) {
		subtractRequest := &calculator.SubtractRequest{A: A.Value, B: B.Value}
		subtractResponse, err := client.Subtract(context.Background(), subtractRequest)
		if err != nil {
			fmt.Printf("Subtract failed: %v\n", err)
			return
		}
		fmt.Printf("Subtract result: %d\n", subtractResponse.Result)
	}

	// Set up the interval (e.g., 2 seconds)
	interval := 2 * time.Second

	// Start the interval function
	stop := setInterval(interval, callSubtract)

	// Run for a certain duration (e.g., 10 seconds)
	time.Sleep(15 * time.Second)

	// Stop the interval function after a certain duration
	close(stop)

}
