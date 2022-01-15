package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//reverses list O(n)
//
func ReverseLinkedList(head *ListNode) *ListNode {

	ptr := head
	dh := &ListNode{}
	for ptr != nil {
		temp := ptr
		ptr = ptr.Next
		temp.Next = dh.Next
		dh.Next = temp
	}
	return dh.Next
}
//you must find the center of the list and reverse  the second half

//since its second half reverse alternate between firsta nd second halves
func reorderList(head *ListNode) {
	slowPtr := head
	fastPtr := head

	//lets go to the center of the ll and reverse
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	//slow ptr is at mid lets reverse the right side
	slowPtr = ReverseLinkedList(slowPtr)
	alterNameFlag := 1

	ptr := head //need to alternate and keep track of the nodess so we don't lose them
	leftPtr := head.Next
	rightPtr := slowPtr
	for leftPtr != nil || rightPtr != nil {
		if leftPtr == rightPtr {
			ptr.Next = leftPtr
			break
		}
		if alterNameFlag > 0 {
			//take from right half
			ptr.Next = rightPtr
			rightPtr = rightPtr.Next
		} else {
			//take from left half
			ptr.Next = leftPtr
			leftPtr = leftPtr.Next
		}
		alterNameFlag *= -1
		ptr = ptr.Next
	}
}

func main() {
	//testData := []int{1,2,3,4}
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	//PrintLL(ReverseLinkedList(n1))

	reorderList(n1)
	PrintLL(n1)
	//reorderList(n1)
}

func PrintLL(head *ListNode) {
	ptr := head
	for ptr != nil {
		fmt.Printf("[%v]->", ptr.Val)
		ptr = ptr.Next
	}
}
