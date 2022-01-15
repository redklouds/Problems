package main

import "testing"

func TestMain(t *testing.T) {
	queens := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
	king := []int{0, 0}
	//expected := [][]int{{0, 1}, {1, 0}, {3, 3}}

	queensAttacktheKing(queens, king)

}
