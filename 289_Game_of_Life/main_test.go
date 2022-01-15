package main

import "testing"

func TestGameOfLife(t *testing.T) {
	testBoard := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(testBoard)

	expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
	if !CheckBoard(testBoard, expected) {
		t.Errorf("They are not the same %v", testBoard)
	}
}

func CheckBoard(board1 [][]int, board2 [][]int) bool {
	if len(board1) != len(board2) {
		return false
	}

	for r := range board1 {
		for c := range board2[0] {
			if board1[r][c] != board2[r][c] {
				return false
			}
		}
	}
	return true
}
