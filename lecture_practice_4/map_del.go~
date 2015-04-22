package main

import "fmt"

func main() {
	
	user_map := map[string]string {
			"Agent23": "On mission!",
			"Agent27": "Lost contact 17 minutes ago!",
			"Agent47": "Hunting the target!",
			"Agent37": "Injured!",
			"Agent53": "Comprimised!",
		}

	var x = 0
	for _ = range user_map {
		x++
		
	}
	
	//get the number of key-val pairs
	fmt.Println()
	fmt.Println("Current Key-Value pairs = ", x)

	//check  for the desired key
	fmt.Println("Key to be deleted: Agent53")
	val, exists := user_map["Agent53"]
	fmt.Println("val: ", val)
	fmt.Println("exists: ", exists)
	fmt.Println()
	
	//delete the key-val pair
	if val, exists := user_map["Agent53"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("Sir, The Agent has shown reluctant efforts")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
		
}
