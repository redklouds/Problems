package main

func countComponents(n int, edges [][]int) int {

	//create the adj list
	adjList := make(map[int][]int)
	//n is how many nodes
	for i := 0; i < n; i++ {
		adjList[i] = []int{}
	}
	for _, edge := range edges {
		//edges[0] - edge[1] both have edges to each other
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	countOfComponents := 0
	visited := make(map[int]bool)
	for node := 0; node < n; node++ {
		if !visited[node] {
			dfs(adjList, visited, node)
			countOfComponents++
		}
	}

	return countOfComponents
	//iterate from 0 to n nodes and keep count JUST like counting islands
}

func dfs(adjList map[int][]int, visited map[int]bool, node int) {
	if visited[node] {
		return
	}

	visited[node] = true

	for _, n := range adjList[node] {
		//for all my neighbors dfs into them
		dfs(adjList, visited, n)
	}
}
