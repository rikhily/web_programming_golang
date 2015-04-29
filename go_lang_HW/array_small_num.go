package main

import "fmt"

func main() {

	x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }
	
	var small = 0

	//the length is taken one less
	//as the iterator needs to check 
	//until the second last element
	for i:=0; i<len(x) - 1; i++ {
		if (x[i] < x[i+1]) {
			small = x[i]
		} else {
			small = x[i+1]
		}
	}
	
	fmt.Print("The smallest number in the array: ")
	fmt.Println(small)
}
