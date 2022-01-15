package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if rotate(mat); !Equals(mat, expected) {
		t.Error()
	}

}

func Equals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		for j := range a[0] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
