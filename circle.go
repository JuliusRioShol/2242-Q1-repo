package main

import (
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
func main() {

}
