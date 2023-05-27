package main

import (
	"context"
	"golang-test-task/rpc"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This is used just to check that server actually responses to client requests.
// It sends a request every second and logs the results.
func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cli error: %v", err)
	}
	defer func() { conn.Close() }()

	cli := rpc.NewSolutionClient(conn)

	for {
		resp, err := cli.CountCycles(context.TODO(), &rpc.CountCyclesRequest{})
		if err != nil {
			log.Printf("cli error: %v", err)
		} else {
			log.Printf("cli response: %v", resp.NCycles)
		}
		time.Sleep(1 * time.Second)
	}
}
