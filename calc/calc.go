// Package calc provides basic calculations
package calc

import (
	"fmt"
)

// Add returns the sum of two signed integers
func Add(a, b int64) (int64, error) {
	// check for overflows
	if (a+b > a) != (b > 0) {
		return 0, newErrIntOverflow(fmt.Sprintf("%d + %d", a, b))
	}

	return a + b, nil
}

// Mod returns the remainder (or modulo) of the dividend divided by the divisor. That's a lot of "divi" in 5 words!
func Mod(dividend, divisor int64) (int64, error) {
	if divisor == 0 {
		return 0, newErrZeroDivision(fmt.Sprintf("%d + %d", dividend, divisor))
	}

	return dividend % divisor, nil
}

// Max returns the biggest of two values
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the biggest of two values
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// NthPrimes : solve Project Euler problem #7 up to the 10001th prime
// using the Sieve of Eratosthenes
// https://projecteuler.net/problem=7
func NthPrimes(n []uint64) ([]uint64, error) {
	// 10001th prime is 104743
	n10001 := uint64(104743)
	primes := SieveOfEratosthenes(n10001 + 1)

	nthPrimes := make([]uint64, len(n))

	for i, nthPrime := range n {
		if nthPrime > 10001 || nthPrime < 1 {
			return nil, newErrOutOfRange("1 <= n <= 10001")
		}
		nthPrimes[i] = primes[nthPrime-1]
	}
	return nthPrimes, nil
}

// SieveOfEratosthenes : pre-calculate all the prime numbers up to a ertain point
// to find the nth prime more easily
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func SieveOfEratosthenes(n uint64) (primes []uint64) {
	isNPrime := make([]bool, n)

	for i := uint64(2); i < n; i++ {
		if isNPrime[i] {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < n; k += i {
			isNPrime[k] = true
		}
	}
	return
}
