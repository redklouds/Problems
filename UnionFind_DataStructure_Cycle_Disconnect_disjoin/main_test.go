package main

import "testing"

func TestUnionSuccesfulOfNoSubGraph(t *testing.T) {
	//we want to make sure that our graph is organized as a UnionFind Data Structure
	// 0 1 2 3
	edgeList := [][]int{

		{0, 1},
		{0, 3},
		{2, 3},
		{1,2},
	}
	numNodes := 4

	unionFindDS := NewUnionFindDS(numNodes)
	for _, pair := range edgeList {
		unionFindDS.InsertEdge(pair[0], pair[1])
	}

	expected := []int{1, 3, 3, -1}

	for idx := range unionFindDS.Arr {
		if expected[idx] != unionFindDS.Arr[idx] {
			t.Errorf("Expected: %d, Got: %d", expected[idx], unionFindDS.Arr[idx])
		}
	}

}
