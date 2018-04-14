package eulerutil

import (
	"fmt"
)

// IsPalindrome - returns true if the string is palindrome and false otherwise
func IsPalindrome(s string) bool {
	ret := true
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			ret = false
			break
		}
	}

	fmt.Printf("Input: %s \t IsPalindrome: %t\n", s, ret)

	return ret
}

// Reverse - string reverse
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	ret := string(r)
	fmt.Printf("Input: %s \t Output: %s\n", s, ret)

	return ret
}

// PrimeFactorization - takes a number on the input and returns an array
// of prime factorsint
func PrimeFactorization(aNumber int) []int {

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

// MaxInArray - returns the largest number in array of integer values
func MaxInArray(arr []int) int {
	max := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return max
}

// MinInArray - returns the smallest number in array of integer values
func MinInArray(arr []int) int {
	min := arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
