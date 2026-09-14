package mathx

import "testing"

func TestAbsDiff(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{5, 3, 2},
		{3, 5, 2},
		{0, 0, 0},
		{-5, 5, 10},
	}
	for _, c := range cases {
		if got := AbsDiff(c.a, c.b); got != c.want {
			t.Errorf("AbsDiff(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
