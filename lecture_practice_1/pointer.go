package main

import "fmt"

func main() {
	text := "Using a pointer"

	var data *string = &text

	fmt.Println(text, data)
}
