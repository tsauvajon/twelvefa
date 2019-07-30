package main

import (
	"context"

	"github.com/tsauvajon/twelvefa/calc"
)

// client contains the methods available with the gRPC client
// it implements the calc.CalcClient interface
type client struct {
	address string
	calc    calc.CalcClient
}

// I decided to use the calc responses for simplicity
// but we should make our response structs in this package
func (c *client) Max(a, b int64) (*calc.MaxResponse, error) {
	req := &calc.MaxRequest{
		A: a,
		B: b,
	}
	return c.calc.Max(context.Background(), req)
}

func (c *client) Add(a, b int64) (*calc.AddResponse, error) {
	req := &calc.AddRequest{
		A: a,
		B: b,
	}

	return c.calc.Add(context.Background(), req)
}

func (c *client) NthPrimes(n []uint64) (*calc.NthPrimesResponse, error) {
	req := &calc.NthPrimesRequest{
		N: n,
	}

	return c.calc.NthPrimes(context.Background(), req)
}
