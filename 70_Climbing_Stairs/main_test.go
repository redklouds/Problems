package main

import "testing"

func TestCanJump(t *testing.T) {

	expected := 5
	targetN := 4
	steps := []int{1, 2}

	if actual := climbStairsDp(steps, targetN); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
	/*
		expected := true
		arr := []int{4, 2, 3, 0, 3, 1, 2}
		start := 5
		if actual := canReach(arr, start); actual != expected {
			t.Errorf("Expected: %t, got: %t", expected, actual)
		}

		expected = true
		arr = []int{4, 2, 3, 0, 3, 1, 2}
		start = 0
		if actual := canReach(arr, start); actual != expected {
			t.Errorf("Expected: %t, got: %t", expected, actual)
		}

		expected = false
		arr = []int{3, 0, 2, 1, 2}
		start = 2
		if actual := canReach(arr, start); actual != expected {
			t.Errorf("Expected: %t, got: %t", expected, actual)
		}
	*/
}
