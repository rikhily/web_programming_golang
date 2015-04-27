package main

import "fmt"

//(1 ft = 0.3048 m)

func main() {
	
	var dimen_ft float64= 1362
	
	fmt.Print("The height of WTC is: ")
	fmt.Println(dimen_ft)
	
	var f2m_conv = 0.3048

	var dimen_m = dimen_ft * f2m_conv

	fmt.Print("The height in meters is: ")
	fmt.Println(dimen_m)

}
