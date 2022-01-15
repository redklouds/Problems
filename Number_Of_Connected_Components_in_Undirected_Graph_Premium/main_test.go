package main

import "testing"

//returns 2, since there exists 2 disconnected components in graph
func TestNumConnectedComp_Return_2(t *testing.T) {

	//2 components exists here
	testEdges := [][]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	n := 5
	Actual := NumConnectedCompUnDirectedGraph(n, testEdges)

	Expected := 2
	if Expected != Actual {
		t.Errorf("Expected: %d, Got: %d", Expected, Actual)
	}
}

//returns 1 since there exists only length of 1 in edges single edge
func TestNumConnectedComp_One_Edge_Return_1(t *testing.T) {
	testEdges := [][]int{
		{0, 1},
	}
	n := 2

	Expected := 1
	if Actual := NumConnectedCompUnDirectedGraph(n, testEdges); Expected != Actual {
		t.Errorf("Expected: %d, Got: %d", Expected, Actual)
	}
}

//returns 0 since numNodes == 0 or length == 0
func TestNumConnectedComp_zero_numNodes(t *testing.T) {

	testEdges := [][]int{
		{},
	}
	n := 0
	Actual := NumConnectedCompUnDirectedGraph(n, testEdges)
	Expected := 0
	if Expected != Actual {
		t.Errorf("Expected: %d, Got: %d", Expected, Actual)
	}
}

//returns number of nodes, if there contains ZERO edges,
//that is because if there exists no edges there are as many components
//as there is edges
func TestNumConnectedComp_zero_edges(t *testing.T) {

	testEdges := [][]int{}
	n := 10
	Actual := NumConnectedCompUnDirectedGraph(n, testEdges)
	Expected := 10
	if Expected != Actual {
		t.Errorf("Expected: %d, Got: %d", Expected, Actual)
	}
}
