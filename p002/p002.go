package main

import "fmt"

// Find the sum of all the multiples of 3 or 5 below 1000
func main() {
	fmt.Printf("\n\n fibonacci sequence for first %d numbers: %d", 10, CalcFibo(10))

	fmt.Printf("\n\n fibonacci sequence for first %d numbers: %d", 50, CalcFiboRecoursive(50))

	limit := 4000000
	fmt.Printf("\n\n Problem two from project euler for %d limit: %d", limit, ProblemTwo(limit))
}

// CalcFibo - calculates Fibonacci sequence for the firs
// n numbers given by its input parameter
func CalcFibo(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	n1 := 1
	n2 := 2
	var fibo = 0
	for index := 2; index < n; index++ {
		fibo = n1 + n2
		n1 = n2
		n2 = fibo
	}
	return fibo
}

// CalcFiboRecoursive recoursive implementation of the fibonacci sequence
func CalcFiboRecoursive(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	n1 := 1
	n2 := 2

	return calcFiboRecursiveInternal(n1, n2, 2, n)
}

func calcFiboRecursiveInternal(n1 int, n2 int, index int, n int) int {
	if index == n {
		return n2
	}

	var fibo = n1 + n2
	n1 = n2
	n2 = fibo

	return calcFiboRecursiveInternal(n1, n2, index+1, n)
}

// ProblemTwo p002 from project euler
func ProblemTwo(limit int) int {
	n1 := 1
	n2 := 2
	customFibo := 0

	for n1 <= limit {
		if n1%2 == 0 {
			customFibo += n1
		}

		var fibo = n1 + n2
		n1 = n2
		n2 = fibo

	}

	return customFibo
}
