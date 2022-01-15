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
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	//extra memory solution,
	//we can create a pathset, to see which nodes we have vbisited, if we visit a node that has already been visited, before, then we know that they intersect since
	//they both reference the same node

	//3 ways, hashing (bviisted hash)
	//edit the nod eot have visited gflah
	//two pointers to euqlize

	//basically traverse the two list, one node at a time, whe one pointer eachs the end
	//put it on the opposite list and continue traversing

	ptr1 := headA
	ptr2 := headB
	ptr1ReAssigned := false
	ptr2ReAssigned := false
	for !ptr1ReAssigned || !ptr2ReAssigned {
		if ptr1 == ptr2 {
			return ptr1 //already found the match
		}
		if ptr1 == nil {
			ptr1 = headB
			ptr1ReAssigned = true
		}
		if ptr2 == nil {
			ptr2 = headA
			ptr2ReAssigned = true
		}
		ptr2 = ptr2.Next
		ptr1 = ptr1.Next
	}

	//now that they both are reassign they are both at the same equal distance to/from the intersection point
	for ptr1 != nil || ptr2 != nil {
		//if one of them equal nil jump out we don't have a intersect
		if ptr1 == ptr2 {
			return ptr1
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return nil

}
