package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// askQuestion prompts the user with the question q and checks against expected.
// Returns true if the user's answer matches expected exactly after trimming spaces.
func askQuestion(q, expected string) bool {
	var answer string
	fmt.Printf("%s = ", q)
	fmt.Scanln(&answer)
	return strings.TrimSpace(answer) == strings.TrimSpace(expected)
}

func main() {
	// Open the CSV file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file
	r := csv.NewReader(f)

	// Read all records
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed to read csv: %v", err)
	}

	// Ask only the first question and track correctness
	if len(records) == 0 {
		log.Println("no questions found in CSV")
		return
	}
	first := records[0]
	if len(first) < 2 {
		log.Println("first record malformed, need question and answer")
		return
	}

	correctCount := 0
	incorrectCount := 0

	if askQuestion(first[0], first[1]) {
		fmt.Println("Correct!")
		correctCount++
	} else {
		fmt.Printf("Incorrect. Correct answer is %s\n", first[1])
		incorrectCount++
	}

	fmt.Printf("Score: %d correct, %d incorrect\n", correctCount, incorrectCount)
}
