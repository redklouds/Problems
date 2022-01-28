package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestProductOfSelf(t *testing.T) {
	testData := []int{1, 2, 3, 4}

	expected := []int{24, 12, 8, 6}

	if actual := ProductExSelf(testData); !Equal(actual, expected) {
		t.Errorf("Got: %v, Expected: %v", actual, expected)
	}
}
