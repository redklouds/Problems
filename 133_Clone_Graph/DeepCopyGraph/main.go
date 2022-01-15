package main

import (
	"DeepCopyGraph/queue"
	"DeepCopyGraph/stack"
	"container/list"
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func BFS(node *Node) {
	q := queue.Queue{}
	q.EnQueue(node)

	//trake if i've visited this specific node before by pointer
	visitedMap := make(map[*Node]bool)
	visitedMap[node] = true
	for !q.IsEmpty() {
		//while q is not emtpy
		n := q.DeQueue()

		//type assertion and cast,, make suir eits not nil and of this type
		v := n.(*Node)
		fmt.Printf("Node: %+v address: %v\n", v.Val, &v.Val)
		//check all children
		for _, neigh := range v.Neighbors {
			_, exists := visitedMap[neigh]
			if !exists {
				visitedMap[neigh] = true
				//we haven't seen this node before lets enqueue it
				q.EnQueue(neigh)
			}
		}

	}
}

// O( |V| + |E|)
//but O(V^2) if you are given an adjacency list, because you will need to iterate through the list
func DFS(node *Node) {
	//using stack (Explictly), you can also use recursion to throw the items into the stack
	stack := stack.Stack{}
	stack.Push(node)
	//using the pointer address we can make this a unique pinter map
	visitedMap := make(map[*Node]bool)
	fmt.Printf("%+v\n", visitedMap)
	visitedMap[node] = true
	for !stack.IsEmpty() {
		nodeV := stack.Pop().(*Node)

		fmt.Printf("Node: %+v Address: %v\n", nodeV.Val, &nodeV.Val)
		for _, ch := range nodeV.Neighbors {
			//for each child
			//check if tis visited and if not Push it and mark it
			if _, ok := visitedMap[ch]; !ok {
				//not been visited
				visitedMap[ch] = true
				stack.Push(ch)
			}
		}
	}
}

func DFSRecursion(node *Node) {
	//we will NOT need an extra stack for this operation, however it doesn't matter since the same allocation exists
	//if we do not use recursion and with a stack memory on the stack is allocated and a stack is created
	//at each child our objective is to continue to push our children into the stack and process the children first
	vm := make(map[*Node]bool)
	vm[node] = true
	DFSRecursiveHelper(node, vm)

}
func DFSRecursiveHelper(node *Node, visitedMap map[*Node]bool) {
	if node != nil {
		fmt.Printf("Node: %v Address: %v\n", node.Val, &node.Val)
		for _, ch := range node.Neighbors {
			if _, exists := visitedMap[ch]; !exists {
				visitedMap[ch] = true
				DFSRecursiveHelper(ch, visitedMap)
			}
		}
	}
}

type Queue struct {
	list.List
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(val interface{}) {
	q.PushBack(val)
}
func (q *Queue) Dequeue() interface{} {
	if !q.IsEmpty() {
		val := q.Front()
		q.Remove(val)
		return val.Value
	}
	return nil
}



func DeepCopy(node *Node) *Node {
	//need two maps one for tracking visited nodes and another for tracking our copies
	copiedMap := make(map[*Node]*Node)
	visitedMap := make(map[*Node]bool)
	queue := Queue{}

	queue.Enqueue(node)
	visitedMap[node] = true

	for !queue.IsEmpty() {
		//while queue not empty, take the values

		//create copy of current node
		//WE SHOULD ONYL VISIT A NODE ONCE, Link it to its neighbors (one way) A -> b, A - > C but not backwards, they wil be linked when we reach B and C respectively

		nVal := queue.Dequeue().(*Node)
		var copiedNode *Node
		//see if a copy of the copied node exists already
		elem, ok := copiedMap[nVal]
		if !ok {
			//DNE
			copiedNode = &Node{Val: nVal.Val}
		} else {
			copiedNode = elem
		}

		copiedMap[nVal] = copiedNode

		//now link the neightbors
		for _, ch := range nVal.Neighbors {
			//check if the child has been created if not create it
			elem, ok := copiedMap[ch]
			var copyCh *Node
			if !ok {
				copyCh = &Node{Val: ch.Val}
			} else {
				copyCh = elem
			}
			copiedMap[ch] = copyCh
			copiedNode.Neighbors = append(copiedNode.Neighbors, copyCh)

			if _, visited := visitedMap[ch]; !visited {
				//has not been visited does not exists in visited map
				//add it to queue, mark it,
				queue.Enqueue(ch)
				visitedMap[ch] = true
			}
		}

	}

	//fmt.Printf("Map %v", copiedMap)
	return copiedMap[node]
}

func main() {
	n1 := &Node{
		Val: 1,
	}

	n2 := &Node{
		Val: 2,
	}

	n3 := &Node{
		Val: 3,
	}
	n4 := &Node{
		Val: 4,
	}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	BFS(n1)

	fmt.Println("")
	DFS(n1)
	fmt.Println("")
	DFSRecursion(n1)

	/*
		testMap := make(map[*Node]*Node)

		var CopiedNode *Node
		testMap[n1] = n1
		if elem, ok := testMap[n1]; !ok {
			//does not exists
			//copy it
			CopiedNode = &Node{Val: 6969}
		} else {
			CopiedNode = elem
		}
		fmt.Printf("N1 Address: %v\n", testMap)
		fmt.Printf("N1 Addressdfdf: %v\n")
		fmt.Printf("ELEMNT value: %v", &CopiedNode)
	*/
	res := DeepCopy(n1)
	fmt.Println("fdgdfg")
	BFS(res)
}
