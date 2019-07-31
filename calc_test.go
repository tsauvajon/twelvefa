package main

import (
	"context"
	"math"
	"testing"

	"github.com/tsauvajon/twelvefa/calc"
)

func TestAdd(t *testing.T) {
	s := &server{}
	req := &calc.AddRequest{
		A: 2,
		B: 3,
	}
	resp, err := s.Add(context.Background(), req)

	if err != nil {
		t.Errorf("Add failed with error: %s", err)
		return
	}

	if resp.Sum != 5 {
		t.Errorf("expected Add(2, 3) to be 5, got %d", resp.Sum)
	}
}

func TestAddError(t *testing.T) {
	s := &server{}
	req := &calc.AddRequest{
		A: 2,
		B: math.MaxInt64,
	}
	_, err := s.Add(context.Background(), req)

	if err == nil {
		t.Error("expected Add to return an error")
	}
}

func TestMax(t *testing.T) {
	s := &server{}
	req := &calc.MaxRequest{
		A: 2,
		B: 3,
	}
	resp, err := s.Max(context.Background(), req)

	if err != nil {
		t.Errorf("Max failed with error: %s", err)
		return
	}

	if resp.Max != 3 {
		t.Errorf("expected Max(2, 3) to be 3, got %d", resp.Max)
	}
}

func TestNthPrimes(t *testing.T) {
	s := &server{}
	req := &calc.NthPrimesRequest{
		N: []uint64{12, 9658},
	}
	expected := []uint64{37, 100801}
	resp, err := s.NthPrimes(context.Background(), req)

	if err != nil {
		t.Errorf("NthPrimes failed with error: %s", err)
		return
	}

	for i, nth := range resp.NthPrimes {
		if expected[i] != nth {
			t.Errorf("Expected %dth prime to be %d, got %d", req.N[i], expected[i], nth)
		}
	}
}

func TestNthPrimesError(t *testing.T) {
	s := &server{}
	req := &calc.NthPrimesRequest{
		N: []uint64{4659658},
	}
	_, err := s.NthPrimes(context.Background(), req)

	if err == nil {
		t.Error("expected Add to return an error")
	}
}
