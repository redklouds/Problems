package main

import "testing"

func TestMain(t *testing.T) {

	graph := [][]int{{1, 2}, {3}, {3}, {}}
	//exp := [][]int{{0, 1, 3}, {0, 2, 3}}

	allPathsSourceTarget(graph)

	graph = [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}

	allPathsSourceTarget(graph)

	
	graph = [][]int{{}}

	allPathsSourceTarget(graph)
}

func MatEquals(mata, matb [][]int) bool {
	if len(mata) != len(matb) {
		return false
	}

	return true
}
