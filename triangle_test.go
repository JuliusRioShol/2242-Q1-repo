package main

import (
	"testing"
)

// 11.
func TestArea(t *testing.T) {
	//call the thing that we want to test.
	t1 := Triangle{
		base:   2.5,
		height: 5,
	}
	got := t1.Area()

	//What i should get back:

	expected := 6.25

	//compare the results received with the expected value
	if got != expected {
		t.Errorf("got %f expected %f", got, expected)
	}

}
