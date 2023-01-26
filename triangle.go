// quiz1
package main

import (
	"fmt"
	"math"
)

// 1.
type Triangle struct {
	base, height float64
}

// 2.
func (t Triangle) area() float64 {
	return (t.base * t.height) / 2
}

// 3.
func (t Triangle) perimeter() float64 {
	//we need the hypotenuse to do:  P = a+b+c
	//hypotenues = square root (b^2 + a^2)
	//find using some usefull function in the math library.
	//	hypot := math.Sqrt((math.Pow(t.base, 2)) + (math.Pow(t.height, 2)))
	hypotenuse := math.Hypot(t.base, t.height)

	return t.base + t.height + hypotenuse

}
func main() {
	// 4.
	t1 := Triangle{
		base:   7,
		height: 10.5,
	}

	//5.
	a := t1.area()
	p := t1.perimeter()

	fmt.Println(a)
	fmt.Println(p)
}
