package main

import "fmt"

type user_string string

func (s1 user_string) return1() {
	fmt.Println("This is S1!")
}

func (s2 user_string) return2() {
	fmt.Println("This is S2!")
}

func main() {
	
	var message user_string = "THIS IS A STRING!"
	fmt.Println("This is main:" + message)

	message.return1()
	message.return2()

}
