package main

import (
	"fmt"
)

/*
type Queue struct {
	list.List
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(val interface{}) {
	q.PushFront(val)
}

func (q *Queue) Dequeue() interface{} {
	val := q.Back()
	q.Remove(val)
	return val.Value
}
*/

type Queue struct {
	arr  []*TreeNode
	size int
}

func NewQueue() Queue {
	return Queue{
		arr:  make([]*TreeNode, 0),
		size: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
func (q *Queue) Enqueue(val *TreeNode) {
	q.size++
	q.arr = append(q.arr, val)
}

func (q *Queue) Dequeue() *TreeNode {
	//THE GOTCHAS are that when you trim the slice, and you are storing OBJECTS you MUST NIL them
	//LIKE THESE
	if !q.IsEmpty() {
		val := q.arr[0]
		q.arr[0] = nil
		q.arr = q.arr[1:]
		q.size--
		return val
	}
	return nil
}

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func LevelByLevelTraversalReturnArray(node *TreeNode) [][]string {

	//insert the first node
	q := NewQueue()
	q.Enqueue(nil)
	q.Enqueue(node)
	encounteredNil := false
	resultArr := make([][]string, 0)
	var level []string
	for !q.IsEmpty() {

		if n := q.Dequeue(); n != nil {
			encounteredNil = false
			//print the node,
			level = append(level, n.Val)
			//fmt.Printf("Node: %v ", n.Val)

			//enqueue the left and right nodes
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		} else {
			if !encounteredNil {
				encounteredNil = true
				//its nil we should add a break point
				q.Enqueue(nil)
				//fmt.Println("")
				if len(level) > 0 {
					resultArr = append(resultArr, level)
				}

				level = nil
			}

		}
		//will panic if we type assert a nil as a treeNode

	}

	return resultArr
}

func main() {

	n1 := &TreeNode{Val: "A"}
	n2 := &TreeNode{Val: "B"}
	n3 := &TreeNode{Val: "C"}
	n4 := &TreeNode{Val: "D"}
	n5 := &TreeNode{Val: "E"}
	n6 := &TreeNode{Val: "F"}
	n7 := &TreeNode{Val: "G"}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	//LevelByLevelTraversal(n1)

	q := NewQueue()

	q.Enqueue(n1)
	q.Enqueue(n3)
	r := q.Dequeue()
	fmt.Println(r)
	q.Enqueue(n1)
	q.Enqueue(n2)

	res1 := LevelByLevelTraversalReturnArray(n1)

	fmt.Printf("%v", res1)

}
