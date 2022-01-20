package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//O(n)
//a function that reverses a linked list
func ReverseLL(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	newHead := ReverseLL(head.Next)
	//THINK: its like [LEFT] -> [RIGHT] -> NILL
	//we want to move the LEFT to the RIGHT SIDE so its like [RIGHT] -> [LEFT] to do this
	//you can honestly say that [LEFT].Next.Next = [LEFT] AKA [LEFT] ->(next)[RIGHT] ->(next)= LEFT
	head.Next.Next = head //we move the current node to its right side
	//now we know that the current head node STILL HAS a reference its next to the next node
	//we have to nil it out
	head.Next = nil // [HEAD].Next = nil
	return newHead
}

func reorderList(head *ListNode) {
	//revesre the second half and then stitch the list back together, alternating

	ptr1 := head
	ptr2 := head

	for ptr2 != nil && ptr2.Next != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next.Next
	}
	p := ReverseLL(ptr1)

	ptr2 = head

	for p.Next != nil && ptr2 != nil {
		//eventually; the LEFT pointer (first Partition, will meet the second one
		//THINK of this [1]->[2]->[3]->[5]->[4]->[3](repeated))
		//BECAUSE WE DIDN'T RELEAVE THE REFERENCE OF THE SECOND PARTITION [4] will POINT back to the same THREE
		//that we onced used
		if p == ptr2 {
			break
		}
		//we have to hold the next nodes so we can link it to the next node later
		//[1]->[5] ->[2](Store this node)
		temp := ptr2.Next
		ptr2.Next = p         //Link the current FIrst Parition to the second list
		p = p.Next            //move the second poitner up FIRST IMPORTANT,
		ptr2.Next.Next = temp //Link back to the parition of its current list

		ptr2 = ptr2.Next.Next //essentially "skipping two forward"
	}

}

func PrintLL(head *ListNode) {
	for head != nil {
		fmt.Printf("[%d]->", head.Val)
		head = head.Next
	}
}
