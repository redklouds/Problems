package main

import "testing"

func TestMain(t *testing.T) {

	root := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Left: &TreeNode{
			Val:   9,
			Right: nil,
			Left:  nil,
		},
	}

	zigzagLevelOrder(root)

	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   5,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		},
	}

	zigzagLevelOrder(root)
}
