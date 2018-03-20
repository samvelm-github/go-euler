package main

import "testing"

func TestRun(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1000, 233168},
		{100, 2318},
		{50, 543},
	}
	for _, c := range cases {
		got := Run(c.in)
		if got != c.want {
			t.Errorf("Run(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
