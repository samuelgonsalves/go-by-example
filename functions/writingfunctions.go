package main

import "fmt"

func main() {
	fmt.Println(adder(2, 3))
	fmt.Println(sameArgsAdder(2, 3))
	fmt.Println(returnManyValues(2, 3))
	fmt.Println(variableNumberOfArguments(1, 2, 3, 4, 5))
}

func adder(a int, b int) int {
	return a + b
}

// Same typed arguments can be simplified
func sameArgsAdder(a, b int) int {
	return a + b
}

// Avoids complexities around wrapping objects just to return them in a single type
func returnManyValues(a, b int) (int, int) {
	return a + 1, b + 1
}

func variableNumberOfArguments(a ...int) int {
	sum := 0
	// Note: the _ is ignored index, num is the value
	for _, num := range a {
		sum += num
	}
	return sum
}
