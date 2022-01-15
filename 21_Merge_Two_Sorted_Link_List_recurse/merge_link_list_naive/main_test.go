package main

import "testing"

func TestMergeLinkList(t *testing.T) {

	n1 := &ListNode{Data: 1}
	n2 := &ListNode{Data: 2}
	n3 := &ListNode{Data: 3}
	n4 := &ListNode{Data: 4}
	n5 := &ListNode{Data: 5}
	n6 := &ListNode{Data: 6}
	n7 := &ListNode{Data: 7}
	var list1 *ListNode
	var list2 *ListNode
	//list1 [1] -> [3] -> [5]
	n1.Next = n3
	n3.Next = n5

	list1 = n1

	//list2 [2] -> [4] -> [6] -[7]
	n2.Next = n4
	n4.Next = n6
	n6.Next = n7

	list2 = n2

	mergedList := MergeList(list1, list2)
	/*
		for list1 != nil || list2 != nil && mergedList != nil {
			if list1 != nil && list2 != nil {
				//both don't we can compare the two
				if list1.Data <= list2.Data {
					if list1.Data != mergedList.Data {
						t.Errorf("Expected %v to be %v", list1.Data, mergedList.Data)
					}
					list1 = list1.Next
				} else {
					if list2.Data != mergedList.Data {
						t.Errorf("Expected %v to be %v", list2.Data, mergedList.Data)
					}
					list2 = list2.Next
				}
			}

			//either list2 or list1 is nil
			if list1 != nil {
				if list1.Data != mergedList.Data {
					t.Errorf("Expected %v to be %v", list1.Data, mergedList.Data)
				}
				list1 = list1.Next
			}

			if list2 != nil {
				if list2.Data != mergedList.Data {
					t.Errorf("Expected %v to be %v", list2.Data, mergedList.Data)
				}
				list2 = list2.Next
			}
		}
	*/

	expectStart := 1
	for mergedList != nil {
		if mergedList.Data != expectStart {
			t.Errorf("Expected %v to be %v", expectStart, mergedList.Data)
		}
		expectStart++
		mergedList = mergedList.Next
	}

}
