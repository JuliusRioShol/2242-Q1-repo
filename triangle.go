// quiz1
package main

import (
	"fmt"
)

// 1.
type Triangle struct {
	base, height float64
}

// 2.
func (t Triangle) area() float64 {
	return (t.base * t.height) / 2
}

func main() {

	fmt.Println("Hello Cookie")

}
