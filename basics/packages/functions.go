package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// grouped params declaration
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Named return values (hoisted to top of func); Naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println(add(42, 13))

	fmt.Println(split(17))
}
