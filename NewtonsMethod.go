package main

/**
	a simple program from the GoLang tutorial which uses Newton's
	method to calculate the square root of a number using a loop
 */

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10000; i++ {
		y := z
		z -= (z * z - x) / (2 * z)
		if math.Abs(z - y) < 3.0e-16 {
			return
		}
	}
	return
}

func main() {
	num := float64(2)
	fmt.Println("The square root of", num, "=", Sqrt(num))
}