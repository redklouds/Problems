package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:

	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	/*
		root2 := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		}
	*/
	if actual := insertIntoBST(root, 5); !validateBSTTree(actual, root) {
		t.Error()
	}
	fmt.Println("actual")

}

func validateBSTTree(root1 *TreeNode, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	if (root1.Left == nil && root2.Left != nil) || (root1.Left != nil && root2.Left == nil) {
		return false
	}
	if (root1.Right == nil && root2.Right != nil) || (root1.Right != nil && root2.Right == nil) {
		return false
	}

	if !validateBSTTree(root1.Left, root2.Left) || !validateBSTTree(root1.Right, root2.Right) {
		return false
	}
	return true
}
