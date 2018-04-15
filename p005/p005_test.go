package main

import "testing"

func TestProblemFive(t *testing.T) {
	limit := 20
	expected := 232792560
	lcm := ProblemFive(limit)

	if lcm != expected {
		t.Errorf("ProblemFive(%d) expected %d, actual %d\n", limit, expected, lcm)
	}
}
