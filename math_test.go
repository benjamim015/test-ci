package main

import "testing"

func TestSum(t *testing.T) {

	sum := Sum(15, 15)

	if sum != 30 {
		t.Errorf("The sum result is invalid: Result %d. Expected: %d", sum, 30)
	}
}
