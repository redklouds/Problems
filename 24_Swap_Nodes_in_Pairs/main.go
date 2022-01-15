package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//LOOK at the NOTE BOOK LOL
//basically you are swapping places from the left placement of the node to the right
//IE A -> B , a is on the "right" and to make A on the Left Ie A -> B -> A , you do node.Next.Next = node that puts you on the right of your neighbor
//then BREAK the cycle from current node bnack to Neighbor ie A -> B, but if toy trhink about it, going backwards, bottom up
//return a new node JUST LIKE reverse a link list recursibly you return the nodes, and remember odd values
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newNode := swapPairs(head.Next.Next)
	head.Next.Next = head
	temp := head.Next
	head.Next = newNode
	return temp
}
