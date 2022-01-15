package main

import (
	"fmt"
	"math"
)

//Depth first search is a search that is performed
//to search a tree or a graph, it is called DFS - Depth First search
//because it prioritizes Depth First, so
//its main objective is to Explore DEEPER into the graph whenever possible
//this if there was a node with 3 adjacent nodees, it will grab the first adjacent node
//and traverse that next, and repeat,

//BFS - breadth first search uses a queue, its objective is too..

//if you have a tree
/*
		A

	/C	B	D
Z	y	k
BFS will search all the adjacent nodes as priority, whre DFS will search the deepest
meaning that when a DFS is performed it will keep continuing each time for the first adjacent unvisited node
//whre as BFS will traverse everyadjacent node FIRST before traversing onto the current nodes adjacent nodes



//EX for DFS willSearch the current node, then it will mark it visited
//then it will traverse the FIRST unvisited Adjacent node of current
//and repeat that process

//EX BFS will search the current Node, Mark it visited,
//then it will search, nieghbor nodes next,




*/

type GraphNode struct {
	Val      string
	AdjNodes []*GraphNode
}

type GraphNodeNum struct {
	Val      int
	adjNodes []*GraphNodeNum
}

func BFS(node *GraphNode) {
	//breadth first search, will search not the DEEPTH but the breadth,  it will prioritize
	//exploring the adjacent neighbors first before going any deeper think LEVEL BY LEVEL exploring, insead of a single Straight shot to the bottom
	if node == nil {
		panic("Empty node")
	}

	visited := make(map[*GraphNode]bool)
	queue := make([]*GraphNode, 1)
	queue[0] = node
	for len(queue) > 0 {
		//while queue is still has items

		//enqeue its adjacent neighbors
		n := queue[0]
		visited[n] = true
		queue[0] = nil
		queue = queue[1:]
		fmt.Printf("Visited :%v\n", n.Val)
		for _, adjNode := range n.AdjNodes {
			if _, hasVisited := visited[adjNode]; !hasVisited {
				visited[adjNode] = true
				queue = append(queue, adjNode)
			}

		}
	}
}

func IterDFSUsingStack(node *GraphNode) {
	//using a stack to explicitly build without recursion

	if node == nil {
		panic("Cannot give me nil as input")
	}
	stack := make([]*GraphNode, 1)
	stack[0] = node

	visited := make(map[*GraphNode]bool)
	//while stack isn't empty
	for len(stack) > 0 {
		//push in its unvisited adjacent node
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[n] = true
		fmt.Printf("Visited: %v\n", n.Val)
		for _, adjN := range n.AdjNodes {
			if _, visited := visited[adjN]; !visited {
				stack = append(stack, adjN) //add it to the head where we will retrieve from
			}
		}

	}
}

func DFS(node *GraphNode) {
	if node != nil {
		visited := make(map[*GraphNode]bool)
		search(node, visited)
	}
}

//Depth First Search helper
//will find the first adjacent node and repeat its Depth FIRST SEARCH,

//Priority to explore DEEPER when possible
//Maps are passed by reference

//Map types are reference types, like pointers or slices[1]
func search(node *GraphNode, visited map[*GraphNode]bool) {

	if node != nil {
		fmt.Printf("Visiting %v\n", node.Val)
		visited[node] = true
		for _, adjN := range node.AdjNodes {
			//if current adj not visited, we HAVE TO EXPLORE DEEPERRRR
			if _, ok := visited[adjN]; !ok {
				//not visited
				search(adjN, visited)
			}
		}
	}
}

func ShortestPath(node *GraphNodeNum, target *GraphNodeNum) ([]*GraphNodeNum, int) {

	q := make([]*GraphNodeNum, 1)
	q[0] = node
	totalDistance := 0
	resultPath := make([]*GraphNodeNum, 0)
	visited := make(map[*GraphNodeNum]bool)
	visited[node] = true
	found := false
	for len(q) > 0 && !found {
		//still have something
		min_dist := math.MaxInt32
		var shortedNode *GraphNodeNum
		qlen := len(q)
		for i := 0; i < qlen; i++ {
			//search for the shorted distance for this current level
			//if we have found target at the current level, we break record, and return
			if q[i] == target {
				shortedNode = q[i]
				min_dist = q[i].Val
				found = true
				break
			}

			if q[i].Val < min_dist {
				shortedNode = q[i]
				min_dist = q[i].Val

			}

			//add these edges to the queue
			for _, adjNode := range q[i].adjNodes {
				if _, ok := visited[adjNode]; !ok {
					//not visited
					q = append(q, adjNode)
					visited[adjNode] = true
				}
			}

		}

		//finshed with getting the shortest path for this current level or the 1 edge away nodes lets record it into the path
		resultPath = append(resultPath, shortedNode)
		totalDistance += min_dist
		//clean the queue, trim it to remove the ones we processed
		q = q[qlen:] //from the length of our items to the end remove the front
	}
	if resultPath[len(resultPath)-1] != target {
		panic("Target does not exists here")
	}
	return resultPath, totalDistance
}

