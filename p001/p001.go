package main

import "fmt"

func main() {
	fmt.Println(run(1000))
}

func run(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
