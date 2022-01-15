package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	//for each list insert them //EASIEST LOL
	if len(lists) == 1 {
		return lists[0]
	}
	resultHead := lists[0] //take the first list as a default list
	for i := 1; i < len(lists); i++ {
		ptr := lists[i] //the current poitner for each link list in the array
		for ptr != nil {
			//while ptr still has something insert it
			resultHead = InsertNode(resultHead, *ptr)
			ptr = ptr.Next
		}

	}
	return resultHead
}

func InsertNode(head *ListNode, toInsert ListNode) *ListNode {
	if head.Val > toInsert.Val {
		//head [4]->[6], To Insert is [2]

		//toInsert.Next = head //[2] -> [4]
		//head = toInsert
		temp := head
		head = &toInsert
		head.Next = temp
	} else {
		//somewhere in the middle
		ptr := head
		for ptr.Next != nil && toInsert.Val > ptr.Next.Val { //while current.Next is greater than the Inserted node
			//continue until target hits a next that's larger than itself
			//OR ptr.Next is nil meaning we are at the end
			//ptr.Next will be the inserted val
			ptr = ptr.Next
		}
		//part our new node here
		toInsert.Next = ptr.Next
		ptr.Next = &toInsert
	}
	return head
}
func main() {

}

func PrintList(node *ListNode) {
	ptr := node
	for ptr != nil {
		fmt.Printf("[%v]->", ptr.Val)
		ptr = ptr.Next
	}
}
