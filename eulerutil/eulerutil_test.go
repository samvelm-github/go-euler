package eulerutil

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		aNumber int
		isPrime bool
	}{
		{5, true},
		{11, true},
		{66, false},
	}
	for _, c := range cases {
		isPrime := IsPrime(c.aNumber)
		if isPrime != c.isPrime {
			t.Errorf("IsPrime(%d) == %t, want %t", c.aNumber, isPrime, c.isPrime)
		}
	}

}

func TestIsPalindrome(t *testing.T) {
	s := "12321"
	isPalindrome := IsPalindrome(s)

	if !isPalindrome {
		t.Errorf("The string %s is palindrome - %t\n", s, isPalindrome)
	}

	s = "123321"
	isPalindrome = IsPalindrome(s)

	if !isPalindrome {
		t.Errorf("The string %s is palindrome - %t\n", s, isPalindrome)
	}

	s = "123219"
	isPalindrome = IsPalindrome(s)

	if isPalindrome {
		t.Errorf("The string %s is palindrome - %t\n", s, isPalindrome)
	}
}

func TestGCD(t *testing.T) {
	numberOne := 54
	numberTwo := 24
	expected := 6
	gcd := GCD(numberOne, numberTwo)

	fmt.Printf("%v\n", gcd)
	if gcd != expected {
		t.Errorf("GCD(%d, %d) == %d, expecting %d", numberOne, numberTwo, gcd, expected)
	}
}

func TestPrimeFactorization(t *testing.T) {
	aNumber := 121
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
	if reverse != expected {
		t.Errorf("Reverse(%s) == %s, want %s", hello, reverse, expected)
	}
}
