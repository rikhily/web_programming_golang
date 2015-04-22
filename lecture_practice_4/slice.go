package main

import "fmt"

func main() {
	
	user_slice := []int{1, 2, 3, 4, 5}

	for _, value := range user_slice{
		fmt.Println(value)
	}
}
