package main

import "fmt"

func main() {

	area := 0
	x := modArea(area)

	fmt.Println(x)

	testData := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	maxAreaOfIsland(testData)

}

func maxAreaOfIsland(grid [][]int) int {

	//idea, for each cell search for '1', if i find a '1', DFS into it assuming that the islands are sepeareted forsure
	//if i find a '1'

	//dfs into it, passing a local int value to change, for each DFS instead of incrementing to couint number of island
	//i want to increment within the dfs function and return the number of recursive calls

	directions := [][]int{
		//X,Y
		{0, -1}, //UP
		{1, 0},  //RIGHT
		{0, 1},  //DOWN
		{-1, 0}, //LEFT
	}
	maxArea := 0

	for row := range grid {

		for col := range grid[0] {

			if grid[row][col] == 1 {
				//we found an island lets dfs into it and get the count

				//area := dfs(row, col, grid, directions)
				area := bfs(grid, directions, Coords{row, col})
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

//since this problem is allllll about local locality, such as a graph or tree, if the nieghbor or target nodes are CLOSE, to the current node
//then BFS is faster at searching a local locality, DFS searching DEEPER, but if soltuions or cells are not so far deep than they are wide and closer
//it would seem to make more sense to search with breadth, HOWEVER our algo really acts like a BFS because it stops once it hits boundaries... hmmm
//so i think we are fine...?

//lets try bfs anyways

type Coords struct {
	Row int
	Col int
}

//returns the max area of the island, given a row and column of a land point
func bfs(grid [][]int, directions [][]int, coord Coords) int {
	//
	q := make([]Coords, 1)
	q[0] = coord
	localArea := 0
	for len(q) > 0 {

		//check the current coords as visited
		//dp, optmization
		item := q[0]

		//clean up?
		q = q[1:]
		//increment, since we are currently on ONE peice of land
		grid[item.Row][item.Col] = 0

		localArea++
		//check its neighbors
		//conditions again.. LEFT, TOP, RIGHT, DOWN (horizontal, vertical)
		for _, dir := range directions {

			//dir[0] = X(Row) dir[1] = Y(Col)
			//conditions to enqueue,
			//If our Neighbor(ADJNODE/CELL) are IN BOUNDARY AND our Neighbor is an 1 (one)
			var adjRow int = item.Row + dir[0]
			var adjCol int = dir[1] + item.Col
			if adjRow >= 0 && adjRow <= len(grid)-1 && adjCol >= 0 && adjCol <= len(grid[0])-1 && grid[adjRow][adjCol] == 1 {
				//we KNOW that this current neighbor cell IS a 1(one) AND within boundary, lets enqueue it and search ITS neighbors
				q = append(q, Coords{item.Row + dir[0], item.Col + dir[1]})
			}
		}

	}

	return localArea

}

func dfs(row int, col int, grid [][]int, directions [][]int) int {
	//exit case...
	//if im out of bounds, that's water
	//if the current cell is  zero(0)
	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == 0 {
		//we are out of bounds
		//return -1//don't need to found this as a valid cell
		return 0
	}
	//mark visited
	grid[row][col] = 0
	//otherwise we are currently sitting inside a valid 1 cell
	//we can search its neighbors
	maxArea := 1
	for _, dir := range directions {
		//return 1 + dfs(row + dir[0], col + dir[1], grid, directions)
		maxArea += dfs(row+dir[0], col+dir[1], grid, directions)
	}
	//maxArea  = maxArea + 1 +
	//return 1
	return maxArea
}

func modArea(a int) int {
	if a != 100 {
		//v:= a++ never undersatnd increment

		v := a + 1
		return 1 + modArea(v)
	}
	return 1
}
