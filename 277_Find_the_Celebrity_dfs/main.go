package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */

/*
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
*/
func FindCelebrity(knows func(a, b int) bool, n int) int {
	//def findCelebrity(self, n: int) -> int://
	//this is the def

	//our idea we can use the knowlege of a DAG/Topological sort
	//Indegrees to our advanctage O(n^2) runtime and O(n) sapce

	//we can iterate all combinations asking if these know each other

	//at each iteration

	//check if a know b
	//if a knows b, then update the neighbor count
	//this means that a has a neighbor++
	//and b has indegree++

	//now at each add check if the indegree== n-1 and the neighbor == n-1 if it does return true

	//initalize our indegree count

	indegreeCount := make(map[int]int)
	for i := 0; i < n; i++ {
		indegreeCount[i] = 0
	}
	//initalize our neighbor count
	neighborCount := make(map[int]int)

	for i := 0; i < n; i++ {
		neighborCount[i] = 0
	}

	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			if know := knows(a, b); know {
				//if a knows b, then we need
				if a == b {
					continue
				}
				indegreeCount[b]++
				neighborCount[a]++

			}

			if know := knows(b, a); know {
				indegreeCount[a]++
				neighborCount[b]++

			}
		}
	}

	for node, idc := range indegreeCount {
		if idc == n-1 && neighborCount[node] == 0 {
			return node
		}
	}

	//have to check here since, the scenario {{1,1}{1,1}}, we only have 2 nodes, so n-1 = 1, the momment that we see that node 0 has a neighbor of 1 we also know that indegree for 1 is n-1 AND the
	//neighbor is zero DEFAULT BEFORE we even can cehck number 1 for their conditions, we exist way to early. so we need to check at the end for the condition
	return -1
	//else at the end return false

}
