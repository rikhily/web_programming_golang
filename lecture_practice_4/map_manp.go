package main

import "fmt"

func main() {
	user_slice := map[string]string { "Agent23": "On mission!",
					  "Agent27": "Lost contact 17 minutes ago!",
					  "Agent47": "Hunting the target!",
					}

	user_slice["Agent37"] = "Injured!"
	user_slice["Agent53"] = "Comprimised!"
	
	for value := range user_slice {
		fmt.Println(key, value)
	}
}
