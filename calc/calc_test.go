package calc

import (
	"fmt"
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

func BenchmarkAdd(b *testing.B) {
	var n int64
	for i := 0; i > b.N; i++ {
		n, _ = Add(math.MaxInt64, 56)
	}
	n++
}

func BenchmarkAddOverflow(b *testing.B) {
	for i := 0; i > b.N; i++ {
		Add(math.MaxInt64, -123457)
	}
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

func TestNthPrimes(t *testing.T) {
	in := []uint64{12, 9658, 789, 10001, 2, 2345}
	expected := []uint64{37, 100801, 6047, 104743, 3, 20849}
	out, err := NthPrimes(in)

	if err != nil {
		t.Errorf("test failed with error: %s", err)
	}

	if len(out) != len(expected) {
		t.Errorf(
			"the slice length doesn't match - expected %d elements, got %d",
			len(expected),
			len(out),
		)
		return
	}

	for i, actual := range out {
		if actual != expected[i] {
			t.Errorf("expected the %dth prime to be %d, got %d", i, expected[i], actual)
		}
	}
}

func BenchmarkSieve10(b *testing.B) {
	var (
		primes []uint64
		p10    = uint64(29)
	)

	for i := 0; i < b.N; i++ {
		primes = SieveOfEratosthenes(p10 + 1)
	}

	_ = fmt.Sprint(primes[9])
}

func BenchmarkSieve10001(b *testing.B) {
	var (
		primes []uint64
		p10001 = uint64(104743)
	)

	for i := 0; i < b.N; i++ {
		primes = SieveOfEratosthenes(p10001 + 1)
	}

	_ = fmt.Sprint(primes[10000])
}
