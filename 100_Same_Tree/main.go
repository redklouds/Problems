package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//o(n) since we only visit each node once
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//think conditions for not being same
	//if the left is nill and the right is not
	//if the right is nil and the left is not
	//perform inorder traversale
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	//what happens to the stack call when both are nil? we want to exit and return true
	if p == nil && q == nil {
		return true
	}

	//else we know both are not NIL and have values we can continue to recurse
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
