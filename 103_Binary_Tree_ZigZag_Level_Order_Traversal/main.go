package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


//O(n) runtime
//O(n) memory
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	return BFSZigZag(root)
}

func BFSZigZag(root *TreeNode) [][]int {
	currentLevel := 1
	q := []*TreeNode{root}
	var results [][]int
	for len(q) != 0 {
		qLen := len(q)
		resultSet := make([]int, qLen)

		for i := 0; i < qLen; i++ {
			//backwards or forwards,
			if currentLevel%2 == 0 {
				//go backwards
				//BOUNDARIES FOR GETTING JUST the parition of the queue, here we want QLEN instead of len(q)
				resultSet[i] = q[qLen-1-i].Val //this will get the last element placed in the front
			} else {
				resultSet[i] = q[i].Val
			}

			//insert the same way...
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		currentLevel++ //refactor i can actually alternate the levels
		//finished processing the curent level we can add the result set to the results
		results = append(results, resultSet)
		//trim out queue since we are finished with this level
		q = q[qLen:] // [inclusiveIDX: excludedIDX], since it will not exclude the last len perfect for us
	}

	return results
}
