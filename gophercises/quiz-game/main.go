package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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

	// Print each record
	for i, rec := range records {
		// Expecting two columns: question and answer
		if len(rec) < 2 {
			log.Printf("skipping malformed record at line %d: %#v", i+1, rec)
			continue
		}
		fmt.Printf("Q%d: %s = %s\n", i+1, rec[0], rec[1])
	}
}
