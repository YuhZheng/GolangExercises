package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prevz := x 
	for ; math.Abs(prevz - z) > 1e-15;  {
		prevz = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(2)-math.Sqrt(2))
}

