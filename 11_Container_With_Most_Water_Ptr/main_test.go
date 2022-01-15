package main

import "testing"

func TestMaxRainWater_AREA(t *testing.T) {
	testData_heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49

	if actual := MaxRaimWater_Area(testData_heights); actual != expected {
		t.Errorf("Got: %d, Expected: %d", actual, expected)
	}

	testData_heights = []int{1, 1}
	expected = 1

	if actual := MaxRaimWater_Area(testData_heights); actual != expected {
		t.Errorf("Got: %d, Expected: %d", actual, expected)
	}

	testData_heights = []int{1, 2, 1}
	expected = 2

	if actual := MaxRaimWater_Area(testData_heights); actual != expected {
		t.Errorf("Got: %d, Expected: %d", actual, expected)
	}
}
