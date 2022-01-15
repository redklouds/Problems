package main

type ListNode struct {
	Data interface{}
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next

		if slowPtr == fastPtr {
			return true
		}
	}
	return false

}
