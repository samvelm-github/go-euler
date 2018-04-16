package main

import "testing"

func TestProblemSix(t *testing.T) {
	sumf1 := ProblemSix(100)
	sumf2 := ProblemSixMath(100)

	if sumf1 != sumf2 {
		t.Errorf("Diffence between sum of squares and square of sums calculated in 2 different ways is not the same: %d, %d", sumf1, sumf2)
	}

}
