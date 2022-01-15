package main

import "math"

func minPathSum(grid [][]int) int {
	dpMat := make([][]int, len(grid))
	//setup
	for i := range dpMat {
		dpMat[i] = make([]int, len(grid[0]))
	}

	directions := [][]int{
		{0, 1}, //RIGHT
		{1, 0}, //DOWN
	}

	for r := len(grid) - 1; r >= 0; r-- {
		for c := len(grid[0]) - 1; c >= 0; c-- {
			if r == len(grid)-1 && c == len(grid[0])-1 {
				dpMat[len(grid)-1][len(grid[0])-1] = grid[len(grid)-1][len(grid[0])-1]
			} else {
				//bounds checks
				celMin := math.MaxInt32
				for _, dir := range directions {
					if r+dir[0] >= 0 && r+dir[0] <= len(grid)-1 &&
						c+dir[1] >= 0 && c+dir[1] <= len(grid[0])-1 {
						//within boundaries now we must get the mimum between my two paths

						//ME plus my neighbor who is smaller? meaning who will i take
						celMin = Min(celMin, grid[r][c]+dpMat[r+dir[0]][c+dir[1]])
					}
				}
				dpMat[r][c] = celMin
			}

		}
	}

	return dpMat[0][0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
