package main

/* Given n Nodes labeled 0 to n-1 and
a list of UNDIRECTED EDGES(each edge is a pair of nodes), write a function
to check whether these edges make up a valid tree.


*TIP -> pay attention to all details, ;-> you can assume that
NO duplicate edges will appear in the edges, since all edges are
 UNDIRECTED, [0,1] is the same as [1,0] TRUE. and thus will not
 appear together in edges

 THINK -> this helps our generation of the graph much easier, since if we encountered

 [0,1] and [1,0] in the edges it will produce DUPLICATE adjacent nodes to each of them
  OR if there dduplicates it would allow us to make the adjacent list not adding both sides but waiting

*/

//O(n) - worst case for Runtime since if there exists no cycles we visit ever node
//(n) - space for the aux space on storing visited nodes
//idea

//we know that a graph is a Tree IFF,
// it contains No Cycles
//		if no cycles means that since undirected,
//		if we find a node thats been VISITED, AND is NOT our parent
//		(parent) the node we traversed from, we know its a cycle
//		otherwise its just the backedge back to our parent as an adjacent node
// DFS is perfectly selectd because it Traverses only the connected componnets, unlike the BFS topological traversal
// is connected
//		if we are given the number of nodes, in the graph,
//		we can dfs and compare # of visited to num of nodes, to see if they match
func IsGraphValidTree(numNodes int, edges [][]int) bool {

	if len(edges) < 2 {
		return true
	}
	//first create the graph, I will be using a adjacency list

	graph := make(map[int][]int)

	for n := 0; n < numNodes; n++ {
		graph[n] = []int{} //empty adjlist
	}

	//populate the edges now
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	} //they are undirected 'pointing to each other'

	visited := map[int]bool{}
	//this function checks the current node with its parents,
	//if its truly connected and its undirected we can select any point and visit
	//all nodes, so i default to zero
	//the parent is -1, since the nodes are labled 0 -> n-1 we can just pick something that is 'dummy' parent
	//that won't flip the flag
	containsCycle := dfsCycleSearch(0, -1, visited, graph)

	//now we know if there is a cycle, return false 'isn't a valid graph'

	if containsCycle {
		return false
	}
	//else if the contains Cycle Is false 'doesn't contain a cycle'
	//we need to check if all nodes have been visited, go maps are pas by reference so
	//contains the visited nodes
	if len(visited) != numNodes {
		return false
	}
	return true //all conditions met, happy path

}

//given the current node, and its  previous parent that it decended from
//tell me if there exists a cycle
func dfsCycleSearch(node int, parentNode int, visited map[int]bool, graph map[int][]int) bool {

	//mark current node
	visited[node] = true
	//search its children/adjacent nodes
	for _, adjN := range graph[node] {
		//if the current adjacent node is VISITED, we must see if its a false posiitve, by checking if its our parent
		//if it is, that's OK, otherwise its a loop
		if visited[adjN] && adjN != parentNode {
			return true //loop exists
		}
		//otherwise if we haven't visited this adjacent node lets visit it
		if !visited[adjN] {
			dfsCycleSearch(adjN, node, visited, graph)
		}
	}
	return false //no cycle we visited all nodes, w/o encountering a loop
}
