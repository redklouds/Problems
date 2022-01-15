package main

func NumConnectedCompUnDirectedGraph(numNodes int, edges [][]int) int {

	//edge cases

	//if there are zero edges we can safly say there are that many comps
	if len(edges) == 0 {
		return numNodes //num Nodes are all individual comps
	}
	if numNodes == 0 {
		return 0
	}
	if len(edges) == 1 {
		return 1
	}
	graph := map[int][]int{}
	//populate the graph

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	//initalize the visited map
	visited := map[int]bool{}

	numComponents := 0
	//loop through every node, checking if its been visited, if not dfs on that node
	for n := 0; n < numNodes; n++ {
		if !visited[n] {
			//not visited, increment count of component and make a search
			numComponents++
			dfsNode(n, visited, graph)
		}
	}
	return numComponents
}

func dfsNode(node int, visited map[int]bool, graph map[int][]int) {
	//mark the current node as visited
	//search all unvisited adjacent nodes
	visited[node] = true
	for _, adjN := range graph[node] {
		if !visited[adjN] {
			dfsNode(adjN, visited, graph)
		}
	}
}
