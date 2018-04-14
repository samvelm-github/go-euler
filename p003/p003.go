package main

import (
	"fmt"

	"github.com/user/go-euler/eulerutil"
)

func main() {
	aNumber := 600851475143
	largestPrimeFactor := GetLargestPrimeFactor(aNumber)
	smallestPrimeFactor := GetSmallestPrimeFactor(aNumber)

	fmt.Printf("%d\tLargest prime factor: %d\t Smallest prime factor: %d", aNumber, largestPrimeFactor, smallestPrimeFactor)
}

// GetLargestPrimeFactor - implements prime factorization on a given number
// and returns the largest one
func GetLargestPrimeFactor(aNumber int) int {
	factors := eulerutil.PrimeFactorization(aNumber)
	largestPrimeFactor := eulerutil.MaxInArray(factors)

	return largestPrimeFactor
}

// GetSmallestPrimeFactor - implements prime factorization on a given number
// and returns the smallest one
func GetSmallestPrimeFactor(aNumber int) int {
	factors := eulerutil.PrimeFactorization(aNumber)
	smallestPrimeFactor := eulerutil.MinInArray(factors)

	return smallestPrimeFactor
}
