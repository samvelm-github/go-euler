package eulerutil

import (
	"fmt"
	"testing"
)

func TestPrimeFactorization(t *testing.T) {
	aNumber := 600851475143
	expected := [2]int{11, 11}
	factors := PrimeFactorization(aNumber)

	fmt.Printf("%v\n", factors)

	passed := true

	if len(factors) != len(expected) {
		t.Errorf("Expected size: %d, got: %d\n", len(expected), len(factors))
	} else {
		for i := range factors {
			if factors[i] != expected[i] {
				passed = false
				break
			}
		}
		if !passed {
			t.Errorf("Expecting: %v, got: %v\n", expected, factors)
		}
	}
}

func TestReversePos(t *testing.T) {
	hello := "Hello World"
	expected := "dlroW olleH"
	reverse := Reverse(hello)
	if reverse != reverse {
		t.Errorf("Reverse(%s) == %s, want %s", hello, reverse, expected)
	}
}
