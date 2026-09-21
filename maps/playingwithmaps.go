package main

import "fmt"

func main() {

	mapping := map[string]int{"user1": 2, "user2": 3}

	fmt.Println(mapping)

	// Set user1 value
	mapping["user1"] = 3
	fmt.Println(mapping)

	// Increment user1 value
	mapping["user1"]++
	fmt.Println(mapping)

	// Adds a user3 to the map
	mapping["user3"] = 1
	fmt.Println(mapping)

	// Returns the zero-value for the type (in this case 0 since it is an int being stored)
	user4 := mapping["user4"]
	fmt.Println(user4)

	// Removing the user3 from the map
	delete(mapping, "user3")
	fmt.Println(mapping)

}
