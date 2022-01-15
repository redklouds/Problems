package main

import "testing"

func TestFindInRotatedSortedArray(t *testing.T) {
	testData := []int{4, 5, 6, 7, 0, 1, 2}
	expected := 0
	target := 0
	if actual := Search(testData, target); actual != expected {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}
