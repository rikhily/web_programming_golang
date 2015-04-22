package main

import "fmt"

type Student struct {
	name string
	id int
}

func main() {
	
	var c = Student{}

	c.name = "invinci"
	c.id = 9989

	fmt.Println(c.name)
	fmt.Println(c.id)
}
