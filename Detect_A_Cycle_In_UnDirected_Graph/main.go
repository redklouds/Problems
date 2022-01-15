package main

type GraphNode struct {
	Val  byte
	AdjN []*GraphNode
}

//basically if you think in a UNDIRECTED GRAPH
//THINK
//UNDIRECTED,
//in a DIRECTED graph to find a cycle you just dfs until you find a a node that points to anothe node that we have viisted
//actually you use DFS, DFS is for searching along edge paths ,so by design if a loop existsed that path would go visit a node that already
//has been visited,

//the ONLY GOTCHA is when undirected graph your edges between two verticies point to each other, now that is a ok case, in visited map that is not considered
//a cycle, how to fix? well we KNOW that as long as the adj node points back to its parent, the one where it came from
//thats NOT a cycle, the only time its a cycle is if the adj node has been VISITED< and NOT my parent
func CycleExists(n *GraphNode) bool {
	if n == nil {
		return false
	}
	dummyGraph := &GraphNode{}
	defer func(g *GraphNode) {
		g = nil
	}(dummyGraph)
	visited := make(map[*GraphNode]bool)
	return dfsDetectLoop(n, dummyGraph, visited)

}

//returns true of a cycle exists, false otherwise
func dfsDetectLoop(curNode *GraphNode, parentNode *GraphNode, visited map[*GraphNode]bool) bool {
	//we need
	//mark the node
	visited[curNode] = true

	//we need to dfs check all adjNodes
	for _, adjNode := range curNode.AdjN {
		//if the node has been visited AND is NOT where we came from(parent) its a cycle

		if visited[adjNode] && adjNode != parentNode {
			return true
		}
		//if the node has nOT been visited, lets search it
		if !visited[adjNode] {
			dfsDetectLoop(adjNode, curNode, visited)
		}
	}
	return false
}
