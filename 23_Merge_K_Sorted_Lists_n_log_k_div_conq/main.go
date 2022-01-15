package main

/*
There are a couple ways the easiest is to grab each list and combin them into the first
list, working out way down to the end
this will run at O(n^1) we can do a O(n log(k)), where k is the list, and O(n) is each list in the array
we know merging is an O(n) operation, with no extra space,

*/

type ListNode struct {
	Next *ListNode
	Val  int
}

func mergeKLists(lists []*ListNode) *ListNode {

	//base case
	if lists == nil || len(lists) == 0 {
		return nil
	}

	//idea is to work k/2 iterations, we know that mergeing two list is done O(n) times
	//we just pair the list and add the result to the left side over and over until we are left with a single list

	lastListIdx := len(lists) - 1
	//go until only a single list is left
	for lastListIdx != 0 {
		leftList, rightList := 0, lastListIdx

		for leftList < rightList {
			//merge the lists right and left tofether into the left
			lists[leftList] = sortedMerge(lists[leftList], lists[rightList])
			
			//consider the next pair
			leftList++
			rightList--

			//check if all pairs are merged, update
			if leftList >= rightList{
				lastListIdx = rightList
			}
		}
	}
	return lists[0]
}

//takes two list sorted in increasin gorder and merges their nodes
//to make one big sorted list returned
func sortedMerge(a, b *ListNode) *ListNode {

	//base case
	//if one of the list are null return their other
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	//ie if a is a the end we return b, vice versa we do this becasue
	//each list is SORTED if we keep processing each list, we know that at the end the
	//element at the top will be the nex tone to process

	var result *ListNode

	//pick either a or b , and recurse
	if a.Val <= b.Val {
		//so if a is smaller, we want to say that the current node is
		//a then its next will be assigned as the smallest from either list
		//we advance a to the next

		//DO KEEP IN MIND THE BOTTOM RETURN STATEMENT
		//will retuhrn the 'result' at the end, back to the callers, which in the stack frame
		//is the CURRENT value
		result = a
		result.Next = sortedMerge(a.Next, b)
	} else {
		result = b
		result.Next = sortedMerge(a, b.Next)
	}
	return result
}
