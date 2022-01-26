package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:

	grid := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'},
	}

	exp := 3
	if actual := getFood(grid); exp != actual {
		t.Errorf("exp %d Got %d", exp, actual)
	}
}

func TestMain2(t *testing.T) {
	//TODO:

	grid := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	}

	exp := 6
	if actual := getFood(grid); exp != actual {
		t.Errorf("exp %d Got %d", exp, actual)
	}
}

func TestMain3(t *testing.T) {
	//TODO:

	grid := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'X', 'O', 'X'},
		{'X', 'O', 'X', '#', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}

	exp := -1
	if actual := getFood(grid); exp != actual {
		t.Errorf("exp %d Got %d", exp, actual)
	}
}
