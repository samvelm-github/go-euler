package main

import (
	"fmt"

	"github.com/user/go-euler/eulerutil"
)

func main() {
	limit := 20
	lcm := ProblemFive(limit)

	fmt.Println(lcm)
}

// ProblemFive - Solution for problem five from project euler
// Least common multiplier for two numbers can be calculated as
// lcm = n1 * n2 / lcd(n1, n2), where lcd is the least common
// divizor for the same two numbers
func ProblemFive(limit int) int {
	var lcm = 1
	for i := 2; i <= limit; i++ {
		lcd := eulerutil.GCD(lcm, i)
		lcm = lcm * i / lcd
	}

	return lcm
}
