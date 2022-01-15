package main

import "testing"

func TestUnionByRankNoCycle(t *testing.T) {
	edgeList := [][]int{

		{0, 1},
		{0, 3},
		{2, 3},
		{1, 2},
	}
	numNodes := 4

	djUnionByRank := DisjointSet{}
	djUnionByRank.MakeSet(numNodes, edgeList)

	expected := []Subset{
		Subset{
			Parent: 0,
			Rank:   3,
		},
		Subset{
			Parent: 0,
			Rank:   0,
		},
		Subset{
			Parent: 0,
			Rank:   0,
		},
		Subset{
			Parent: 0,
			Rank:   0,
		},
	}
	for idx := range djUnionByRank.Arr {
		if expected[idx].Parent != djUnionByRank.Arr[idx].Parent {
			t.Errorf("\nExpected: %d, Got: %d", expected[idx], djUnionByRank.Arr[idx].Parent)
		}
		if expected[idx].Rank != djUnionByRank.Arr[idx].Rank {
			t.Errorf("\nExpected: %d, Got: %d", expected[idx].Rank, djUnionByRank.Arr[idx].Rank)
		}
	}
}
