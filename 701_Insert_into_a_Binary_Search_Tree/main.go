package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return InsertBST(root, val)

}

func InsertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		//go left, since the val is less than current value
		root.Left = InsertBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = InsertBST(root.Right, val)
	}

	return root
}
