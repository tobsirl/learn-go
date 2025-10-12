package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

// slowReader returns first answer quickly then blocks long enough to trigger timeout.
type slowReader struct {
    reads int
}

func (s *slowReader) Read(p []byte) (int, error) {
    if s.reads == 0 {
        copy(p, []byte("2\n"))
        s.reads++
        return 2, nil
    }
    time.Sleep(50 * time.Millisecond)
    return 0, io.EOF
}

// helper to run askQuestion with simulated user input
func runAskQuestion(t *testing.T, question, expected, userInput string) bool {
    t.Helper()
    in := strings.NewReader(userInput + "\n")
    out := &bytes.Buffer{}
    return askQuestion(question, expected, in, out)
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
        in := strings.NewReader("10\n")
        out := io.Discard
        _ = askQuestion("5+5", "10", in, out)
    }
}

func TestRunQuizAllCorrect(t *testing.T) {
    records := [][]string{{"1+1", "2"}, {"2+3", "5"}}
    // Provide answers quickly
    in := strings.NewReader("2\n5\n")
    out := &bytes.Buffer{}
    correct, incorrect, timedOut := runQuiz(records, 2*time.Second, in, out)
    if timedOut {
        t.Fatalf("did not expect timeout")
    }
    if correct != 2 || incorrect != 0 {
        t.Fatalf("expected 2 correct 0 incorrect, got %d %d", correct, incorrect)
    }
}

func TestRunQuizTimeout(t *testing.T) {
    records := [][]string{{"1+1", "2"}, {"2+3", "5"}}
    // Immediate timeout: timer channel already has a value before any answers processed.
    in := strings.NewReader("2\n5\n")
    out := &bytes.Buffer{}
    timerCh := make(chan time.Time, 1)
    // Fire immediately
    timerCh <- time.Now()
    correct, incorrect, timedOut := runQuizWithTimer(records, in, out, timerCh)
    if !timedOut {
        t.Fatalf("expected timeout")
    }
    if correct != 0 || incorrect != 0 {
        t.Fatalf("expected 0 correct 0 incorrect for immediate timeout, got %d %d", correct, incorrect)
    }
}

func TestRunQuizSkipsMalformed(t *testing.T) {
    records := [][]string{{"1+1", "2"}, {"bad"}, {"2+2", "4"}}
    in := strings.NewReader("2\n4\n")
    out := &bytes.Buffer{}
    correct, incorrect, timedOut := runQuiz(records, time.Second, in, out)
    if timedOut {
        t.Fatalf("did not expect timeout")
    }
    if correct != 2 || incorrect != 0 {
        t.Fatalf("expected 2 correct 0 incorrect, got %d %d", correct, incorrect)
    }
    if !strings.Contains(out.String(), "skipping malformed record") {
        t.Fatalf("expected malformed record message in output; got: %s", out.String())
    }
}
