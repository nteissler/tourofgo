package main

import (
	"fmt"
	"math"
)

var expected float64 = float64(math.Sqrt(2))

func Sqrt(x float64) float64 {
	if x < 0 {
		return Sqrt(-x)
	}

	z := x/2.0
	d := z - expected
	fmt.Printf("d is %v\n", d)
	for math.Abs(d) > (.0001) {
		z = z - (z*z-x)/2*z
		d = z - expected
	}
	return z
}

func main() {
	fmt.Println("Your value is",
				Sqrt(2))
	fmt.Println("True value is", expected)
}
