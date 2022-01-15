package main

import "testing"

func TestMain(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if actual := spiralOrder(matrix); !Equals(actual, expected) {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	matrix := [][]int{{3}, {2}}
	expected := []int{3,2}
	if actual := spiralOrder(matrix); !Equals(actual, expected) {
		t.Error()
	}
}


func TestMain3(t *testing.T) {
	matrix := [][]int{{3,3,4}}
	expected := []int{3,3,4}
	if actual := spiralOrder(matrix); !Equals(actual, expected) {
		t.Error()
	}
}

func TestMain4(t *testing.T) {
	matrix := [][]int{{3,3,4}, {5,4,2}}
	expected := []int{3,3,4,2,4,5}
	if actual := spiralOrder(matrix); !Equals(actual, expected) {
		t.Error()
	}
}

func TestMain5(t *testing.T) {
	matrix := [][]int{{3}}
	expected := []int{3}
	if actual := spiralOrder(matrix); !Equals(actual, expected) {
		t.Error()
	}
}

func Equals(a, b []int) bool {
	if len(a) != len(b){
		return false
	}

	for i := range a{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}
