package main

import (
	"context"

	"github.com/tsauvajon/twelvefa/calc"
)

// server contains the methods available with the gRPC server
// it implements the calc.CalcServer interface
type server struct{}

func (s *server) Add(ctx context.Context, in *calc.AddRequest) (*calc.AddResponse, error) {
	sum, err := calc.Add(in.A, in.B)

	return &calc.AddResponse{Sum: sum}, err
}

func (s *server) Max(ctx context.Context, in *calc.MaxRequest) (*calc.MaxResponse, error) {
	max := calc.Max(in.A, in.B)

	return &calc.MaxResponse{Max: max}, nil
}

func (s *server) NthPrimes(ctx context.Context, in *calc.NthPrimesRequest) (*calc.NthPrimesResponse, error) {
	primes, err := calc.NthPrimes(in.GetN())
	return &calc.NthPrimesResponse{NthPrimes: primes}, err
}
