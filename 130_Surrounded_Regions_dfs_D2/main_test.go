package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)

	fmt.Printf("%c", board)
	for row := range board {
		for col := range board[0] {
			if board[row][col] != expected[row][col] {
				t.Error()
			}
		}
	}

}
