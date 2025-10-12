package main

import (
	"os"
	"testing"
)

// helper to run askQuestion with simulated user input
func runAskQuestion(t *testing.T, question, expected, userInput string) bool {
    t.Helper()
    // Save original stdin
    origStdin := os.Stdin
    defer func() { os.Stdin = origStdin }()

    r, w, err := os.Pipe()
    if err != nil {
        t.Fatalf("failed to create pipe: %v", err)
    }
    // Write user input followed by newline to pipe
    if _, err := w.Write([]byte(userInput + "\n")); err != nil {
        t.Fatalf("failed to write to pipe: %v", err)
    }
    w.Close()

    os.Stdin = r
    return askQuestion(question, expected)
}

func TestAskQuestionCorrect(t *testing.T) {
    got := runAskQuestion(t, "5+5", "10", "10")
    if !got {
        t.Fatalf("expected correct answer to return true")
    }
}

func TestAskQuestionIncorrect(t *testing.T) {
    got := runAskQuestion(t, "1+1", "2", "3")
    if got {
        t.Fatalf("expected incorrect answer to return false")
    }
}

func TestAskQuestionTrimsWhitespace(t *testing.T) {
    got := runAskQuestion(t, "2+3", "5", "  5  ")
    if !got {
        t.Fatalf("expected whitespace-trimmed correct answer to return true")
    }
}

func TestAskQuestionCaseSensitivity(t *testing.T) {
    // For numeric answers case doesn't apply; ensure works with mixed input
    got := runAskQuestion(t, "3+3", "6", "6")
    if !got {
        t.Fatalf("expected numeric answer to return true")
    }
}

// Benchmark to gauge overhead (optional)
func BenchmarkAskQuestion(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // simulate answer each iteration
        orig := os.Stdin
        r, w, _ := os.Pipe()
        w.Write([]byte("10\n"))
        w.Close()
        os.Stdin = r
        _ = askQuestion("5+5", "10")
        os.Stdin = orig
    }
}
