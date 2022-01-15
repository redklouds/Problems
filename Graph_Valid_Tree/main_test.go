package main

import "testing"

func TestIsGraphValid_With_Loop_Return_False(t *testing.T) {
	//return false, signifying that this is not a valid tree graph

	testEdges := [][]int{
		{1, 3},
		{0, 2},
		{2, 4},
		{3, 4},
	}
	numNodes := 4
	Expected := false

	Actual := IsGraphValidTree(numNodes, testEdges)
	if Actual != Expected {
		t.Errorf("Expected: %t, Got: %t", Expected, Actual)
	}
}

func TestIsGraphValid_NoLoop_Return_True(t *testing.T) {
	testEdges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	testEdges2 := [][]int{
		{0, 1},
		{2, 0},
		{3, 0},
		{4, 3},
	}

	//first testCase
	n := 4
	Expected := true

	Actual := IsGraphValidTree(n, testEdges)
	if Expected != Actual {
		t.Errorf("Expected: %t, Got: %t", Expected, Actual)
	}

	n = 5
	Expected = true
	if Actual = IsGraphValidTree(n, testEdges2); Actual != Expected {
		t.Errorf("Expected: %t, Got: %t", Expected, Actual)
	}
}
