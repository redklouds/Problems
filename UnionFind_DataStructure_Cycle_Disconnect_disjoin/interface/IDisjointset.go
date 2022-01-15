package main

type IDisjointset interface {
	Find(a int) int
	Union(a, b int)
	MakeSet(numNodes int, edges [][]int)

	HasMultipleDisjoinSet() bool
}
