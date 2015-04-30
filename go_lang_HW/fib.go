package main

import "fmt"

func fib(num int) int {
	
	if(num == 0){
		return 1
	} else if (num == 1) {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

func main() {
	x := 11
	
	y := fib(x)
	
	fmt.Println("The fibonacci number of ", x, " is : ", y)
}

