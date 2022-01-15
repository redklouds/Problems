package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	exp := -1
	if actual := FindCelebrity(knows, 2); actual != exp {
		t.Error()
	}
}

func knows(a, b int) bool {

	//data := [][]int{{1, 1, 0}, {0, 1, 0}, {1, 1, 1}}
	//data := [][]int{{1, 0}, {1, 1}}
	data := [][]int{{1, 1}, {1, 1}}
	if data[a][b] == 1 {
		return true
	}
	return false
}
