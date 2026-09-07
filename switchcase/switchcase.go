package main

import (
	"fmt"
	"strings"
)

func main() {
	printMatchingNames("sam")
	printMatchingNames("helen")
	printMatchingNames("hope")

	printNumbers(2)
	printNumbers(5)
}

func printMatchingNames(name string) {
	// acts as an if-else if chain
	switch {
	case strings.HasPrefix(name, "s"):
		fmt.Println(name, "starts with s")
	case strings.HasPrefix(name, "he") && strings.HasSuffix(name, "en"):
		fmt.Println(name, "starts with he and ends with en")
	default:
		fmt.Println("No checks matched")
	}
}

func printNumbers(num int) {
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Some other value!")
	}
}
