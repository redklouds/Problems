package main

import "testing"

func TestMain(t *testing.T) {

	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	expected := 1

	if actual := numIslands(grid); actual != expected {
		t.Error()
	}

}
