package main

import "fmt"

//(C = (F - 32) * 5/9)

func main() {
	
	fmt.Println("Code converts the temperature from Fahrenheit to Celsius")
	const temp_F = 425.0

	var temp_C = ((temp_F - 32) * (5 / 9.0))

	fmt.Print("Temperatur in F: " )
	fmt.Println(temp_F)

	fmt.Print("Temperature in C: ")
	fmt.Println(temp_C)

}
