package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root //err return nil
	}
	if val < root.Val {
		//go to the left
		root = searchBST(root.Left, val)
	} else if val > root.Val {
		//go to thr right
		root = searchBST(root.Right, val)
	}

	return root

}
