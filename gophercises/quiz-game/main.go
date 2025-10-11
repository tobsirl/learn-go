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

	// Loop through all questions and track correctness
	if len(records) == 0 {
		log.Println("no questions found in CSV")
		return
	}

	correctCount := 0
	incorrectCount := 0

	for i, rec := range records {
		if len(rec) < 2 {
			log.Printf("skipping malformed record at line %d: %#v", i+1, rec)
			continue
		}
		if askQuestion(rec[0], rec[1]) {
			fmt.Println("Correct!")
			correctCount++
		} else {
			fmt.Printf("Incorrect. Correct answer is %s\n", rec[1])
			incorrectCount++
		}
	}

	fmt.Printf("Final Score: %d correct, %d incorrect (Total: %d)\n", correctCount, incorrectCount, correctCount+incorrectCount)
}
