package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	got := greeting()
	expected := "cake"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
