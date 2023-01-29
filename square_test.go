package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	//call the thing we want to test.
	sA, sP := Square(5)

	sA_expected := 25.00
	sP_expected := 20.000

	if (sA_expected != sA) || (sP_expected != sP) {
		if sA_expected != sA {
			t.Errorf("got %f, expected %f", sA, sA_expected)
		}
	}
	if sP_expected != sP {
		t.Errorf("got %f, expected %f", sP, sP_expected)
	}
}
