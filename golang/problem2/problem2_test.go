package problem2_test

import (
	"euler/golang/problem2"
	"testing"
)

func BenchmarkProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem2.FibConv(4_000_000)
	}
}

func TestFibConv(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		want int
		N    int
	}{
		{want: 44, N: 89},
	}
	for _, tc := range tcs {
		got := problem2.FibConv(tc.N)
		if tc.want != got {
			t.Errorf("Want %d,got %d", tc.want, got)
		}
	}
}
