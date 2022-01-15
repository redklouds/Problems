package main

import (
	"fmt"
	"sort"
)

//SLOWWWW

func allPathsSourceTarget(graph [][]int) [][]int {

	//idea is using the strategy 'save and revert state' pattern
	//basically we will keep a pathset that will keep record of all the paths we walk,
	//then when we walk the path we add the neighbors when we find target add the pathset to the result array and exist the

	//recursive function
	//pathSet := make(map[int]bool)

	//DESIGN OF PATH SET, for there will be an indexs, that tells me which is the last index to iterate until i get the current path
	//and that value also tells me which value to overwrite the pathset too
	pathSet := make([]int, len(graph)) //this is pass by copy but since i will not be changing the header, i can safely pass
	//this pathset around
	//pathSet[0] = 0
	resultSet := [][]int{}
	//resultMap := map[int][][]int{}
	//pathSet[0] = true
	//ordered pathset = we can use each index as the node itself, initalize the nodes with -1 to signify end of path, or
	//that we
	dfs(pathSet, &resultSet, graph, 0, len(graph)-1, 0) //pass by reference to the array
	//second issue is the pathset is not sorted or ordered
	return resultSet //resultMap[0]

}

func dfs(pathStack []int, resultSet *[][]int, adjList [][]int, node int, targetNode int, pathIdx int) {
	pathStack[pathIdx] = node
	//basecase
	if node == targetNode {
		//we found our target, add the current pathset to the result set
		*resultSet = append(*resultSet, getPathSetFromArr(pathStack, pathIdx))
		return //getPathSet(pathset) //exit this loop
	}

	//if we are not the target, lets search its neighbors
	for _, neighbor := range adjList[node] {
		//pathset[neighbor] = true
		//push this onto the path stack
		//	pathStack = append(*pathStack, neighbor)

		pathIdx++
		dfs(pathStack, resultSet, adjList, neighbor, targetNode, pathIdx)
		//pop off the stack since we finish this neibhor on our path stack/set
		pathIdx--
		//delete(pathset, neighbor) //for a set you need to REMOVE the key from the map else it just sites there

	}

}

func getPathSetFromArr(pathSet []int, lastIndex int) []int {
	resultSet := make([]int, lastIndex+1)
	for i := range resultSet {
		resultSet[i] = pathSet[i]
	}
	return resultSet

}

func getPathSet(pathSet map[int]bool) []int {
	resultSet := make([]int, len(pathSet)) //we know the length/num keys in pathset // so use it
	i := 0
	for k := range pathSet {
		resultSet[i] = k //adding key to index
		i++
	}
	sort.Ints(resultSet)
	return resultSet
}

func main() {

	//Go does pass by value when we modify the array header

	//but willm work if you only add or modiy to exists spaces

	//meaning that if i have a allocated array and just change the index value it works
	arr := make([]int, 5)
	modifyArray(arr)
	fmt.Println(arr)

}

func modifyArray(arr []int) {
	arr[3] = 60606
}
