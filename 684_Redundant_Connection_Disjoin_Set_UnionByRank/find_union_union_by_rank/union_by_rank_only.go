package union_by_rank

type SetNode struct {
	Rank   int
	Parent int
}

//WORST CASE O(logN) for Union
//WE CAN FURTHER optimize this by using Path Compression for O(1) Find , but we weon't its only useful if we call this multiple times
func findRedundantConnection(edges [][]int) []int {

	//using disjoint set, we can find two nodes, whome are already connected and would produce a cycle

	//subproblems

	//create the disjoint set
	djSet := make([]SetNode, len(edges)+1) //plus 1 because the nodes are not zero index'd, and we will have an extra 1
	//add edges until we find two edges, whom produce a cycle, return those edges
	for i := range djSet {
		djSet[i] = SetNode{0, i}
	}

	for _, edge := range edges {
		if result, found := Union(djSet, edge[0], edge[1]); found {
			return result
		}
	}
	return []int{}
}

//find_union
func Find(arr []SetNode, a int) int {
	if arr[a].Parent == a {
		return arr[a].Parent
	}
	return Find(arr, arr[a].Parent)
}

//Basically Union by Rank drops from O(n) runtime to O(logN), by shifting the parents
//of big subtrees and making the tree 'flatter' by moving smaller tries
//under taller tries
func Union(arr []SetNode, a int, b int) ([]int, bool) {
	parentA := Find(arr, a)
	parentB := Find(arr, b)
	//so we know the two aren't of the same set, so we can check further
	if parentA == parentB {
		return []int{a, b}, true
	}

	if arr[parentB].Rank > arr[parentA].Rank {
		//parent B is the new parent of parentA
		arr[parentA].Parent = parentB
		arr[parentB].Rank++ //we added a another level to parent A subtree, to 'flatten' rearrange
	} else {
		//rank rank is equal or parentA is higher, we make parentA parent of parent B
		arr[parentB].Parent = parentA
		//since we MERGED the two subsets, we need this new parentA rank to be combination of parent B and A plus 1
		arr[parentA].Rank = arr[parentA].Rank + arr[parentB].Rank + 1
	}
	return nil, false
}
