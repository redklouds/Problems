package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	exp := 1
	if act := shortestBridge(grid); act != exp {
		t.Error()
	}
}

func TestMain2(t *testing.T) {

	grid := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}
	exp := 2
	if actual := shortestBridge(grid); actual != exp {
		t.Error()
	}
}

func TestMain3(t *testing.T) {

	grid := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	exp := 1
	if actual := shortestBridge(grid); actual != exp {
		t.Error()
	}
}

func TestDistanceFormula(t *testing.T) {

	exp := 1
	c1 := Coords{0, 1}
	c2 := Coords{1, 0}
	if act := distance(c1, c2); act != exp {
		t.Error()
	}

	exp = 3
	c1 = Coords{0, 0}
	c2 = Coords{2, 2}
	if act := distance(c1, c2); act != exp {
		t.Error()
	}
}

func TestAbs(t *testing.T) {
	exp := 1
	if act := Abs(-1); act != exp {
		t.Error()
	}

	exp = 111
	if act := Abs(-111); act != exp {
		t.Error()
	}

	exp = 68
	if act := Abs(-68); act != exp {
		t.Error()
	}

	exp = 68
	if act := Abs(68); act != exp {
		t.Error()
	}
}

func addToArray(arr *[]int, i int) {
	//*arr = append(*arr, i)
	if i != 20 {
		*arr = append(*arr, i)
		addToArray(arr, i+1)

	}

}
