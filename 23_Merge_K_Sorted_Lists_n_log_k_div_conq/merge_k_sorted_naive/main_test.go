package main

import (
	"fmt"
	"testing"
)

func Test_hi(t *testing.T) {
	n1 := &ListNode{Val: 5}
	n2 := &ListNode{Val: -2}
	n1 = InsertNode(n1, ListNode{Val: 20})
	n1 = InsertNode(n1, *n2)
	PrintList(n1)

	var SortedLLArr []*ListNode

	//[1,4,5]
	var list1Head *ListNode

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 4}
	l4 := &ListNode{Val: 5}
	list1Head = l1

	fmt.Println(list1Head)
	l1.Next = l2
	l2.Next = l4

	//[1,3,4],
	l5 := &ListNode{Val: 1}
	l6 := &ListNode{Val: 3}
	l7 := &ListNode{Val: 4}
	l5.Next = l6
	l6.Next = l7

	//[2,6]
	l8 := &ListNode{Val: 2}
	l9 := &ListNode{Val: 6}
	l8.Next = l9

	SortedLLArr = append(SortedLLArr, l1)
	SortedLLArr = append(SortedLLArr, l5)
	SortedLLArr = append(SortedLLArr, l8)

	mergedList := mergeKLists(SortedLLArr)

	PrintList(mergedList)
}
