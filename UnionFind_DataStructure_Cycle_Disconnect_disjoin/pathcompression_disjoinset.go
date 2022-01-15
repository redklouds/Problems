package main

//this is a another optimization for the naive disjoint set
//called PATH COMPRESSION
//the idea is, if you have multiple nodes that the "FIND()" method calls
//and recursivly runs up the call stack to the ABSOLUTE parent node, there is a chance
//that for instance if you do arr[b] = arr[a] every time that it creates a linear linked list
//therefore the worst case is O(n), also if we think about it if we have a stucture 0(AbsoluteParent) <- 1 <- 2 <- 3
// or array [-1, 0, 1,2] , meaning when you called Find(3) = Find(2) = find(1) = Find(0) = -1 finally O(n) why??
//if we KNOW that 2 equals the parent of zero, we can optmize by setting Find(2) to return ZERO right away, skipping
//the whole call chain, therefore we can edit the subset sutrcutre,

//either have this extra array calledd "Parent" who store the parent of the current node, or make each element in the data array
//a struct that store the Parent ,

//what we can do is on the first call of Find(3), we return the absolute parent down the Find call, setting all the nodes that are 1 2 and 3
//their parent will be set to the aboslute parent instead of its immediate local parent, making the Find call amoriztized O(1) time so the next time Find()
//is called its Find() O(1) , how we do this is we reciursivly follow the flow up to find abs parent same, however when we find it, we return the abs parent
//update the current nodes parent to the abs parent, then all subseqetnt calls are also updated to the parent of its parent

type DisjoinsetPathCompression struct {
	Arr []Subset
}

//O(1) amortized
//so given a node to Find its parent/subgroup/abs parent
//we will be using a element called Subset type for each index, but can be done with a seperate array
func (d *DisjoinsetPathCompression) Find(node int) int {

	//recursive, basecase is since we inintalize the array with each node's parent as itself (self loop) to signify
	//its its own abs parent
	//just check if the pass in node is the same as my parent, that means i'm the abs parent
	if d.Arr[node].Parent != node {
		//if our parent doesn't equal the current node, we want to keep recursing to follow the array to find abs parent
		//BUT we also want to set THIS nodes parent to the abs parent
		d.Arr[node].Parent = d.Find(d.Arr[node].Parent) //recursivly call Find, giving its 'local' parent to keep going up, BUT we also are awaiting
		//a return statement , therefore when it does return we are setting THIS local parent to absolute parent, we want to make sure that
		//alllllll subseqent parents equal the SAME parent, therefore 'flattening' the tree structure, and amoritizing O(1) Find() runtime
	}

	//recursive case, where if/when we do find that the current stack arr[node] where NODE is the SAME as parent, we have the absolute parent reutnr it
	//return it down the recursive stack so that every subsequent call can uypdate its parent respectivly
	return d.Arr[node].Parent

}

func (d *DisjoinsetPathCompression) MakeSet(numNodes int, edges [][]int) {
	//initalize the array
	d.Arr = make([]Subset, numNodes)
	for i := 0; i < numNodes; i++ {
		temp := Subset{
			Parent: i,
			Rank:   0,
		}
		d.Arr[i] = temp
	}

	for _, edge := range edges {
		//add edges
		d.Union(edge[0], edge[1])
	}
}

//REMEMBER DISJOIN SET ONLY WORKS WHEN THINGS ARE UNDIRECTIONAL
//we want to merge the two,
func (d *DisjoinsetPathCompression) Union(a, b int) {
	//get the abs parent of both nodes
	parentA := d.Find(a)
	parentB := d.Find(b)

	//if the parent's are in the same group, this edge/realtionship is going to cause a cycle, redundant, we can add or do something
	if parentA == parentB {
		return //cycle on this edge// if we want to detect, Union can Return FALSE, to signify that the given edge is a cycle edge
	}

	//now since they are different, we can just set one the parent of the other 'see union by rank for optimizing this'
	//here we will just set A parent of B everytime
	d.Arr[parentA].Parent = parentB

}

func (d *DisjoinsetPathCompression) HasMultipleDisjoinSet() bool {
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
