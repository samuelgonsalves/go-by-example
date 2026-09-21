package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}

	fmt.Println(arr)

	s := make([]int, 3)
	printSlice(s)

	// Appending to the slice
	s = append(s, 3)

	// Interestingly, the capacity doubles here since the underlying array is full
	printSlice(s)

	// Setting values
	s[0] = 0
	s[1] = 1
	s[2] = 2
	printSlice(s)

	// Using the colon operator in this fashion a slice can be created
	a := arr[1:3]
	printSlice(a)

	// [i:] From index i to end of array
	b := arr[2:]
	printSlice(b)

	// [:j] From index 0 to j-1
	c := arr[:2]
	printSlice(c)
}

func printSlice(slice []int) {
	fmt.Println("slice =", slice, "len =", len(slice), "cap =", cap(slice))
}
