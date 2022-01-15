package iter

import (
	"206_Reverse_Linked_List_DFS_Recur/reverse_list_original_iteritive/models"
	"fmt"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

//GOREGOUS SOLUTION O(n) runtime no extra space, if we keep
//inserting at the head of a new listNode we esstentially create a stack in place
func ReverseListNode(node *models.ListNode) *models.ListNode {
	dummyHead := models.ListNode{}

	//ptr for node
	ptr := node
	for ptr != nil {
		//iter over list
		temp := dummyHead.Next //hold temp ptr to next
		tempObj := *ptr        //actual object, dereferrence the obj

		tempObj.Next = temp
		dummyHead.Next = &tempObj //since dummy head .Next is a pointer
		//you need to give it NOT the obect but return a address pointer to that object
		//by putting ampersand there
		ptr = ptr.Next
	}

	return dummyHead.Next
}

//using a slice/stack using stack[:endOfCutOffIndex] <- retrn a slice from beginning upto but not including this index
//O(N) space and O(n) time
func ReverseLinkedList(node *models.ListNode) *models.ListNode {
	//using a stack

	stack := []models.ListNode{}
	ptr := node
	for ptr != nil {
		stack = append(stack, *ptr)
		ptr = ptr.Next
	}
	dummyHead := models.ListNode{}
	ret := &dummyHead
	for len(stack) != 0 {
		ret.Next = &stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = ret.Next
		ret.Next = nil
	}
	return dummyHead.Next
}

func main() {
	head := models.ListNode{Data: 1}
	n2 := models.ListNode{Data: 2}
	n3 := models.ListNode{Data: 3}
	n4 := models.ListNode{Data: 4}
	n5 := models.ListNode{Data: 5}

	head.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	Print(&head)

	newHead := ReverseLinkedList(&head)
	Print(newHead)

	reversedHead := ReverseListNode(newHead)

	Print(reversedHead)
}

func Print(node *models.ListNode) {
	fmt.Println("")
	ptr := node
	for ptr != nil {
		fmt.Printf("[%v]->", ptr.Data)
		ptr = ptr.Next
	}
}
