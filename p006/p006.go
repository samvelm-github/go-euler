package main

import (
	"fmt"
)

func main() {
	sumf1 := ProblemSix(100)
	sumf2 := ProblemSixMath(100)

	fmt.Printf(">> sumf1: %d\t sumf2: %d", sumf1, sumf2)
}

// ProblemSix - find the diffence between sum of squares and square of sums
// of first N subsequent numbers
func ProblemSix(aNumber int) int {
	sumOne := 0
	sumTwo := 0
	for i := 1; i <= aNumber; i++ {
		sumOne = sumOne + i
		sumTwo = sumTwo + i * i
	}

	return sumOne * sumOne - sumTwo
}

// ProblemSixMath - implementation is based on math formula
// sum of first n subsequent numbers - n * (n + 1) / 2
// sum of squares of n subsequent numbers - n * (n + 1) * (2 * n + 1) / 6
func ProblemSixMath(aNumber int) int {
	sumOne := aNumber * (aNumber + 1) / 2
	sumTwo := aNumber * (aNumber + 1) * (2 * aNumber + 1) / 6 

	return sumOne * sumOne - sumTwo
}
