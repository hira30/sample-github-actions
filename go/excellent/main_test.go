package main

import "testing"

func TestEvenOrAdd(t *testing.T) {
	result := EvenOrAdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}