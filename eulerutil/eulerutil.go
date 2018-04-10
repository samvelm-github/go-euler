package eulerutil

import (
	"fmt"
)

// Reverse - string reverse
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// PrimeFactorization - takes a number on the input and returns an array
// of prime factorsint
func PrimeFactorization(aNumber int) []int {

	fmt.Printf(">> aNumber: %d\n", aNumber)

	n := aNumber
	factors := make([]int, 0)

	for i := 2; i <= n; i++ {
		for {
			if n%i != 0 {
				break
			} else {
				factors = append(factors, i)
				n = n / i
			}
		}
	}

	return factors
}
