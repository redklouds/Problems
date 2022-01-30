package naive

func findRedundantConnection(edges [][]int) []int {

	//using disjoint set, we can find two nodes, whome are already connected and woui;d produce a cycle

	//subproblems

	//create the disjoint set
	djSet := make([]int, len(edges)+1) //plus 1 because the nodes are not zero index'd, and we will have an extra 1
	//add edges until we find two edges, whom produce a cycle, return those edges
	for i := range djSet {
		djSet[i] = -1
	}

	for _, edge := range edges {
		if result, found := Union(djSet, edge[0], edge[1]); found {
			return result
		}
	}
	return []int{}

}

//find_union
func Find(arr []int, a int) int {
	if arr[a] == -1 {
		return a
	}
	return Find(arr, arr[a])
}

//returns true if cycle found, false and nil otherwise
func Union(arr []int, a int, b int) ([]int, bool) {
	parentA := Find(arr, a)
	parentB := Find(arr, b)

	if parentA == parentB {
		return []int{a, b}, true
	}
	arr[parentA] = parentB
	return nil, false
}
