package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	//grid := [][]int{{0}, {1}}
	maxArea := maxAreaOfIsland(grid)

	fmt.Println(maxArea)
}
func maxAreaOfIsland(grid [][]int) int {
	//basically we want to find the LARGEST island here, so

	//1: first find the island
	//2: we need to run a DFS on that start of the island to verify wheather its an island
	//3: collect its size, after we finish running dfs on it, we return the size we have observed
	//4: the edges doesn't effect our count we just want absolute size

	//edge cases grid is smaller? does that handle it? [[0],[1]]

	//create our beautiful search directions
	// Row,Col,  moving Rows makes you go up and down
	//moving columns make you goLeft and Right
	dir := [][]int{
		//ROW, COL (U/D, L/R)
		{-1, 0}, //Go Up
		{0, -1}, //Go LEFT
		{1, 0},  //Go DOWN
		{0, 1},  //Go RIGHT
	}

	/*
		visited := make([][]int, len(grid))
		for i := range visited {
			visited[i] = make([]int, len(grid[0]))
		}
	*/

	//search for the first peice of island
	maxArea := 0
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 1 {
				//curCoords := Coords{col, row} //X is the column, col controls the X, Row's control the Y
				area := dfs(dir, grid, row, col)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

type Coords struct {
	X int
	Y int
}

//we want to run dfs, on the first sighting of an island, and return the observed size
func dfs(directions [][]int, grid [][]int, row int, col int) int {
	//our rules for searching
	//if we hit an edge (OOB) OR we hit an already visited land, OR we hit a zero, we want to STOP our iterations
	//stopping our search for MORE LAND TO CONCOURQQQQ

	//check oob if visited, and if its not an island
	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == 0 {
		return 0
	}

	//check if we already visited the current node

	//we are now within bounds, not visited, and we have hit a circle

	//WE KNOW THIS IS A 1 so we can count here..

	grid[row][col] = 0
	area := 1
	for _, dir := range directions {
		//search the other directions and mark this one visited
		nextRow := row + dir[0]
		nextCol := col + dir[1]
		area += dfs(directions, grid, nextRow, nextCol)
	}

	return area
}
