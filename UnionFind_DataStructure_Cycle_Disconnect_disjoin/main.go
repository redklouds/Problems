package main

import (
	"fmt"
	"testing"
)

//Making  Union Find Data Structure
var ValtoFind = 11

type unionFind struct {
	Arr          []int
	numVerticies int
}

//given number of nodes in this graph
//Given N , NumNodes, labeled from 0 to N - 1
func NewUnionFindDS(numNodes int) unionFind {
	//initalize the Union Find Array
	defaultDelimiter := -1 //because we know that nodes labels are from 0 to N -1
	arr := make([]int, numNodes)
	for i := range arr {
		arr[i] = defaultDelimiter
	}
	return unionFind{
		Arr:          arr,
		numVerticies: numNodes,
	}
}

func (u *unionFind) InsertEdge(a, b int) {
	//to Insert an Edge we need to first Find the nodes absolute parent
	//then see if it's in a different sub group of the graph/subgraph/subgroup/disconnected
	//then if we want we can stitch them together
	//or notify that there exists a cycle at this edge
	parentA := u.dfsFind(a)
	parentB := u.dfsFind(b)

	if parentA == parentB {
		//they are in the same group, we have a cycle IFF we add this edge in, since
		//if they are in the same group WE KNOW UNDIRECTED GRAPHS of the same group have paths
		//that connect all nodes within the group, that's why they are in the group in the first place..
		fmt.Printf("If we add Edge {%d, %d} we will create a cycle", a, b)

	} else {
		//we don't want a cycle , they exists in two different sub groups lets union Them
		u.union(parentA, parentB)
	}
}

//given two absolute PARENT NODES join them from A <- B
func (u *unionFind) union(a, b int) {
	//important* it doesn't mnatter which way we merge , BUT WE NEED TO BE CONSISSTANT, either Left to Right or Right to left
	//since its undirected [0,1] to [1,0] are the SAME *** THINK ABOUT IT *** the eduges are unique and undirected
	u.Arr[a] = b
}

//performs Find on the array data structure
func (u *unionFind) dfsFind(a int) int {

	//if the current A[i] != -1 we are not at the absolute parent yet continue to dfs search
	if u.Arr[a] == -1 {
		return a
	}
	//we have found that A IS the absolute parent therefore the node 'a' is the index which is ALSO the node
	return u.dfsFind(u.Arr[a])
}

func BenchMarkNaive(b *testing.B) {
	edgeList := [][]int{
		{0, 2},
		{6, 7},
		{1, 5},
		{0, 3},
		{6, 10},
		{9, 5},
		{0, 1},
		{7, 8},
		{5, 6},
		{4, 6},
		{11, 8},
	}
	numNodes := 12
	u := UnionFind{}
	for i := 0; i < b.N; i++ {

		u.MakeSet(numNodes, edgeList)
		u.Find(ValtoFind)
	}
	//fmt.Printf("Contains Multiple Comp?: %t\n", u.HasMultipleDisjoinSet())
	//fmt.Printf("DS : %v\n\n", u.Arr)

}

func BenchmarkPathCompression(b *testing.B) {
	edgeList := [][]int{
		{0, 2},
		{6, 7},
		{1, 5},
		{0, 3},
		{6, 10},
		{9, 5},
		{0, 1},
		{7, 8},
		{5, 6},
		{4, 6},
		{11, 8},
	}
	numNodes := 12
	djPathCompression := DisjoinsetPathCompression{}
	for i := 0; i < b.N; i++ {

		djPathCompression.MakeSet(numNodes, edgeList)
		djPathCompression.Find(ValtoFind)

	}
	//fmt.Printf("Contains Multiple Comp?: %t\n", djPathCompression.HasMultipleDisjoinSet())
	//fmt.Printf("DS : %v\n\n", djPathCompression.Arr)
}

func BenchmarkUnionByRank(b *testing.B) {
	edgeList := [][]int{
		{0, 2},
		{6, 7},
		{1, 5},
		{0, 3},
		{6, 10},
		{9, 5},
		{0, 1},
		{7, 8},
		{5, 6},
		{4, 6},
		{11, 8},
	}
	numNodes := 12
	djUnionByRank := DisjointSet{}
	for i := 0; i < b.N; i++ {

		djUnionByRank.MakeSet(numNodes, edgeList)
		djUnionByRank.Find(ValtoFind)
	}
	//fmt.Printf("Contains Multiple Comp?: %t\n", djUnionByRank.HasMultipleDisjoinSet())
	//fmt.Printf("DS : %v\n\n", djUnionByRank.Arr)

}

func main() {

	a := testing.Benchmark(BenchMarkNaive)
	fmt.Printf("BenchMarkNaive: %s\n", a)

	b := testing.Benchmark(BenchmarkPathCompression)
	fmt.Printf("BenchmarkPathCompression %s\n", b)

	c := testing.Benchmark(BenchmarkUnionByRank)
	fmt.Printf("BenchmarkUnionByRank: %s\n", c)

	/*
			BenchMarkNaive: 16216194	        70.67 ns/op
		BenchmarkPathCompression 14457412	        82.79 ns/op
		BenchmarkUnionByRank: 13714238	        88.27 ns/op
	*/

}
