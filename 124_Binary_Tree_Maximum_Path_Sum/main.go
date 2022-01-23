package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	ret := -math.MaxInt32
	maxPathSumHelper(root, &ret)
	return ret
}

func maxPathSumHelper(root *TreeNode, ret *int) int {
	if root == nil {
		return -math.MaxInt32
	}

	left := maxPathSumHelper(root.Left, ret)
	right := maxPathSumHelper(root.Right, ret)

	SumSplit := Max(left, right, 0) + root.Val
	SumNoSplit := left + right + root.Val
	*ret = Max(*ret, SumSplit, SumNoSplit)
	return SumSplit
}

func Max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}
