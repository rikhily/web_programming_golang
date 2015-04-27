package main

import "fmt"

type Residence struct {
	name string
	city string
}

func main() {
	
	user_slice := []Residence{ {"Him", "Fresno"}, 
				   {"Her", "Clovis"},
				   {"They", "Sanger"},
			}

	fmt.Println("#Initial Slice elements")
	for _, value := range user_slice {
		fmt.Println(value)
	}

	users_slice := []Residence {
				{"I", "SFO"},
				{"You", "Berkeley"},
				{"We", "It's LA baby!"},
			}

	//now to append the two slices
	user_slice = append(user_slice, users_slice...)

	fmt.Println()
	fmt.Println("#After Appending")
	for _, value := range user_slice{
		fmt.Println(value)
	}

	//now slicing the slices
	user_slice = append(user_slice[:2], users_slice[1:]...)
	
	fmt.Println()
	fmt.Println("#After deletion")
	for _, value := range user_slice {
		fmt.Println(value)
	}
}
