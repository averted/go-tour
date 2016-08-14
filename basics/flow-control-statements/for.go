package main

import "fmt"

func main() {
	sum := 0

	// no parenthesis, braces are always required
	for i := 0; i < 10; i++ {
		sum += i
	}

	// init and post statements are optional..
	// making it exact replica of C's while loop
	for sum < 1000 {
		sum += sum
	}

	// infinite loop:
	// for {
	// }

	fmt.Println(sum)
}
