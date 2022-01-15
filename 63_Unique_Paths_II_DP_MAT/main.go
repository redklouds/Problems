package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 {
		if obstacleGrid[0][0] == 1 {
			return 0
		}
		return 1
	}
	//the idea is JUST like uniquePath,s however, we know there are obsticles
	//so if there are we just don't count those paths as one of our potential path

	dpBoard := make([][]int, len(obstacleGrid))

	//initalize board
	for i := range dpBoard {
		dpBoard[i] = make([]int, len(obstacleGrid[0]))
	}

	directions := [][]int{
		//ROW, COL
		{0, 1}, //RIGHT
		{1, 0}, //DOWN
	}

	//we know that there is ONE way to get to the bottom corner as our default case
	dpBoard[len(obstacleGrid)-1][len(obstacleGrid[0])-1] = 1

	//now we can iterate from the bottom corner

	for r := len(dpBoard) - 1; r >= 0; r-- {
		for c := len(dpBoard[0]) - 1; c >= 0; c-- {

			//ok for every direction if its valid directions lets take that direction
			//and add it to our list of possible paths
			//if not oob, and the direction is NOT a obstacle

			if obstacleGrid[r][c] != 1 { //EDGE CASE WHAT DO WE DO IF I AM A OBSTACLE? no reason to calculate it
				for _, dir := range directions {
					if r+dir[0] >= 0 && r+dir[0] <= len(dpBoard)-1 &&
						c+dir[1] >= 0 && c+dir[1] <= len(dpBoard[0])-1 &&
						obstacleGrid[r+dir[0]][c+dir[1]] != 1 {
						//we are in bounds AND our neighbors are NOT an obstacle so we can take their paths
						dpBoard[r][c] += dpBoard[r+dir[0]][c+dir[1]]
					}
				}
			}

		}
	}

	return dpBoard[0][0]
}
