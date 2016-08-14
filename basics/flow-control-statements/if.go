package main

import "fmt"
import "math"

/**
 * No parentesis again, but braces are required.
 */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(math.Abs(x)) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

/**
 * Like FOR, IF statement can start with a short statement to execute before the condition.
 *
 * Variables declared by the statement are only in scope until the end of the IF
 */
func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}

	return limit
}

/**
 * Variables declared inside an if short statement are also available
 * inside any of the else blocks.
 */
func elseExample(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return limit
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
