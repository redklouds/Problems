package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var x int
	testData := []int{2, 7, 11, 15}
	targetVal := 9
	expRes := []int{1, 2}
	result := twoSum(testData, targetVal)
	if !Equal(result, expRes) {
		t.Errorf("Got: %+v Expected: %+v", result, expRes)
	}
	x = 0
	fmt.Println(x)
}
