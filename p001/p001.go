package main

import "fmt"

// Find the sum of all the multiples of 3 or 5 below 1000
func main() {
	fmt.Println(Run(50))
}

func Run(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
