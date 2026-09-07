package main

import "fmt"

func main() {
	// Short-hand variables
	intValue := 1
	intValue++
	fmt.Println("Incremented value:", intValue)

	isSuccess := true
	fmt.Println("Is success?", isSuccess)

	// Cannot be assigned again
	const constantInt = 1
	fmt.Println("Constant value:", constantInt)
}
