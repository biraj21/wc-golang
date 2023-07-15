package main

import (
	"testing"
)

var numLines, numWords, numBytes, numChars = GetStats("test.txt")

func TestNumLines(t *testing.T) {
	want := int64(7137)
	if numLines != want {
		t.Fatalf("want %d, got %d", want, numLines)
	}
}

func TestNumWords(t *testing.T) {
	want := int64(58159)
	if numWords != want {
		t.Fatalf("want %d, got %d", want, numWords)
	}
}
func TestNumChars(t *testing.T) {
	want := int64(339120)
	if numChars != want {
		t.Fatalf("want %d, got %d", want, numChars)
	}
}
func TestNumBytes(t *testing.T) {
	want := int64(341836)
	if numBytes != want {
		t.Fatalf("want %d, got %d", want, numBytes)
	}
}
