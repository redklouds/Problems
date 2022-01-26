package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	exp := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	if actual := updateMatrix(grid); !Equal(actual, exp) {
		t.Error()
	}

}


func TestMain2(t *testing.T) {
	//TODO:
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	exp := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}

	if actual := updateMatrix(grid); !Equal(actual, exp) {
		t.Error()
	}

}

func Equal(grid1, grid2 [][]int) bool {

	if len(grid1) != len(grid2) {
		return false
	}
	for row := range grid1 {
		for col := range grid1[0] {
			if grid1[row][col] != grid2[row][col] {
				return false
			}
		}
	}
	return true
}
