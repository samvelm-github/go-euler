package main

import "testing"

func TestGetLargestPrimeFactor(t *testing.T) {
	aNumber := 600851475143
	expectedValue := 6857
	actualValue := GetLargestPrimeFactor(aNumber)
	if actualValue != expectedValue {
		t.Errorf("GetLargestPrimeFactor(%d) == %d, expected %d", aNumber, actualValue, expectedValue)
	}
}
func TestGetSmallestPrimeFactor(t *testing.T) {
	aNumber := 600851475143
	expectedValue := 71
	actualValue := GetSmallestPrimeFactor(aNumber)
	if actualValue != expectedValue {
		t.Errorf("GetSmallestPrimeFactor(%d) == %d, expected %d", aNumber, actualValue, expectedValue)
	}
}
