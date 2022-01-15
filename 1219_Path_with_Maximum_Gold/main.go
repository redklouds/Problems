package main

func getMaximumGold(grid [][]int) int {
	/*
	   It is still necessary to use recursive thinking to solve the problem,
	   each node can only go up, down, left, and right.
	   If the gold at that point is 0, then you cannot go to that point.
	   After withdrawing money at this point, go to the next node.
	   Record the total gold amount at this time and update the maximum
	   value of the total gold amount.
	   To restore the gold value of the node for the next recursive search.
	*/
	directions := [][]int{
		//ROW, COL UP/DOWN, LEFT/RIGHT
		{0, -1}, //LEFT
		{-1, 0}, //UP
		{0, 1},  //RIGHT
		{1, 0},  //DOWN
	}
	maxGold := 0
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] != 0 {
				tempGoldAmt := grid[row][col]
				grid[row][col] = 0
				dfs(directions, grid, row, col, tempGoldAmt, &maxGold)
				grid[row][col] = tempGoldAmt //this reverts our dfs search
			}
		}
	}
	return maxGold
}

func dfs(directions [][]int, grid [][]int, row int, col int, gold int, maxGold *int) {
	*maxGold = Max(*maxGold, gold)
	for _, dir := range directions {
		nextRow := row + dir[0]
		nextCol := col + dir[1]

		//we want to only search WITHIN the boundaries of the board
		if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) && grid[nextRow][nextCol] != 0 {
			tempGold := grid[nextRow][nextCol]
			grid[nextRow][nextCol] = 0
			dfs(directions, grid, nextRow, nextCol, gold+tempGold, maxGold)
			grid[nextRow][nextCol] = tempGold
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
