package main

import (
	"fmt"
	"math"
)

// 6.
type Circle struct {
	radius float64
}

// 7.
func (c Circle) area() float64 {
	const Pi = 3.14159
	return Pi * (math.Pow(c.radius, 2))
}

// 8.
func (c Circle) perimeter() float64 {
	const Pi = 3.14159
	return 2 * (Pi * c.radius)
}
func main() {

	c1 := Circle{
		radius: 3,
	}

	fmt.Println(c1.area())
	fmt.Println(c1.perimeter())
}
