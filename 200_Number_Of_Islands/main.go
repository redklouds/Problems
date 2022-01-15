package main

func numIslands(grid [][]byte) int {

	directions := [][]int{
		//Row, Col -> (Row=Up/Down, Col="Left/Right")
		{0, -1}, //LEFT
		{-1, 0}, //UP
		{0, 1},  // RIGHT
		{1, 0},  // DOWN
	}

	//IDEA:
	//1: Find the first spot of land
	//2: Run DFS on that plot of land to search it through, running as we connect other lands together to make an island
	//3: each time we finish running DFS on a land we record that as a island (island can be a single index/spot)

	numIsland := 0
	//search entire grid
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == '1' {
				//search it
				dfs(grid, row, col, directions)
				numIsland++
			}
		}
	}
	//if found island, lets search it, and since we can alter the grid, lets mark the island as we walk it
	return numIsland
}

//dfs - we do not need to return anything since we are simply walking the distance of this respective island
//and marking the ones we visit as water,
//ASSUMPTION: once we walk a land , we know that's considered an island
func dfs(grid [][]byte, row int, col int, directions [][]int) {

	//bounday control
	//OOB and/or hit water, cancel this recursive step

	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == '0' {
		return
	}

	//we know we are on a peaice of land lets search its neighbor
	grid[row][col] = '0'
	for _, dir := range directions {
		dfs(grid, row+dir[0], col+dir[1], directions)
	}
}
