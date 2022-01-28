package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:

	grid := [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 2, 2, 0},
	}

	island1 := 1
	island2 := 2
	expected := 2

	if actual := shortestDistance(grid, island1, island2); expected != actual {
		t.Error()
	}

}

func TestShortestDistance(t *testing.T) {

	coord1 := Coords{0, 1}
	coord2 := Coords{3, 1}

	expect := 2
	if act := ShortestDistance(coord1, coord2); act != expect {
		t.Error()
	}
}

func TestAbs(t *testing.T) {
	exp := 1
	if act := Abs(-1); act != exp {
		t.Error()
	}

	exp = 55
	if act := Abs(-55); act != exp {
		t.Error()
	}
	exp = 10
	if act := Abs(10); act != exp {
		t.Error()
	}

	exp = 1
	if act := Abs(-1); act != exp {
		t.Error()
	}
}
