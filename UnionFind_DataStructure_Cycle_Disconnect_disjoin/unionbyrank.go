package main

//this is union by rank, what this does it it keeps track of a 'ranking' basically every time we have
//a node that has another node under it, we increment its rank, and we union by rank, the higher rank
//will always be the parent node of the smaller rank, incrementing the rank approperiately
//what this does, is it makes the disjoint  set more flatter, meaning there is not chance of a linear
//run time, we now make our union O(logN) run time, because: putting lower ranked sub trees directly
//underneath the higher (absolute parent) doesn't allow us to keep statically add to the absolute balue creating a linear linked list

//naivve approach think nodes 1 2 3 4, with ediges (1,2) (2, 3) (3,4)
//you will have a disjoint set datastructure of (if merging from left to rightg everytime
// 1 <- 2 <- 3 <- 4, where instead youu can have 1 <-2 , 1 <-3, 1<- 4 they are all connected to the same parent, anyways.. right??

//we can model this datastrucure in two(2) ways
//either use the object
/*
type Subset struct {
	Parent int
	Rank   int
}
*/

//or using two arrays,
//1 for Ranks storing
//1 for Parent storing
//parent := []int{}
//rank := []int{}

//using a struct

type Subset struct {
	Parent int
	Rank   int
}

type DisjointSet struct {
	Arr []Subset //remember if you do this, the functions ARE NOT POINTERS, you want to pass the actual array
}

//given a vertext to look for lets, search and find it
//by default we delimite all nodes parent to be themselves
//so our base case would be we recurse up the disjoint set until our vertext is ourselves (you need to look at notes)
func (d *DisjointSet) Find(vertext int) int {
	/*
		if d.Arr[vertext].Parent != vertext { //********GOTCHASSSS, THINK IN RECURSION, IF YOU DON'T 'HOLD' THE RECURSIVE STACK
		//THEN IT WILL JUST CONTINUE THE CALL STACK, ONLY DO THIS IF YOU WANT IT TO MIS BEHAVE LOL
			d.Find(d.Arr[vertext].Parent)
		}
		return d.Arr[vertext].Parent
	*/
	if d.Arr[vertext].Parent == vertext {
		return d.Arr[vertext].Parent //or return vertext
	}
	return d.Find(d.Arr[vertext].Parent)
}

//given two nodes, merge/union them
func (d DisjointSet) Union(a, b int) {
	//if the two nodes are in the same subset/group/tree
	//we return don't do anything, there exists a cycle if we add this edge
	parentA := d.Find(a)
	parentB := d.Find(b)

	if parentA == parentB {
		return //there exists a cycle on this edge
	}

	//now we can compare them by rank, the higher rank will be the parent, if the ranks are the same
	//we make parentA be the parent
	if d.Arr[parentB].Rank > d.Arr[parentA].Rank {
		//parentB is the parent
		d.Arr[parentA].Parent = parentB //make Node A parent B
		d.Arr[parentB].Rank++           //increment ParentB
	} else {
		d.Arr[parentB].Parent = parentA
		d.Arr[parentA].Rank = d.Arr[parentA].Rank + d.Arr[parentB].Rank + 1
	}
}

func (d *DisjointSet) MakeSet(numNodes int, edges [][]int) {
	d.Arr = make([]Subset, numNodes)
	for i := 0; i < numNodes; i++ {
		temp := Subset{
			Parent: i,
			Rank:   0,
		}
		d.Arr[i] = temp
	}

	for _, edge := range edges {
		d.Union(edge[0], edge[1])
	}
}

func (d *DisjointSet) HasMultipleDisjoinSet() bool {
	numComps := 0
	for node, val := range d.Arr {
		if node == val.Parent { //if the index/node is the same as its parent, this is the absolute parent
			numComps++
			if numComps > 1 {
				return true
			}
		}
	}
	return false

}
