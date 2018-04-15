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
func ProblemFive(limit int) int {
	var lcm = 1
	for i := 2; i <= limit; i++ {
		lcd := eulerutil.GCD(lcm, i)
		lcm = lcm * i / lcd
	}

	return lcm
}
