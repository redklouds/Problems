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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	sortedList := sortedMerge(list1, list2)
	//one of the list are nil find out which and append to end of dummy
	return sortedList
}

//using recursion we can perform this in O(n) timne

//takes two SORTED linked list and returnes a single merged sorted ascending order linked list
func sortedMerge(list1, list2 *ListNode) *ListNode {
	//base case, if either list are nil we return the opposite and append it to the end of the list that is not nil

	//base case if list1 is nil it returns list 2,
	//if list2 is nil it returns list1
	//if both list1 is nil and list 2is nil it return the other, which is nil, nice
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	//it is NOt a list1 else list2 retunr, becausse we only want to do tis base case if one of the list are nil

	//now we will have a pointer to the current lower node
	var lowerNodePtr *ListNode

	//now check which of the two list are the smaller value

	if list1.Val <= list2.Val {
		//list1 is lower point to it as reference to return
		lowerNodePtr = list1
		lowerNodePtr.Next = sortedMerge(list1.Next, list2)
	} else {
		//list2 is lower
		lowerNodePtr = list2
		lowerNodePtr.Next = sortedMerge(list1, list2.Next)
	}
	return lowerNodePtr
}
