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
	fmt.Println(his.house_no)
	fmt.Println(his.street)
	fmt.Println(his.apt_no)
	fmt.Println(his.city)
	fmt.Println(his.state)
	fmt.Println(his.zip)
}

func main() {
	var invinci = Residence{}
	invinci.house_no = 4969
	invinci.street = "N. Backer Avenue"
	
	//call the func
	address(invinci)
}
