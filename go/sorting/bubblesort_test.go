package main

import (
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{2, 5, 1, 3, 4}
	expected := []int{1, 2, 3, 4, 5}

	bubblesort(arr)

	if len(expected) != len(arr) {
		t.Fatalf("Length of arrays differs: expected=%d, actual=%d", len(expected), len(arr))
	}

	for i := 1; i < len(expected); i++ {
		if expected[i] != arr[i] {
			t.Fatalf("Element on position %d doesn't match: expected=%d, actual=%d", i, expected[i], arr[i])
		}
	}
}
