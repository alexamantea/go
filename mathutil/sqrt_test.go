package mathutil

import "testing"

func TestSqrt(t *testing.T) {
	cases := []struct {
		in   float64
		want float64
	}{
		{4, 2},
		{9, 3},
		{144, 12},
		{10000, 100},
		{625, 25},
	}
	for _, c := range cases {
		got := Sqrt(c.in)
		if got != c.want {
			t.Errorf("SqrtTest(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
