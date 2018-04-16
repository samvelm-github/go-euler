package main

import (
	"fmt"

	"github.com/user/go-euler/eulerutil"
)

func main() {
	limit := 10001
	lcm := ProblemSeven(limit)

	fmt.Println(lcm)
}

// ProblemSeven - returns nth prime number
func ProblemSeven(n int) int {
	counter := 0
	aPrime := 0
	for {
		if eulerutil.IsPrime(aPrime) {
			counter++
			if counter == n {
				break
			}
		}
		aPrime++
	}
	return aPrime
}
