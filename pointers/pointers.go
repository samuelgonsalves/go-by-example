package main

import "fmt"

func main() {
	ptr := 1
	incrementer(&ptr)

	fmt.Println(ptr)
}

func incrementer(ptr *int) {
	fmt.Println("memory addr", ptr)
	*ptr++
}
