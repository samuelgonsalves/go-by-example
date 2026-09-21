package main

import "fmt"

func main() {
	s := "hello"

	// This does not return the rune but rather the Unicode value
	for _, c := range s {
		fmt.Println(c)
	}

	// This also doesn't return the individual runes but the hex values of them
	for index := range s {
		fmt.Printf("%x", s[index])
		fmt.Println()
	}

	// This formats it into the rune characters
	for index := range s {
		fmt.Printf("%c", s[index])
		fmt.Println()
	}

}
