package main

import "testing"

func TestMain(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	exp := 2
	if actual := countComponents(n, edges); actual != exp {
		t.Error()
	}


	n = 5
	edges = [][]int{{0, 1}, {1, 2},{2,3}, {3, 4}}
	exp = 1
	if actual := countComponents(n, edges); actual != exp {
		t.Error()
	}

	n = 1
	edges = [][]int{{}}
	exp = 1
	if actual := countComponents(n, edges); actual != exp {
		t.Error()
	}
}
