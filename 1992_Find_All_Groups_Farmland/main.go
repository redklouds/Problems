package main

import "fmt"

func findFarmLands(land [][]int) [][]int {
	solutionSet := [][]int{}

	directions := [][]int{
		//X,Y
		{0, 1},  //DOWN
		{1, 0},  //RIGHT
		{-1, 0}, //LEFT
	}

	visited := make([][]bool, len(land))
	for i := range visited {
		visited[i] = make([]bool, len(land[0]))
	}

	for row := range land {
		for col := range land[0] {

			if land[row][col] == 1 && !visited[row][col] {
				//we found new plot of land as the FIRST top left corner
				topLeftCoords := []int{row, col}  //top left start of the farm land
				localSolution := make([][]int, 0) //here we store the coords
				dfsFarmLandBRUTEFORCE(land, row, col, visited, localSolution, directions)
				//after we finsedh searching search the local solution set ofr the latest x and y coords
				localMaxX := -1
				localMaxY := -1
				for _, pair := range localSolution {
					localMaxX = Max(localMaxX, pair[0])
					localMaxY = Max(localMaxY, pair[1])
				}
				topLeftCoords = append(topLeftCoords, localMaxX, localMaxY)
				solutionSet = append(solutionSet, topLeftCoords)
			}
		}
	}

	return solutionSet
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfsFarmLandBRUTEFORCE(land [][]int, row int, col int, visited [][]bool, solutionSet [][]int, directions [][]int) {
	//basically lets just get ALL the cells/coords of the land masses
	//then we find the largest values in each
	if row < 0 || row > len(land)-1 || col < 0 || col > len(land[0])-1 || land[row][col] == 0 || visited[row][col] {
		return // don't search this
	}

	//otherwise we have found a nice plot of farm land lets record its coords
	solutionSet = append(solutionSet, []int{row, col})

	//now lets search its adjcent cells
	//ONLY BFS you need to mark visited when you enqueue the nodes, else you visite nodes twice
	//ie if you have like a
	/*
		A--|
			|B that like 'L' shape you enqueue B from C but then from D you also see B and enqueue it again

			A---B in BFS if you are on Node B you enqueue node D, then proceed to process Node C, and C you see that
			|	| C D is also adj to Node C  and Enqueue it again, in BFS if you don't check visited in the LOOP, then you run dups
			|	| With DFS you jst check visited before the adj loop
			C---D
	*/

	visited[row][col] = true
	for _, dir := range directions {
		adjRow := row + dir[0]
		adjCol := col + dir[1]
		dfsFarmLandBRUTEFORCE(land, adjRow, adjCol, visited, solutionSet, directions)
	}
}

func main() {
	/*
		landData := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}

		res := findFarmLands(landData)
		fmt.Println(res)

	*/
	a := []int{6, 999}
	UpdateArr(&a)

	fmt.Println(a)

	LtoR := true

	for i := 0; i < 10; i++ {
		LtoR = !LtoR
		fmt.Printf("%t\n", LtoR)
	}

	testD := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	reverseList(testD)
	fmt.Println(testD)
}

func UpdateArr(arr *[]int) {
	*arr = append(*arr, 1, 2, 3, 4)
}

func reverseList(arr [][]int) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
}
