package main

import "testing"

func TestCalcFibo(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 89},
		{50, 20365011074},
	}
	for _, c := range cases {
		got := CalcFibo(c.in)
		if got != c.want {
			t.Errorf("Run(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCalcFiboRecoursive(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 89},
		{50, 20365011074},
	}
	for _, c := range cases {
		got := CalcFiboRecoursive(c.in)
		if got != c.want {
			t.Errorf("Run(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestProblemTwo(t *testing.T) {
	limit := 4000000
	expectedValue := 4613732
	actualValue := ProblemTwo(limit)
	if actualValue != expectedValue {
		t.Errorf("ProblemTwo(%d) == %d, expected %d", limit, actualValue, expectedValue)
	}
}
