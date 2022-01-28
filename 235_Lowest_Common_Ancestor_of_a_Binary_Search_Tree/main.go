package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//run time O(logN) where N is the height of the tree, since we are searching each partition
//O(1) because we do not use extra space
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	//basically when we search we want to find the subtree that p and q belong too, 

	//we know that if there is a split that mean that it is the lowest common ancestors
	//OR if we find p or q that means we are at the lowest

	//find the partition
	cur := root
	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val{
			cur = cur.Right
		}else if p.Val < cur.Val && q.Val < cur.Val{
			//search left
			cur = cur.Left
		}else{
			return cur //found a split or one of the values are equal to p ot r
		}
	}
	return cur
}
