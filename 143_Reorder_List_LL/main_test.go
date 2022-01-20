package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reorderList(head)
	expected := []int{1, 5, 2, 4, 3}
	if !validateList(expected, head) {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
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
	reorderList(head)
	expected := []int{1, 4, 2, 3}
	if !validateList(expected, head) {
		t.Error()
	}
}

func TestMain3_V1(t *testing.T) {
	//TODO:
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	reorderList_V1(head)
	expected := []int{1, 4, 2, 3}
	if !validateList(expected, head) {
		t.Error()
	}
}

func TestMain4_V1(t *testing.T) {
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

	reorderList(head)
	expected := []int{1, 4, 2, 3}
	if !validateList(expected, head) {
		t.Error()
	}
}

//verify that the linklist matches the order and amount in the given answer
func validateList(ans []int, head *ListNode) bool {

	for i := range ans {
		if head == nil {
			return false
		}
		if head.Val != ans[i] {
			return false
		}
		head = head.Next
	}
	return true
}
