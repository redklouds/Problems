package main

import (
	"sort"
	"testing"
)


//U(n^2)  even though we save space, we can use a hash to find the two some
//we save space but we also have to iterate thhrough the values at each interval 
//then again at each intervals interval
func Test3Sum(t *testing.T) {
	testNums := []int{-1, 0, 1, 2, -1, -4}
	//{-4,-1,-1,0,1,2}
	sort.Ints(testNums)

	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if actual := ThreeSome(testNums); !Compare(actual, expected) {
		t.Errorf("Got: %v, expected: %v", actual, expected)
	}
}

func Compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for row := range a {
		for col := range b {
			if a[row][col] != b[row][col] {
				return false
			}
		}
	}

	return true
}
