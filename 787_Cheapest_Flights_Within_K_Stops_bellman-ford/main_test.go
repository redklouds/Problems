package main

import "testing"

func TestMain(t *testing.T) {

	n := 3
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 1
	//TODO:

	expected := 200
	if actual := findCheapestPrice(n, flights, src, dst, k); actual != expected {
		t.Error()
	}
}
