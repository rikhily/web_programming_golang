package main

import "fmt"

func main() {
	var x = 0
	for i :=0; i<10; i++ {
		for j :=0; j<10; j++ {
			x = i + j
			fmt.Println(x)
		}
	}
}
