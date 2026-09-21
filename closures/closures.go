package main

import "fmt"

// Here getInt function returns another function and closes over the variable i
// A closure is a function that "closes over" variables from its outer scope, remembering them even after the outer function returns.
func getInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	incrementingFunction := getInt()

	fmt.Println(incrementingFunction())
	fmt.Println(incrementingFunction())
	fmt.Println(incrementingFunction())

	secondIncrementingFunction := getInt()
	fmt.Println(secondIncrementingFunction())
}
