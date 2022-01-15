package main

import "testing"

func TestMaxProductSubArrayV3(t *testing.T) {
	testData := []int{-2, 0, -1}
	expected := 0
	if actual := MaxProductSubArrayV3(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}

	testData = []int{-4, -3, -2}
	expected = 12
	if actual := MaxProductSubArrayV3(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}
}

func TestMaxProductSubArrayV1(t *testing.T) {
	testData := []int{-2, 0, -1}
	expected := 0
	if actual := MaxProductLinearSpace(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}

	testData = []int{-1, -2, -3, 0, 3, 5}
	expected = 15
	if actual := MaxProductLinearSpace(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}
}

func TestMaxProductDPMaxArraysWithDpArray(t *testing.T) {
	testData := []int{2, -1, 1, 1}
	expected := 2
	if actual := MaxProductDPMaxArrays(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}

	testData = []int{-1, -2, -3, 0, 3, 5}
	expected = 15
	if actual := MaxProductLinearSpace(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}
}

func TestMaxProduct_Dp(t *testing.T) {
	testData := []int{-2, 3, -4}

	expected := 3
	if actual := maxProduct(testData); actual != expected {
		t.Errorf("Got: %d, expected: %d", actual, expected)
	}
}
