package main

import "fmt"

func greatest(num ...int) int {
	curr := 0
	for _, value := range num {
		if value > curr {
			curr = value
		}
	}
	return curr
}

func main() {

	val := greatest(12, 43, 56, 34, 87, 75, 98, 65)
	
	fmt.Println("The greatest value in the list is: ", val)
	
}


