package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:

	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}
	expected := 42
	if actual := maxPathSum(root); expected != actual {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	//TODO:

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
	}
	expected := 6
	if actual := maxPathSum(root); expected != actual {
		t.Error()
	}
}

func TestMain3(t *testing.T) {
	//TODO:

	root := &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}

	expected := -3
	if actual := maxPathSum(root); expected != actual {
		t.Errorf("expected %d Got %d", expected, actual)

	}
}
