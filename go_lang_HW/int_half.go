package main

import "fmt"

func half(num int) (int, bool){
	
	value := num/2
	
	if num%2 == 0 {
		return value, true
	}
	return value, false
}

func main() {

	val := 35
	
	fmt.Println("The original number: ", val)
	
	a, b := half(val)
	
	fmt.Println("Halved number: ", a, b)
}
