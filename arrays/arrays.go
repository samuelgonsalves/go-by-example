package main

import "fmt"

var nums [5]int

func main() {
	nums := []int{1, 2, 3, 4, 5}
	printArray(nums)

	changeSign(nums)

	printArray(nums)

	primes := []int{2, 3, 5, 7}

	printArray(primes)
}

func changeSign(nums []int) {
	for i := range nums {
		nums[i] = nums[i] * (-1)
	}
}

func printArray(nums []int) {
	for i := range nums {
		fmt.Print(nums[i], ",")
	}
	fmt.Println()
}
