package main

import "testing"

func TestMain(t *testing.T) {
	testData := [][]int{{1, 0}}
	expected := true

	if actual := canFinish(2, testData); expected != actual {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	testData := [][]int{{1, 0}, {0, 1}}
	expected := false

	if actual := canFinish(2, testData); expected != actual {
		t.Error()
	}
}


func TestMain3(t *testing.T) {
	testData := [][]int{{2,0}, {1,0}, {3,1}, {3,2}, {1,3}}
	expected := false

	if actual := canFinish(4, testData); expected != actual {
		t.Error()
	}
}
