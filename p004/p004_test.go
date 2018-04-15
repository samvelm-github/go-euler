package main

import "testing"

func TestProblemFour(t *testing.T) {
	maxPod, maxI, maxJ := ProblemFour()

	if maxPod != 906609 || maxI != 913 || maxJ != 993 {
		t.Errorf("Expected values: 906609, 913, 993\t received: %d, %d, %d", maxPod, maxI, maxJ)
	}
}
