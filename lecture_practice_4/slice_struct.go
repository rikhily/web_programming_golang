package main

import "fmt"

type Residence struct {
	name string
	city string
}

func main() {
	user_slice := []Residence { {"His", "Fresno"},
				    {"Her", "Clovis"},
				}
	
	for _, value := range user_slice {
		fmt.Println(value.name + " --> " + value.city)
	}
}
