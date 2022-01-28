package main

import "math"

type Coords struct {
	Row int
	Col int
}

var directions = [][]int{
	//ROW, COL UP/DOWN, LEFT/Right
	{0, -1}, //LEFT
	{-1, 0}, //UP
	{0, 1},  //RIGHT
	{1, 0},  //DOWN
}

func shortestBridge_v1(grid [][]int) int {
	//the idea is to

	//2 soltuions 1: get ALLL COOORDINATES of the two island, then run a manhattan dsitance tog et teh SMALLEST distance between the two island, this will answer the question, "how many fips are required to connect to the two"

	//answering that this is the smallest number of distant cells between the two

	//brute force, O(N^2) becasue you will to break the problem into a couple steps

	//1: find and get a list of all coords of the first island

	//2: find and get a list of all the coords of the second island

	//3: apply a distance formula (manhattan distance) since we know it is only four direction

	//4: return the smallest distance

	//all coordinates of island1
	var island1Coords []*Coords

	var island2Coords []*Coords

	island1Done, island2Done := false, false
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 1 && !island1Done {
				//we found some island lets get all coords
				//we know that after we run this function all the rest of the cells will be zero, so we can continue where we left off
				getCoordsIfIslandDfs(row, col, grid, &island1Coords)
				island1Done = true
			}
			//after we finish lets continue to island2 - Optimize we can actually move it forward if we want to the last coords, however we it won't change the run time
			if grid[row][col] == 1 && !island2Done {
				getCoordsIfIslandDfs(row, col, grid, &island2Coords)
				island2Done = true
			}
		}
	}

	//we have the coords for the two island and all their coords, let run the
	//distance fornumula for each pair of coords to find the smallest distances

	minDistance := math.MaxInt32
	for _, coord1 := range island1Coords {
		for _, coord2 := range island2Coords {
			minDistance = Min(minDistance, distance(*coord1, *coord2))
		}
	}
	return minDistance
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//if i may modify the grid i can use it as a visited grid and not need to allocate a visited map
func getCoordsIfIslandDfs(row int, col int, grid [][]int, islandCoords *[]*Coords) {
	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == 0 {
		return
	}
	//else we are in bounday and the current cell is a land
	*islandCoords = append(*islandCoords, &Coords{row, col}) //add the current coords

	//flip this cell as visited
	grid[row][col] = 0
	for _, dir := range directions {
		getCoordsIfIslandDfs(row+dir[0], col+dir[1], grid, islandCoords)
	}
}

//given two coordinates perform the manhattan distance formula
func distance(coords1, coords2 Coords) int {
	//dist = Abs(x2 - x1) + Abs(y2 - y1)
	return (Abs(coords2.Col-coords1.Col) + Abs(coords2.Row-coords1.Row)) - 1
}

//if val is negative return inverse of it
func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
