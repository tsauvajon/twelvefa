package main

import (
	"os"
	"testing"
	"time"

	"github.com/tsauvajon/twelvefa/calc"
	"google.golang.org/grpc"
)

// This test requires a 'twelvefa' service to be up and running
func TestClient(t *testing.T) {
	address := os.Getenv("TWELVEFA_ADDRESS")

	if address == "" {
		address = ":3000"
	}

	client := &client{
		address: address,
	}

	conn, err := grpc.Dial(client.address,
		grpc.WithBlock(),                // check we can reach the server
		grpc.WithTimeout(5*time.Second), // but not forever
		grpc.WithInsecure(),
	)

	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client.calc = calc.NewCalcClient(conn)

	resm, err := client.Max(2, 999)

	if err != nil {
		t.Errorf("Max failed: %v", err)
		return
	}

	if resm.Max != 999 {
		t.Errorf("expected Max(2, 999) to be 999, got %d", resm.Max)
		return
	}

	resa, err := client.Add(2, 999)

	if err != nil {
		t.Errorf("Add failed: %v", err)
		return
	}

	if resa.Sum != 1001 {
		t.Errorf("expected Add(2, 999) to be 1001, got %d", resa.Sum)
		return
	}

	resnp, err := client.NthPrimes([]uint64{12})

	if err != nil {
		t.Errorf("NthPrimes failed: %v", err)
		return
	}

	if len(resnp.NthPrimes) != 1 || resnp.NthPrimes[0] != 37 {
		t.Errorf("expected NthPrimes([12]) to be [37], got %v", resnp.NthPrimes)
		return
	}
}
