package main

//using a disjoint set, we can find number of components, the tricky thing is to understand what is a edge

//runtime for this is O(n) since we only visit each input value ONCE, and O(n) for the disjoin set inspace

func findCircleNum(isConnected [][]int) int {

	//just a disjoin set of how many groups there are

	djSet := make([]int, len(isConnected))

	for i := range djSet {
		djSet[i] = -1
	}

	for row := range isConnected {
		for col := range isConnected[0] {
			if isConnected[row][col] == 1 {
				//found an edge
				Union(djSet, row, col)
			}
		}
	}

	numComp := 0
	for i := range djSet {
		if djSet[i] == -1 {
			numComp++
		}
	}

	return numComp
}

func Find(arr []int, a int) int {
	if arr[a] == -1 {
		return a
	}
	return Find(arr, arr[a])
}

func Union(arr []int, a int, b int) {
	parentA := Find(arr, a)
	parentB := Find(arr, b)

	if parentA == parentB {
		return
	}
	arr[parentA] = parentB
}