func ShortestPathBFSModified(src *GraphNode, target *GraphNode, adjMap map[*GraphNode][]*GraphNode, pred map[*GraphNode]*GraphNode) bool {
	if src == target {
		return true
	}

	q := make([]*GraphNode, 1)
	q[0] = src
	visited := make(map[*GraphNode]bool)
	//pred := make(map[*GraphNode]*GraphNode)

	visited[src] = true
	for len(q) > 0 {

		//while ther are nodes BFS search level by level

		u := q[0]
		q = q[1:] //trim back

		for _, adjNode := range adjMap[u] {
			if _, hasVisited := visited[adjNode]; !hasVisited {
				//not visited yet
				pred[adjNode] = u

				visited[adjNode] = true
				if adjNode == target {
					for key,val := range pred{
						fmt.Printf("Key %s, Val %s\n", key.Val,val.Val)
					}
					return true
				}
				q = append(q, adjNode)
			}
		}

	}

	return false
}

//the idea is that BFS NATURALLY UNweighted and Undirected will give you the shortest path in a unweighted undirected graph
//you wants to make a modified versxion of BFS, by using a adjacent map, to simply TRACK ALL PREDESSOR NODES IE, the Current Adjacent Nodes Parent
//basically if you are visiting A, and then its adjacent nodes B and C, you want to update the map for nodes B and C's PARENT NODE AS A, since you 
// you visited B and C from A node, so a predecessor map [adjNode]=Parent node will give you a mapping of the rouyte you took to a particulare target node

//ie

/*
Key C, Val A
Key G, Val A
Key O, Val C
Key Y, Val G
Key X, Val O
Key Z, Val Y
*/

//if src is A and Target is Z, you can follow the map from target up to src, key Z = value Y, and Key Y = Value G, Key G = Value A, so the route from A to Z is A->G->Y->Z,
//you can add this toi a list and print backwards to get the string forwards
func DisplayShortestPath(src *GraphNode, target *GraphNode, pred map[*GraphNode]*GraphNode) {
	//need to get the
	//ONLY USED IF WE HAVE FOUND A PATH
	//while key != src keep going up

	key := target
	resultPath := make([]string, 0)
	for key != src {

		val := pred[key]
		resultPath = append(resultPath, key.Val)
		key = val
	}
	resultPath = append(resultPath, src.Val)

	//print backwards

	for i := len(resultPath)-1; i >= 0; i-- {
		fmt.Printf("%v ->", resultPath[i])
	}

}
func main() {

	gnA := &GraphNode{
		Val: "A",
	}

	gnB := &GraphNode{
		Val: "B",
	}

	gnC := &GraphNode{
		Val: "C",
	}

	gnD := &GraphNode{
		Val: "D",
	}

	gnE := &GraphNode{
		Val: "E",
	}

	gnF := &GraphNode{
		Val: "F",
	}

	gnG := &GraphNode{
		Val: "G",
	}
	gnO := &GraphNode{
		Val: "O",
	}

	gnY := &GraphNode{
		Val: "Y",
	}

	gnZ := &GraphNode{
		Val: "Z",
	}

	gnX := &GraphNode{
		Val: "X",
	}

	gnA.AdjNodes = []*GraphNode{gnC, gnD, gnB}

	gnB.AdjNodes = []*GraphNode{gnA, gnF}
	gnC.AdjNodes = []*GraphNode{gnA, gnE}
	gnD.AdjNodes = []*GraphNode{gnA, gnG, gnF}
	DFS(gnA)
	fmt.Println("")
	IterDFSUsingStack(gnA)
	fmt.Println("")
	BFS(gnA)

	adjMap := make(map[*GraphNode][]*GraphNode)

	adjMap[gnA] = []*GraphNode{gnB, gnC, gnG}
	adjMap[gnB] = []*GraphNode{gnG}

	adjMap[gnC] = []*GraphNode{gnO}
	adjMap[gnG] = []*GraphNode{gnY}
	adjMap[gnO] = []*GraphNode{gnX}
	adjMap[gnB] = []*GraphNode{}

	adjMap[gnY] = []*GraphNode{gnZ}
	adjMap[gnX] = []*GraphNode{gnZ}

	pred := make(map[*GraphNode]*GraphNode)
	res := ShortestPathBFSModified(gnA, gnZ, adjMap, pred)
	fmt.Printf("a shortest path exits?  %v", res)
	DisplayShortestPath(gnA, gnZ, pred)
}
