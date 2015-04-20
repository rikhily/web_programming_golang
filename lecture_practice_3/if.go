package main

import "fmt"

func check(condition bool) {
	if condition {
		fmt.Println("That's how we roll")
	} else {
		fmt.Println("This behaious isn't intended")
	}
}

func main() {
	check(true)

	check(false)
}
