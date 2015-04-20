package main

import "fmt"

type Residence struct {
	street string
	city string
}

func address(his Residence) {
	his_street, his_city := DisplayAddress(his.street, his.city, "hey")

	fmt.Println(his_street)
	fmt.Println(his_city)	
}

//variadic function
func DisplayAddress(street string, city ...string) (user_street string, user_city string) {
	user_street = street[1] + street[0]
	user_city = city
	return
}

func main() {
	var text = Residence{"N. Backer Ave", "Fresno"}
	address(text)
}
