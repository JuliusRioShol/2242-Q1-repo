package main

import (
	"fmt"
	"math"
)

// 9.
func square(side float64) (float64, float64) {
	area := math.Pow(side, 2)
	perimeter := 4 * side

	return area, perimeter

}
func main() {

	area, perimeter := square(5)

	fmt.Println("Square file. ", area, perimeter)
}
