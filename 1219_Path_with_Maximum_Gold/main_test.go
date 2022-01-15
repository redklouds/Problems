package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:

	grid := [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}
	expected := 24
	/*
	   Explanation:
	   [[0,6,0],
	    [5,8,7],
	    [0,9,0]]
	*/
	if actual := getMaximumGold(grid); actual != expected {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	//TODO:

	grid := [][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	}
	expected := 28
	/*
		Explanation:
		[[1,0,7],
		 [2,0,6],
		 [3,4,5],
		 [0,3,0],
		 [9,0,20]]
	*/
	if actual := getMaximumGold(grid); actual != expected {
		t.Error()
	}
}
