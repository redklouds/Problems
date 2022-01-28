package main

//O(m*n) solution, where we use DFS and BFS
//we NEED TO THINK about this
//think about it.,. if we want to find the shortest distance.. that immediately is BFS.. however we need to find which EDGE is the shortest... the way we can actually do
//this is not to record the edges themselves... but to

func shortestBridge(grid [][]int) int {
	// 1 use DS to mark the first island, and enqueue it, we also change island 1 as
	// and mark as 2 to mark as visited node,

	var island1Coords []*Coords

	firstOne := firstIslandCoord(grid)
	dfs(firstOne.Row, firstOne.Col, &island1Coords, grid)

	//2 preform BFS on the visited nodes, level by level, and expand by 1 step each level, KNOWNING
	//that BFS will search around the island 'level by level'

	distance := 0
	for len(island1Coords) > 0 {
		//perform level by level traversal known as bounday traversal in Matrix
		qLen := len(island1Coords)
		for i := 0; i < qLen; i++ {

			coord := island1Coords[i]

			//if we see a 1/other island return the distance
			for _, dir := range directions {
				adjRow := coord.Row + dir[0]
				adjCol := coord.Col + dir[1]
				if adjRow >= 0 && adjRow < len(grid) && adjCol >= 0 && adjCol < len(grid[0]) && grid[adjRow][adjCol] != 2 {
					if grid[adjRow][adjCol] == 1 {
						return distance
					}
					grid[adjRow][adjCol] = 2
					island1Coords = append(island1Coords, &Coords{adjRow, adjCol})
				}
			}
		}
		//after each level, we know we have completed a step or walked a aroudn teh boundary
		distance++

	}

	//if we find the second island return the steps/level immediately
	return distance
}

func firstIslandCoord(grid [][]int) Coords {
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 1 {
				return Coords{row, col}
			}
		}
	}
	return Coords{}
}

//mark and querue in the first island
func dfs(row int, col int, island1Coords *[]*Coords, grid [][]int) {

	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == 2 || grid[row][col] == 0 {
		//oob
		return
	}

	//we are on a one and within bounds
	//mark as 2, and enqueue to the queue
	coord := &Coords{row, col}

	*island1Coords = append(*island1Coords, coord)
	grid[row][col] = 2
	for _, dir := range directions {
		dfs(row+dir[0], col+dir[1], island1Coords, grid)
	}
}
