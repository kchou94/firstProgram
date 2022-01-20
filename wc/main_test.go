package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("The quick brown fox jumps over the lazy dog!\n")
	exp := 9
	res := count(b)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
