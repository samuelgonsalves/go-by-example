package main

import "fmt"

func main() {
	fmt.Println(adder(2, 3))
	fmt.Println(sameArgsAdder(2, 3))
}

func adder(a int, b int) int {
	return a + b
}

func sameArgsAdder(a, b int) int {
	return a + b
}
