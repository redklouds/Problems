package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	exp := 2
	if act := findCircleNum(isConnected); exp != act {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	//TODO:
	isConnected := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	exp := 3
	if act := findCircleNum(isConnected); exp != act {
		t.Error()
	}
}
