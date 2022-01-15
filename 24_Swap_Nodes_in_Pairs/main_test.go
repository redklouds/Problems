package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	expected := []int{2, 1, 4, 3}
	if actual := swapPairs(head); !ContainsSameValuesInOrder(expected, actual) {
		t.Error()
	}

}

func TestMain6(t *testing.T) {
	//TODO:
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	expected := []int{2, 1, 3}
	if actual := swapPairs(head); !ContainsSameValuesInOrder(expected, actual) {
		t.Error()
	}

}

func TestMain2(t *testing.T) {
	//TODO:
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	expected := []int{1}
	if actual := swapPairs(head); !ContainsSameValuesInOrder(expected, actual) {
		t.Error()
	}

}

func TestMain3(t *testing.T) {
	//TODO:

	expected := []int{}
	if actual := swapPairs(nil); !ContainsSameValuesInOrder(expected, actual) {
		t.Error()
	}

}

func ContainsSameValuesInOrder(arr []int, list *ListNode) bool {
	currPtr := list
	for i := range arr {
		if currPtr.Val != arr[i] {
			return false
		}
		currPtr = currPtr.Next
	}
	return true
}
