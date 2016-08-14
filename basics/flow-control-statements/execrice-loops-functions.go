package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	prev, z := 0.0, 1.0

	for math.Abs(prev-z) > 1e-8 {
		prev = z

		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(36) == math.Sqrt(36))
}
