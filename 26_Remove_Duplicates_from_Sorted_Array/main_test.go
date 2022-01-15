package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:
	nums := []int{1, 1, 2}
	Expected := 2
	if actual := removeDuplicates(nums); actual != Expected {
		t.Error()
	}

}

func TestMain2(t *testing.T) {
	//TODO:
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	Expected := 5
	if actual := removeDuplicates(nums); actual != Expected {
		t.Error()
	}

}
