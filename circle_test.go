package main

import (
	"testing"
)

func TestCircleArea(t *testing.T) {
	//call the thing we want to test.
	c1 := Circle{
		radius: 5.0,
	}
	got := c1.Area()

	//What I should get.
	expected := 78.53982

	//an helper function, that compares the two float numbers to see if they are
	//almost the same.
	equal := withTolerance(got, expected)

	if equal != true {
		t.Errorf("got %f, expected %f", got, expected)
	}

}

func TestCirclePerimeter(t *testing.T) {
	//call the thing we want to test.
	c1 := Circle{
		radius: 5.0,
	}
	got := c1.Perimeter()

	//What i should get.
	expected := 31.41593
	//an helper function, that compares the two float numbers to see it they are
	//almost the same.
	equal := withTolerance(got, expected)
	if equal != true {
		t.Errorf("got%f, expected %f", got, expected)
	}

}
