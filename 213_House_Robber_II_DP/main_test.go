package main

import "testing"

func TestHouseRobberII_DP(t *testing.T) {
	testHouseStash := []int{1, 2, 3, 1}

	expected := 4

	if actual := rob(testHouseStash); actual != expected {
		t.Errorf("Expected: %d, Got: %d", expected, actual)
	}
}
