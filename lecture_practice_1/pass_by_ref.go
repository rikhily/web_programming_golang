package main

import "fmt"

func main() {

	text := "using pass by reference"
	fmt.Println("message is", text)

	var indirect *string = &text
	fmt.Println("accessing memory location", indirect)

	*indirect = "yo"
	fmt.Println("")
	fmt.Println(text, *indirect)	
}
