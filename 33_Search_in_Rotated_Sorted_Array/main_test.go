package main

import "testing"

func TestMain(t *testing.T) {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4

	if actual := search(nums, target); actual != expected {
		t.Error()
	}

}

func TestMain2(t *testing.T) {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1

	if actual := search(nums, target); actual != expected {
		t.Error()
	}

}

func TestMain3(t *testing.T) {

	nums := []int{1}
	target := 0
	expected := -1

	if actual := search(nums, target); actual != expected {
		t.Error()
	}

}

func TestMain4(t *testing.T) {
	nums := []int{1}
	target := 1
	expected := 0

	if actual := search(nums, target); actual != expected {
		t.Error()
	}
}

func TestMain5(t *testing.T) {
	nums := []int{1, 3}
	target := 3
	expected := 1
	if actual := search(nums, target); actual != expected {
		t.Error()
	}
}


func TestMain6(t *testing.T) {
	nums := []int{3,5,1}
	target := 3
	expected := 0
	if actual := search(nums, target); actual != expected {
		t.Error()
	}
}

