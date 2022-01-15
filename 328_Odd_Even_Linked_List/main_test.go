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
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	expected := []int{1, 3, 5, 2, 4, 6}
	if actual := oddEvenList(head); !Equal(expected, actual) {
		t.Error()
	}

}

func Equal(ans []int, list *ListNode) bool {
	for i := range ans {
		if list.Val != ans[i] {
			return false
		}
	}
	return true
}
