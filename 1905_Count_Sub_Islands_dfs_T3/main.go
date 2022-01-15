package main

func main() {

	//idea is to invert the dfs, so that we don't return something, we ONLY

	//base case is ONLY return false if a cell exists in grid2 and NOT grid1

	//return TRUE if we can continue traversing
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	visited := make([][]bool, len(grid2))
	for i := range visited {
		visited[i] = make([]bool, len(grid2[0]))
	}

	directions := [][]int{
		//ROW, COL = UP/DOWN, LEFT/RIGHT
		{0, -1}, //LEFT
		{-1, 0}, //UP
		{0, 1},  //RIGHT
		{1, 0},  //DOWN
	}

	//NOTE : I can pass the matrix visited in because its ALREADY allocated, and im NOT
	//Changing the mat header so the values will be kept
	count := 0
	for row := range grid2 {
		for col := range grid2[0] {
			if !visited[row][col] && grid2[row][col] == 1 {
				if dfs(row, col, grid1, grid2, directions, visited) == 1 {
					count++
				}
			}

		}
	}
	return count
}

func dfs(row int, col int, grid1 [][]int, grid2 [][]int, directions [][]int, visited [][]bool) int {

	///the idea to solve this , the question was.. HOW DO we iterate over the sub island terrain, without returning false before
	//we have finished all the islands?
	//var isValid bool
	if row < 0 || row > len(grid1)-1 || col < 0 || col > len(grid1[0])-1 || grid2[row][col] == 0 || visited[row][col] {
		// oob or hit water, we return false, REMEMBER THAT SO LONG AS A SINGLE TRUE EXISTS, our statement is true
		return 0 //AS LONG AS A SINGLE STACK RETURN TRUE WE ARE GOOD, SO CATCH ALL THE CASES THAT ARE FALSE
	}
	//if its out of boundaries or grid2 spot is water return false

	//if grid2 is land and grid1 is water return false
	//we know grid2 is 1
	if grid1[row][col] == 0 {
		return -1
	}
	visited[row][col] = true
	//if we made it past this we are ON a cell that grid1 and grid2 are 1, think about "SQUEEZING the condition" so that only a single true will be returned to the stack
	//visited OR use grid2, since we aren't going to need to process this
	/*
		for _, dir := range directions {
			isValid := dfs(row+dir[0], col+dir[1], grid1, grid2, directions, visited)
			if isValid == -1 { //does any of these calls find a cell that g2 has that g1 doesn't
				return -1
			}
		}
	*/
	//dont' know wht the hell is the difference 

	a := dfs(row+directions[0][0], col+directions[0][1], grid1, grid2, directions, visited)
	b := dfs(row+directions[1][0], col+directions[1][1], grid1, grid2, directions, visited)
	c := dfs(row+directions[2][0], col+directions[2][1], grid1, grid2, directions, visited)
	d := dfs(row+directions[3][0], col+directions[3][1], grid1, grid2, directions, visited)
	if a == -1 || b == -1 || c == -1 || d == -1 {
		return -1
	}

	return 1 //SO AS LONG AS ONE OF THE Stacks return true
}

/*
func dfs_old(row int, col int, grid1 [][]int, grid2 [][]int, directions [][]int, visited [][]bool) bool {

	///the idea to solve this , the question was.. HOW DO we iterate over the sub island terrain, without returning false before
	//we have finished all the islands?

	var isValid bool
	//so if we are OOB or we hit zero lets not continue
	if row >= 0 && row < len(grid1)-1 && col >= 0 && col < len(grid1[0])-1 && grid2[row][col] != 0 {
		//we are within boundaries and NOT on water, AND THIS CURRENT CELL IS a 1 on grid2

		if grid1[row][col] == 1 && grid2[row][col] != 1 {
			return false
		}

		//lets recurse
		visited[row][col] = true
		isValid = true
		for _, dir := range directions {
			isValid = isValid || dfs(row+dir[0], col+dir[1], grid1, grid2, directions, visited)
		}

	}
	return isValid
}
*/
