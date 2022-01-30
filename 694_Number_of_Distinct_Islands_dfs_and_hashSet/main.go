package main

import (
	"fmt"
	"strings"
)

var DIRECTIONS = [][]int{
	//ROW/COL = UP/DOWN | LEFT/RIGHT
	{0, -1}, //LEFT
	{-1, 0}, //UP
	{0, 1},  //RIGHT
	{1, 0},  //DOWN
}

//O(n) - > Fastest using a strinb buiolder this is STRING HEAVY, so use strinb builder to optmize
//we can also think about reusing the strinb builder, and a single hashset
func numDistinctIslands(grid [][]int) int {
	//the hash structure, has to be a island coord hash,
	//you can pass a strinb uilder object around

	//but lets make it simple - we have a array of strings, that become our hash

	hashSet := make(map[string]bool)
	//break down problem

	var sb strings.Builder
	//1: dfs to collect all islands
	for row := range grid {

		for col := range grid[row] {
			if grid[row][col] == 1 {

				//2:when we encounter an island, save its top left coords(err the first coords) we want to baseline them to the (0,0) moving all objects to this point
				//islandCoords := make([]string, 0)
				//3: collect all the island coords, and store them into a hash object
				dfs(row, col, row, col /*, &islandCoords*/, grid, &sb)
				//4: after collecting the hash coord string, store them into our global hashSet, islands that are the same (dups) will be removed
				//hashSet[strings.Join(islandCoords, "")] = true
				hashSet[sb.String()] = true
				sb.Reset()
			}
		}
	}
	//5: count the number of island coords inside the hashset
	return len(hashSet)
}

//1: dfs to collect all islands
//if not OOB we want to build a string, that repersents if this island was shifted to the top left corner of the grid (0,0)
func dfs(row int, col int, topLeftRow int, topLeftCol int /* islandCoords *[]string,*/, grid [][]int, sb *strings.Builder) {

	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == 0 {
		return
	}
	//found a valid island

	fmt.Fprintf(sb, "%d_%d", (row - topLeftRow), (col - topLeftCol))
	//s := fmt.Sprintf("%d_%d", (row - topLeftRow), (col - topLeftCol))

	//*islandCoords = append(*islandCoords, s)
	grid[row][col] = 0 //mark visited
	for _, dir := range DIRECTIONS {
		dfs(row+dir[0], col+dir[1], topLeftRow, topLeftCol /*islandCoords,*/, grid, sb)
	}
}
