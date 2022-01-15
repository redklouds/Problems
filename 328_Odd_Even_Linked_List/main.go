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
func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	nodeIdx := 1

	oddPtr := head
	evenPtr := head.Next

	if evenPtr == nil {
		return head
	}
	currPtr := evenPtr.Next

	evenPtrDummy := evenPtr
	for currPtr != nil {
		if nodeIdx%2 == 0 {
			//its even
			evenPtr.Next = currPtr
			evenPtr = evenPtr.Next
		} else {
			//its odd
			oddPtr.Next = currPtr
			oddPtr = oddPtr.Next
		}
		currPtr = currPtr.Next
	}

	//now we finished
	oddPtr.Next = evenPtrDummy
	return head
}
