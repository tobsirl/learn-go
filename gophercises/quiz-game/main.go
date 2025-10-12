package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// askQuestion prompts the user with the question q and checks against expected using provided IO.
// Returns true if the user's answer matches expected exactly after trimming spaces.
func askQuestion(q, expected string, in io.Reader, out io.Writer) bool {
	// We read a single token (line). Use fmt.Fscan to read until whitespace; acceptable for numeric answers.
	fmt.Fprintf(out, "%s = ", q)
	var answer string
	fmt.Fscan(in, &answer)
	return strings.TrimSpace(answer) == strings.TrimSpace(expected)
}

// runQuiz runs through the provided records with a time limit, reading answers from input and writing prompts/results to output.
// Returns counts of correct/incorrect and whether the quiz timed out before completion.
func runQuiz(records [][]string, limit time.Duration, in io.Reader, out io.Writer) (correct, incorrect int, timedOut bool) {
	timer := time.NewTimer(limit)
	defer timer.Stop()
	return runQuizWithTimer(records, in, out, timer.C)
}

// runQuizWithTimer allows injecting a timer channel for tests.
func runQuizWithTimer(records [][]string, in io.Reader, out io.Writer, timerCh <-chan time.Time) (correct, incorrect int, timedOut bool) {
	if len(records) == 0 {
		fmt.Fprintln(out, "no questions found")
		return 0, 0, false
	}
	for i, rec := range records {
		if len(rec) < 2 {
			fmt.Fprintf(out, "skipping malformed record at line %d: %#v\n", i+1, rec)
			continue
		}
		answerCh := make(chan bool)
		go func(q, ans string) {
			answerCh <- askQuestion(q, ans, in, out)
		}(rec[0], rec[1])

		select {
		case correctAns := <-answerCh:
			if correctAns {
				fmt.Fprintln(out, "Correct!")
				correct++
			} else {
				fmt.Fprintf(out, "Incorrect. Correct answer is %s\n", rec[1])
				incorrect++
			}
		case <-timerCh:
			fmt.Fprintln(out)
			fmt.Fprintln(out, "Time's up!")
			fmt.Fprintf(out, "Final Score: %d correct, %d incorrect (Answered: %d / %d)\n", correct, incorrect, correct+incorrect, len(records))
			return correct, incorrect, true
		}
	}
	fmt.Fprintf(out, "Final Score: %d correct, %d incorrect (Total: %d)\n", correct, incorrect, correct+incorrect)
	return correct, incorrect, false
}

func main() {
    // Flags
    limit := flag.Int("limit", 30, "time limit for the quiz in seconds")
    flag.Parse()

    f, err := os.Open("problems.csv")
    if err != nil {
        log.Fatalf("failed to open file: %v", err)
    }
    defer f.Close()

    r := csv.NewReader(f)
    records, err := r.ReadAll()
    if err != nil {
        log.Fatalf("failed to read csv: %v", err)
    }

    runQuiz(records, time.Duration(*limit)*time.Second, os.Stdin, os.Stdout)
}
