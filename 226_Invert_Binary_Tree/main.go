package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	//just THINK, you just want to swap you NEED to store the
	//reference of the newly returned pointers so you don't over
	//write them so just return the respective sides
	if root == nil {
		return root
	}
	newLeft := invertTree(root.Right)
	newRight := invertTree(root.Left)

	root.Left = newLeft
	root.Right = newRight

	return root
}
