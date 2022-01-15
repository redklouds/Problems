package main

import "fmt"

//OPTIMIZED WITH USING A SUBSET Struct instead of two(2) arrays for parent and rank
//OPMIZED TO UNION BY RANK AKA Union "By RANK"
//AND PATH COMPRESSION -> AKA "Set each subsequent Find() parent to its abs parent"

//O(1) for Find,
//O(n) for making
//O(n) space
type Subset struct {
	Parent int
	Rank   int
}

func findRedundantConnection(edges [][]int) []int {
	//i shall use a disjoin Set Data structure to detect wether there is a redundant edge
	//IE that the edge im adding is going to result in a cycle, since by def the disjoint set,
	disjoinSet := make([]Subset, len(edges)+1)

	//** SO UNION BY RANK NEEDS A PARENT ARRAY, Might as well impl the Path Compression
	//by default rank array is zero based
	for i := 1; i <= len(edges); i++ {
		temp := Subset{
			Parent: i,
			Rank:   0,
		}
		disjoinSet[i] = temp
	}

	var redundantEdges [][]int
	for _, edge := range edges {
		//for each edge union them
		if CheckIfCycleUnion(disjoinSet, edge[0], edge[1]) {
			//here is a redundant edge
			redundantEdges = append(redundantEdges, []int{edge[0], edge[1]})
		}
	}
	fmt.Println(redundantEdges)
	return redundantEdges[len(redundantEdges)-1]

}

//Optimized with Path Compression, using Parent
//given an array find the parent node and return it
func Find(parArr []Subset, a int) int {
	//watch for the recursive GOTCHAS, IE making if arr[a] != -1 { Find(arr, arr[a])}, return arr[a], because remember stack will continue
	if parArr[a].Parent != a {
		parArr[a].Parent = Find(parArr, parArr[a].Parent)
	}
	//NOW the gotchas is not a problem, BECAUSE in the if clause we 'pause the program' and wait until we return
	return parArr[a].Parent
}

//given two nodes(a edge), return true if this edge creates a cycle
func CheckIfCycleUnion(arr []Subset, a, b int) bool {
	parentA := Find(arr, a)
	parentB := Find(arr, b)

	if parentA == parentB {
		return true
	}

	//now we compare the rank of each sub tree
	if arr[parentA].Rank > arr[parentB].Rank {
		//parent A has higher rank make parentA the parent of ParentB
		arr[parentB].Parent = parentA
		arr[parentA].Rank++
	} else {
		//the ranks are either equal or B is higher
		arr[parentA].Parent = parentB
		arr[parentB].Rank = arr[parentB].Rank + arr[parentA].Rank + 1
	}
	return false
}

func main() {
	TestEdges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{1, 5},
	}

	result := findRedundantConnection(TestEdges)
	fmt.Println(result)
}
