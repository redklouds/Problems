package main

import "testing"

func TestUniquePaths_BruteForce(t *testing.T) {

}

func TestUniquePaths_DP(t *testing.T) {
	testRow := 3
	testCol := 2
	expected := 3

	if actual := UniquePaths(testRow, testCol); actual != expected {
		t.Errorf("Expected: %d, got %d", expected, actual)
	}
}
