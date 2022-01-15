package main

import "testing"

func TestMain(t *testing.T) {
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1 := &Node{Val: 1}

	//create the test Nodes
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)

	cloneGraph(node1)


}


