package main

import "fmt"

func main() {
	
	text := "Trying out pointer data and memory access"

	var p1 *string = &text
	
	var p2 string = *p1

	var p3 *string = &p2

	fmt.Println("p1 has", p1)
	fmt.Println("p2 has", p2)
	fmt.Println("p3 has", p3)
}
