package main

import "fmt"

func swap(a , b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 1, 2
	fmt.Println("Before swap: ", x, y, )
	
	swap(&x, &y)
	fmt.Println("After swap: ", x, y)
}	
