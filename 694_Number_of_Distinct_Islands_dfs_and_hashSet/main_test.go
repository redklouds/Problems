package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:

	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
	exp := 1
	if act := numDistinctIslands(grid); act != exp {
		t.Error()
	}
}
