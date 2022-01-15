package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	//idea have a visited map

	//walk through the node,and fore each of its neighbors oh... we need to create its neighbors nodes

	//add them to a nieghbor map, and add them to the new nodes neighbors

	//basically if we have

	//copying is a little difficult since you need to add the neigbors to self.. but once yuou do that.. how
	//do you retrieve them?, you must keep track of the new nodes somewhere

	//so create a cloned Node map

	//using DFS or bfs, we want to first see if the cloned node exists in the clone map

	//if it does not clone it, then  we want to check if its neigbors exists in the clone map

	//if they do not create them and add them to the clone map AND add them to the current nodes neighbors

	//this way, when we encounter the new nodes, we can query for them as we encounter them, and we can add respective neibhors as we go

	clonedNodeStore := make(map[int]*Node) //can use map or 'set' (which is map in go)

	visitedStore := make(map[int]bool)
	//we will use bfs here/just idk just bfs , we can search as we encounter rather than the depth

	var que []*Node
	que = append(que, node)

	//while the q is not empty
	for len(que) != 0 {
		//q is not empty lets continue to search and process

		//dequeue a node
		n := que[0]
		que = que[1:] //resize to remove the first element
		//or use
		//copy(q, q[1:])

		//now that we have our element
		//see if it exists int he copied node store

		//KEY, THE PROBLEM STATES THAT ALL NODES.VAL ARE UNIQUE, MEANING WE CAN SE THIS
		//AS OUR STORE KEY
	
		if _, exists := clonedNodeStore[n.Val]; !exists {
			//if exists then we do not need to create a new one, else create a new one and add it

			newNode := &Node{
				Val: n.Val,
			}
			clonedNodeStore[n.Val] = newNode
		}

		//now we do the same for all neighbors
		for _, neighbor := range n.Neighbors {

			//see if the neighbors exists in the copy if they do add them to
			//the current new nodes neighbors, otherwise create them and add them
			if _, neighborExists := clonedNodeStore[neighbor.Val]; !neighborExists {
				clonedNodeStore[neighbor.Val] = &Node{Val: neighbor.Val}
			}
			clonedNodeStore[n.Val].Neighbors = append(clonedNodeStore[n.Val].Neighbors, clonedNodeStore[neighbor.Val])
			if _, visited := visitedStore[neighbor.Val]; !visited {
				que = append(que, neighbor)
				visitedStore[neighbor.Val] = true// WE NEEED THIS, because if youy think about a square with 4 nodes 1-4, when we reach TWO FOR FOUR(4)

				//SINCE WE ARE ONLY MARKING "VISITED" once we actually process that node, if 2 and 4 were proceeed before 3, you would see that the queue, would enquerue
				//3 TWICE, because it hasne;'t tehnically be visited, so 2 and 4 on their turn would add 3 twice, to mediate this, we need to MARK visited nodes as visited, when 

				//we add them to the queue, because, doing so won't break anything or make it so they aren't visited
				/*
					1 ----- 2
					|		|
					|		|
					4 ----- 3

				*/
				//we need to mark as visited
			}

		}
	}
	return clonedNodeStore[node.Val]
}

type NodeSet struct {
	set map[*Node]bool
}

func NewNodeSet() NodeSet {
	s := NodeSet{
		set: make(map[*Node]bool),
	}
	return s
}
