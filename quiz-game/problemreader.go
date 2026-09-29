package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// Read the problems line by line
// Store them in a slice
// Iterate one by one over the slice asking the user the question and comparing user answer with the answer in the problem
// Keep a counter of correct responses and return that

// TODO: Need to refactor to use a global timer instead

func main() {
	problems := ReadProblems()

	fmt.Printf("Welcome to the quiz game! There are a total of %d problems to solve\nEach problem gives you 3s to solve before marking it as incorrect\n", len(problems))

	c := 0
	i := make(chan string, 1)
	go userInput(i)

	for index, problem := range problems {
		select {
		case <-i:
		default:
		}

		t := time.NewTimer(3 * time.Second)

		fmt.Printf("Problem #%d: %s\n", index+1, problem[0])

		userSolution := waitOrInput(i, t.C)

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

func userInput(c chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %s", err)
	}
}

func waitOrInput(i chan string, t <-chan time.Time) string {
	select {
	case <-t:
		return ""
	case input := <-i:
		return input
	}
}
