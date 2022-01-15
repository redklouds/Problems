package main

import "fmt"
type ListNode struct {
	Data int
	Next *ListNode
}


func MergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	ptr := dummyHead
	for list1 != nil && list2 != nil {
	
			if list1.Data <= list2.Data{
				//choose list 1
				ptr.Next = list1 // pointer poiting to another Poiinter? needs to be dereferenced to thta tactual object
				list1 = list1.Next
			}else{
				ptr.Next = list2
				list2 = list2.Next
			}
		ptr = ptr.Next

	}

				if list1 != nil {
				ptr.Next = list1
				list1 = list1.Next
			}
			if list2 != nil {
				ptr.Next = list2
				list2 = list2.Next
			}
	return dummyHead.Next

}
func main() {

	/*
	//these are ACTUAL objects to LISTNODES, to get a pointer to them
	//get the addresspointer (&) if you have a pointer and want the actual object use (*)
	l1 := ListNode{1, nil} //literal, using name space, if you don't want you can add all params
	l2 := ListNode{Data: 2}
	l3 := ListNode{Data: 3}
	l4 := ListNode{Data: 4}
	l5 := ListNode{Data: 5}

	l6 := ListNode{Data: 6}
	l7 := ListNode{Data: 7}

	var list1 *ListNode
	//list1 [1] -> [3] -> [5]
	list1 = &l1
	l1.Next = &l3
	l3.Next = &l5
	//list2 [2] -> [4] -> [6] -> [7]

	var list2 *ListNode

	list2 = &l2
	l2.Next = &l4

	l4.Next = &l6
	l6.Data = &l7
	l6.Data = &l7

	res := MergeList(list)
*/

	l1 := &ListNode{Data:10}
	l2 := &ListNode{Data: 5}
	if l1.Data > l2.Data{
		fmt.Printf("L1 is larger\n")
	}

	


}
