package main

import (
	"golang-test-task"
	"golang-test-task/rpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Start the server.
func main() {
	solutionServer := golang_test_task.NewSolutionServer()
	grpcServer := grpc.NewServer()
	rpc.RegisterSolutionServer(grpcServer, solutionServer)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error running grpc server: %v", err)
	}
}
