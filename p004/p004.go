package main

import (
	"fmt"
	"strconv"

	"github.com/user/go-euler/eulerutil"
)

func main() {
	maxPod, maxI, maxJ := ProblemFour()

	fmt.Printf("Largest palindrome: %d\t product of %d and %d", maxPod, maxI, maxJ)
}

// ProblemFour - solution for problem four from Project Euler
func ProblemFour() (int, int, int) {
	maxPod := 0
	maxI := 0
	maxJ := 0
	// prod := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			prod := i * j
			sProd := strconv.Itoa(prod)
			if prod > maxPod && eulerutil.IsPalindrome(sProd) {
				maxPod = prod
				maxI = i
				maxJ = j
			}
		}
	}

	return maxPod, maxI, maxJ
}
