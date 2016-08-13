package main

import "fmt"

// var initializers
var i, j int = 1, 2

func main() {
	// type can be omitted if initializer is present
	var c, python, java = true, false, "no!"

	// := can only be used inside a func
	a, b := true, "yes!"

	fmt.Println(i, j, c, python, java)

	fmt.Println(a, b)
}
