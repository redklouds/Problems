package main

import "testing"

func TestMain(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	expected := true

	if actual := exist(board, word); actual != expected {
		t.Error()
	}

}

func TestMain2(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	expected := true

	if actual := exist(board, word); actual != expected {
		t.Error()
	}

}

func TestMain3(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	expected := false

	if actual := exist(board, word); actual != expected {
		t.Error()
	}

}
