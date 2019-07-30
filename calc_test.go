package twelvefa

import (
	"math"
	"testing"
)

type testCase struct {
	a, b     int64
	expected int64
	hasErr   bool // error expected?
}

// tests functions that look like fn func(a, b int64) (int64, error)
func testWithError(t *testing.T, fn func(a, b int64) (int64, error), name string, testCases []*testCase) {
	for _, tc := range testCases {
		if res, err := fn(tc.a, tc.b); (err != nil) != tc.hasErr {
			if tc.hasErr {
				t.Errorf("expected %s(%d, %d) to return an error", name, tc.a, tc.b)
				return
			}
			t.Errorf("did not expect %s(%d, %d) to return an error", name, tc.a, tc.b)
			return
		} else if res != tc.expected {
			t.Errorf("expected %s(%d, %d) == %d, got %d", name, tc.a, tc.b, tc.expected, res)
			return
		}
	}
}

// tests functions that look like fn func(a, b int64) int64
func testWithoutError(t *testing.T, fn func(a, b int64) int64, name string, testCases []*testCase) {
	for _, tc := range testCases {
		if res := fn(tc.a, tc.b); res != tc.expected {
			t.Errorf("expected %s(%d, %d) == %d, got %d", name, tc.a, tc.b, tc.expected, res)
			return
		}
	}
}

func TestAdd(t *testing.T) {
	testCases := []*testCase{
		&testCase{
			a:        1,
			b:        1,
			expected: 2,
			hasErr:   false,
		},
		&testCase{
			a:        -7,
			b:        -5,
			expected: -12,
			hasErr:   false,
		},
		&testCase{
			a:        7,
			b:        -5,
			expected: 2,
			hasErr:   false,
		},
		&testCase{
			a:        -7,
			b:        5,
			expected: -2,
			hasErr:   false,
		},
		&testCase{
			a:        8,
			b:        0,
			expected: 8,
			hasErr:   false,
		},
		&testCase{
			a:        math.MaxInt64,
			b:        465,
			expected: 0,
			hasErr:   true,
		},
		&testCase{
			a:        -1 * math.MaxInt64,
			b:        -465,
			expected: 0,
			hasErr:   true,
		},
	}

	testWithError(t, Add, "Add", testCases)
}
func TestMod(t *testing.T) {
	testCases := []*testCase{
		&testCase{
			a:        1,
			b:        1,
			expected: 0,
			hasErr:   false,
		},
		&testCase{
			a:        -7,
			b:        -5,
			expected: -2,
			hasErr:   false,
		},
		&testCase{
			a:        7,
			b:        -5,
			expected: 2,
			hasErr:   false,
		},
		&testCase{
			a:        -7,
			b:        5,
			expected: -2,
			hasErr:   false,
		},
		&testCase{
			a:        8,
			b:        0,
			expected: 0,
			hasErr:   true,
		},
		&testCase{
			a:        math.MaxInt64,
			b:        465,
			expected: 7,
			hasErr:   false,
		},
		&testCase{
			a:        -1 * math.MaxInt64,
			b:        -465,
			expected: -7,
			hasErr:   false,
		},
	}

	testWithError(t, Mod, "Mod", testCases)
}

func TestMax(t *testing.T) {
	testCases := []*testCase{
		&testCase{
			a:        1,
			b:        1,
			expected: 1,
		},
		&testCase{
			a:        -7,
			b:        -5,
			expected: -5,
		},
		&testCase{
			a:        7,
			b:        -5,
			expected: 7,
		},
		&testCase{
			a:        -7,
			b:        5,
			expected: 5,
		},
		&testCase{
			a:        8,
			b:        0,
			expected: 8,
		},
		&testCase{
			a:        math.MaxInt64,
			b:        465,
			expected: math.MaxInt64,
		},
		&testCase{
			a:        -1 * math.MaxInt64,
			b:        -465,
			expected: -465,
		},
	}

	testWithoutError(t, Max, "Max", testCases)
}

func TestMin(t *testing.T) {
	testCases := []*testCase{
		&testCase{
			a:        1,
			b:        1,
			expected: 1,
		},
		&testCase{
			a:        -7,
			b:        -5,
			expected: -7,
		},
		&testCase{
			a:        7,
			b:        -5,
			expected: -5,
		},
		&testCase{
			a:        -7,
			b:        5,
			expected: -7,
		},
		&testCase{
			a:        8,
			b:        0,
			expected: 0,
		},
		&testCase{
			a:        math.MaxInt64,
			b:        465,
			expected: 465,
		},
		&testCase{
			a:        -1 * math.MaxInt64,
			b:        -465,
			expected: -1 * math.MaxInt64,
		},
	}

	testWithoutError(t, Min, "Min", testCases)
}

func TestSieve(t *testing.T) {
	type testCase struct {
		n        uint64
		expected uint64
	}

	primes := SieveOfEratosthenes(104743 + 1)

	testCases := []*testCase{
		&testCase{
			n:        12,
			expected: 37,
		},
		&testCase{
			n:        9658,
			expected: 100801,
		},
		&testCase{
			n:        789,
			expected: 6047,
		},
		&testCase{
			n:        10001,
			expected: 104743,
		},
		&testCase{
			n:        2,
			expected: 3,
		},
		&testCase{
			n:        2345,
			expected: 20849,
		},
	}

	for _, tc := range testCases {
		if actual := primes[tc.n-1]; actual != tc.expected {
			t.Errorf("expected %dth prime to be %d, got %d", tc.n, tc.expected, actual)
			// return
		}
	}
}

func TestNthPrime(t *testing.T) {
	if nth := NthPrime(10001); nth != 104743 {
		t.Errorf("expcted 104743, got %d", nth)
	}
}

func BenchmarkNthPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// random number between 1 and 10001
	}
}
