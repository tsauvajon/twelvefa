package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/tsauvajon/twelvefa/calc"

	"google.golang.org/grpc"
)

func init() {
	log.SetOutput(os.Stdout)
}

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

	// Capture os signals to a channel
	c := make(chan os.Signal, 2)

	// Both interruption signals are to be caught
	// It will prevent from quitting
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Wait for the next interruption signal
		<-c

		// Make sure everything is stopped properly before actually stopping
		// the server, and display a message to inform the user that we
		// understood his request
		log.Println("stopping the server")
		s.GracefulStop()

		// Actually exit
		os.Exit(1)
	}()
}
