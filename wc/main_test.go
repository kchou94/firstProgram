package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("The quick brown fox jumps over the lazy dog!\n")
	exp := 9
	res := count(b, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("The quick brown\nfox jumps over the lazy\ndog!")
	exp := 3
	res := count(b, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
