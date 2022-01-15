package main

import "fmt"

//https://wentao-shao.gitbook.io/leetcode/graph-search/547.number-of-provinces

//I WILL PERFORM a disjoint SET here
//implemention Path Compression AND union by rank to get O(1) Find and O(logN ) union amoritized

type Subset struct {
	Parent int
	Rank   int
}

type UF struct {
	Arr []Subset
}

func (u *UF) Find(a int) int {
	if u.Arr[a].Parent != a {
		//recurse, because initalizally abs parent are equal parents of themselves
		//pathcompression
		u.Arr[a].Parent = u.Find(u.Arr[a].Parent)
	}
	return u.Arr[a].Parent
}

func (u *UF) Union(a, b int) {
	//get both parent nodes
	parentA := u.Find(a)
	parentB := u.Find(b)

	//if same skip
	if parentA == parentB {
		return //contains cycle we don't want dat
	}

	//check the rank higher is parent
	if u.Arr[parentA].Rank > u.Arr[parentB].Rank {
		//Parent A is parent of b
		u.Arr[parentB].Parent = parentA
		//increment parentA
		u.Arr[parentA].Rank++
	} else {
		//default goes to b equal or b is greater
		u.Arr[parentA].Parent = parentB
		u.Arr[parentB].Rank = u.Arr[parentB].Rank + u.Arr[parentA].Rank + 1
	}
}

func (u *UF) MakeUF(numNodes int, edges [][]int) {
	//im guessing its starts from 1 to n incluseive, and not zero, directions suck. LOL
	u.Arr = make([]Subset, numNodes+1)
	//initalize the subset to parent of itself
	for i := 1; i <= numNodes; i++ {
		temp := Subset{
			Parent: i,
			Rank:   0,
		}
		u.Arr[i] = temp
	}

	//now build the disjointset
	//EDGES IS AN ADJMATRIX, meaning its undirected and SYMMETRICAL, meaning half of the matrix is useless

	for row := 0; row < len(edges); row++ {
		for col := 0; col < row; col++ {
			if edges[row][col] == 1 {
				u.Union(row, col)
			}
		}
	}
	/*
	   for row:=0; row< len(edges); row++{
	       for col:=0; col< len(edges); col++ {
	           if row != col && edges[row][col] == 1{
	               u.Union(row, col)
	           }
	       }
	   }
	*/
}

func (u *UF) findNumGroups() int {
	num := -1
	//the first one ops
	for node, val := range u.Arr {
		if node == val.Parent {
			num++
		}
	}
	return num
}
func findCircleNum(isConnected [][]int) int {

	if len(isConnected) == 0 {
		return 0
	}

	uf := UF{}
	uf.MakeUF(len(isConnected), isConnected)
	return uf.findNumGroups()
}

type DJ []Subset

func main() {
	dj := DJ{}

	dj = []Subset{}
	input := [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}
	fmt.Println("Num : ", findCircleNum(input))
}
