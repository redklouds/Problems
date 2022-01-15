package main

import "testing"

func TestMain1(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}

	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
	}

	expected := 3
	if actual := countSubIslands(grid1, grid2); actual != expected {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	grid1 := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}

	grid2 := [][]int{
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}

	expected := 0
	if actual := countSubIslands(grid1, grid2); actual != expected {
		t.Error()
	}

}

func TestMain3(t *testing.T) {
	g1 := [][]int{
		{1, 1, 1, 1, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 0}}
	g2 := [][]int{
		{1, 1, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 0}}
	expected := 0
	if actual := countSubIslands(g1, g2); actual != expected {
		t.Error()
	}

}
