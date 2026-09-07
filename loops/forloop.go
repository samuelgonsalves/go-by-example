package main

import "fmt"

func main() {
	fmt.Println(sumOf10NaturalNumbersTheHardWay())
	printOddNumbers()

	whileCondition(5)
	loopWithARange(5)
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

func whileCondition(n int) {
	i := 0
	for i < n {
		i++
	}

	fmt.Println(i)
}

func loopWithARange(n int) {
	fmt.Println("Range loop")

	// range is 0 to n-1
	for i := range n {
		fmt.Print(i, ",")
	}
}
