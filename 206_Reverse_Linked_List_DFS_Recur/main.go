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
func reverseList(head *ListNode) *ListNode {
	//insert at the tail LOL ez
	//go all the way down.. now say that the next equals the function
	if head == nil {
		return nil
	}
	return recur(head)

}

func recur(node *ListNode) *ListNode {

	if node.Next == nil {
		return node
	}

	newHead := recur(node.Next) //we want to fetch the new head
	//REMEMBER THAT AT EACH RECURSIVE STACK, IN THE CURRENT stack we are holding two peice of info
	//the CURRENT HEAD HERE and returning the last known head
	node.Next.Next = node //if we were at 3, 3.Next Is 4 and 4.Next is
	node.Next = nil

	return newHead //WE LITERALLY BRING the last node ALL the way back Notice that we do not modify the returl statement at all, the same object is returned at each recursive step
}
