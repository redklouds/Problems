package main

//this is the naive Non Optimized disjoint set data structure , very easy where we use -1 as a delimitor to say
//that this current node is the absolute parent

type UnionFind struct {
	Arr          []int
	numVerticies int
}

//O(n)
func (u *UnionFind) Find(a int) int {
	if u.Arr[a] == -1 {
		//we haven't reached the abs parent yet
		//**********(u.Find(u.Arr[a]) //FUCKING GOTCHAs, if you don't have left hand side something = RecursiveFunc() then it just goes to the next line
		//and doesn't do what you think it does, so you need to reverse
		return a
	}
	//***** RECURSIVE GOTCHAS return a
	return u.Find(u.Arr[a])
}

//O(n)
//merges, combines, two subsets/disconnected graphs
func (u *UnionFind) Union(a, b int) {

	//if they are from the same parent, return, this edge creates a cycle
	parentA := u.Find(a)
	parentB := u.Find(b)
	if parentA == parentB {
		return //creates a cycle here, modify this function to return boolean if you want to find cycles on edges here
	}
	//else they are different sub groups make parent of A, the Parent Of B
	u.Arr[parentB] = parentA
}

func (u *UnionFind) MakeSet(numNodes int, edges [][]int) {
	u.Arr = make([]int, numNodes)
	for i := 0; i < numNodes; i++ {
		u.Arr[i] = -1
	}

	//adding edges, REMEMBER UNDIRECTED GRAPHS are only working with disjoin set
	for _, edge := range edges {
		u.Union(edge[0], edge[1])
	}
}

func (u *UnionFind) HasMultipleDisjoinSet() bool {
	numComps := 0
	for _, val := range u.Arr {
		if val == -1 {
			numComps++
			if numComps > 1 {
				return true
			}
		}
	}
	return false

}
