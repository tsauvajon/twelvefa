package main

import (
	"log"
	"net"
	"os"

	"github.com/tsauvajon/twelvefa/calc"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	port := os.Getenv("PORT")

	// default value
	if port == "" {
		port = "80"
	}

	// Open a TCP listener
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register the Calc gRPC server and serve it on the listener
	s := grpc.NewServer()
	calc.RegisterCalcServer(s, &server{})
	s.Serve(listener)
}
