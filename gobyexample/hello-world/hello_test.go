package main

import "testing"

func TestHello(t *testing.T) {
    got := hello()
    want := "Hello from hello function!"
    if got != want {
        t.Errorf("hello() = %q; want %q", got, want)
    }
}
