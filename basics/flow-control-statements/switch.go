package main

import "fmt"
import "time"
import "runtime"

/**
 * Case body breaks automatically, unless it ends with 'fallthrough' statement.
 */
func main() {
	t := time.Now()
	today := time.Now().Weekday()

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}

	/**
	 * Cases are evaluated from top to bottom, stopping when case succeeds.
	 */
	fmt.Print("When's Saturday? - ")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away..")
	}

	/**
	 * Switch without conditional (e.g, switch true) - same as if else.
	 */
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
