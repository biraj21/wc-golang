package main

import (
	"testing"
)

var stats = GetStats("test.txt")

func TestNumLines(t *testing.T) {
	want := 7137
	if stats.numLines != want {
		t.Fatalf("want %d, got %d", want, stats.numLines)
	}
}

func TestNumWords(t *testing.T) {
	want := 58159
	if stats.numWords != want {
		t.Fatalf("want %d, got %d", want, stats.numWords)
	}
}

func TestNumChars(t *testing.T) {
	want := 339120
	if stats.numChars != want {
		t.Fatalf("want %d, got %d", want, stats.numChars)
	}
}

func TestNumBytes(t *testing.T) {
	want := 341836
	if stats.numBytes != want {
		t.Fatalf("want %d, got %d", want, stats.numBytes)
	}
}
