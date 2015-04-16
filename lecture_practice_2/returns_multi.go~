package main

import "fmt"

type Residence struct {
	house_no int
	street string
	apt_no int
	city string
	state string
	zip int
}

func address(his Residence) {
	fmt.Println(DisplayAddress(his.house_no, his.street, his.apt_no, his.city, his.state, his.zip))
}

func DisplayAddress(house_no int, street string, apt_no int, city string, state string, zip int) (int, string, int, string, string, int) {
	return house_no, street, apt_no, city, state, zip
}

func main() {

	var invinci = Residence{9989, "inct st", 89, "vinci", "vinc", 900990 }
	
	address(invinci)
}
