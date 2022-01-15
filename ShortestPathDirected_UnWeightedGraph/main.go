package main

import "fmt"

//given an UNWEIGHTED
//DIRECTED graph, find the shortest path from
//node A to Node B
//*THINK* here we KNOW that BFS natually will find the shortest
//path because it traverses LEVEL by LEVEL
//if there was a node along the path it would find the SHORTEST path from start
//node to the target node, because breadth searching search level by level not leaving any of themk
//unsearched, also DFS just takes a route and go DEEPER

//Q1 print out the shortest path from src vertex to EACH of its vertexes
//approach, either modify the node to keep a count "# of paths" in the node
//OR use a map to store the shortest paths

//using an adjList as input, a target and a source
//print out the number of vertices that the src is from target
func PrintOutDistanceFromSrcNode(adjList map[int][]int, srcNode int) {
	if len(adjList) < 1 {
		return
	}

	//set up
	//create a holder for number of distances between src and target
	//MAP INTALIZE SUBSTAINTIALLY DIFFERENT, WHEN YOU KNOW THE SIZE OF THE MAP AHEAD
	//USE MAKE TO INITALIZE WITH DEFINED SIZE OTHER WISE USE LITERAL, BECAUSE YOU WILL HAVE A PREALLOCATED
	//MAP INSTAED OF RESIZING THE UNDERLYING ARRAY
	numDistance := make(map[int]int, len(adjList)) //since we know the number of nodes we can preallote the space
	//visited Map
	visited := make(map[int]bool, len(adjList)) //preallocate the space for optimizations

	//enqueue the first node src

	q := []int{srcNode}
	//q[0] = srcNode
	for len(q) > 0 {
		n := q[0]
		q = q[1:] //shrink, no need to clear its zero based- we would if it was a pointer to obj

		//print out
		fmt.Printf("Node: %d is %d verticies away to Node: %d\n", n, numDistance[n], srcNode)
		//foreach adjacent node, if its unvisited, lets update its distance and enqueue it
		for _, adjN := range adjList[n] {
			if !visited[adjN] {
				numDistance[adjN] = numDistance[n] + 1
				visited[adjN] = true
				q = append(q, adjN)
			}
		}
	}
}

func main() {
	testData := map[int][]int{}

	testData[0] = []int{1, 2}
	testData[1] = []int{2, 3}
	testData[2] = []int{3}
	testData[3] = []int{4, 5}
	testData[4] = []int{6}
	testData[5] = []int{6}
	testData[6] = []int{}

	PrintOutDistanceFromSrcNode(testData, 0)
}
