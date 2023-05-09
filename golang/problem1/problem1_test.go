package problem1_test

import (
	"euler/golang/problem1"
	"testing"
)

func TestMul3of5(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		N    int
		want int
	}{
		{N: 10, want: 23},
		{N: 5, want: 3},
		{N: 20, want: 78},
	}
	for _, tc := range testCases {
		got := problem1.Mul3of5(tc.N)
		if got != tc.want {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}
