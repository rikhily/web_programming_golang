package main

import "fmt"

type Residence struct {
	state string
	country string
}

func (user_res Residence) displayRes() {
	fmt.Println("In method : " + user_res.state + " and " + user_res.country)
}

func main() {
	
	var his = Residence{"California", "U.S.A"}
	fmt.Println("From method access: " + his.state)
	fmt.Println("From method access: " + his.country)

	fmt.Println()
	
	his.displayRes()

}
