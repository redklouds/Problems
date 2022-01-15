package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow := head
	fast := head
	front := head
	for i := 1; i < k; i++ {
		if fast.Next == nil {
			if i == k-1 {
				head = head.Next
			}
			return head

		}
		fast = fast.Next
		//slow = slow.Next
		front = front.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if slow != nil {
		//the slow is now nth from the last
		temp := front.Val
		front.Val = slow.Val
		slow.Val = temp

		//we want to take SLOW.Next = to the front pointer
		//and front.Next = to the slow
		//slow.Next = slow.Next.Next
	}
	return head

}
