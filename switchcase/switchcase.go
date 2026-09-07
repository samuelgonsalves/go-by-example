package main

import (
	"fmt"
	"strings"
)

func main() {
	printMatchingNames("sam")
	printMatchingNames("helen")
	printMatchingNames("hope")
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
