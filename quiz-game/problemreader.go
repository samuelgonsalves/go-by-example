package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Read the problems line by line
// Store them in a slice
// Iterate one by one over the slice asking the user the question and comparing user answer with the answer in the problem
// Keep a counter of correct responses and return that

func main() {
	problems := ReadProblems()

	c := 0
	for index, problem := range problems {

		fmt.Printf("Problem #%d: %s\n", index+1, problem[0])
		userSolution := userInput()

		if userSolution == problem[1] {
			c++
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("You got %d/%d answers correct", c, len(problems))
}

func ReadProblems() [][]string {
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatalf("Failed to open problems CSV file %s", err)
	}

	defer file.Close()

	problemsRecords := csv.NewReader(file)
	problems, err := problemsRecords.ReadAll()
	return problems
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	s := make([]string, 0)
	scanner.Scan()

	s = append(s, scanner.Text())

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan input %s", err)
	}
	return strings.Join(s, "")
}
