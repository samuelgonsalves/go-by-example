package main

import "fmt"

func main() {
	intValue := 1
	intValue++
	fmt.Println("Incremented value:", intValue)

	isSuccess := true
	fmt.Println("Is success?", isSuccess)

	const constantInt = 1
	fmt.Println("Constant value:", constantInt)
}
