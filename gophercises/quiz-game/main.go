package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
	// Flags
	limit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

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

	if len(records) == 0 {
		log.Println("no questions found in CSV")
		return
	}

	correctCount := 0
	incorrectCount := 0

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, rec := range records {
		if len(rec) < 2 {
			log.Printf("skipping malformed record at line %d: %#v", i+1, rec)
			continue
		}

		answerCh := make(chan bool)
		go func(q, ans string) {
			answerCh <- askQuestion(q, ans)
		}(rec[0], rec[1])

		select {
		case correct := <-answerCh:
			if correct {
				fmt.Println("Correct!")
				correctCount++
			} else {
				fmt.Printf("Incorrect. Correct answer is %s\n", rec[1])
				incorrectCount++
			}
		case <-timer.C:
			fmt.Println()
			fmt.Println("Time's up!")
			fmt.Printf("Final Score: %d correct, %d incorrect (Answered: %d / %d)\n", correctCount, incorrectCount, correctCount+incorrectCount, len(records))
			return
		}
	}

	fmt.Printf("Final Score: %d correct, %d incorrect (Total: %d)\n", correctCount, incorrectCount, correctCount+incorrectCount)
}
