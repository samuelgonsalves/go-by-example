package main

import "fmt"

func main() {
	fmt.Println(sumOf10NaturalNumbersTheHardWay())
	printOddNumbers()
}

func sumOf10NaturalNumbersTheHardWay() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum = sum + i
	}

	return sum
}

func printOddNumbers() {
	fmt.Print("Odd Numbers: ")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Print(i, ",")
		}
	}
	fmt.Println()
}
