package main

import "testing"

func TestPathCompression(t *testing.T) {
	edgeList := [][]int{

		{0, 1},
		{0, 3},
		{2, 3},
		{1, 2},
	}
	numNodes := 4

	djPathCompression := DisjoinsetPathCompression{}
	djPathCompression.MakeSet(numNodes, edgeList)

	expected := []Subset{
		{
			Parent: 1,
			Rank:   0,
		},
		{
			Parent: 3,
			Rank:   0,
		},
		{
			Parent: 3,
			Rank:   0,
		},
		{
			Parent: 3,
			Rank:   0,
		},
	}
	for idx := range djPathCompression.Arr {
		if expected[idx].Parent != djPathCompression.Arr[idx].Parent {
			t.Errorf("\nExpected: %d, Got: %d", expected[idx], djPathCompression.Arr[idx].Parent)
		}
		if expected[idx].Rank != djPathCompression.Arr[idx].Rank {
			t.Errorf("\nExpected: %d, Got: %d", expected[idx].Rank, djPathCompression.Arr[idx].Rank)
		}
	}
}
