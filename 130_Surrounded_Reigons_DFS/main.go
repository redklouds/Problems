package main

import "fmt"

var isConfined = false

func solve(board [][]byte) {

	visited := make([][]bool, len(board))

	//initalize bisited map
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	//now we need to say for row,
	//for col

	//if the current cell is a 'o' lets search it
	//call dfs on it,
	directions := [][]int{
		//X,Y
		{0, -1}, //UP
		{1, 0},  //RIGHT
		{0, 1},  //DOWN
		{-1, 0}, //LEFT
	}
	//our Bondary, we WANT to only search from a level within the boundary outside, and search from inside out, to see
	//if there is a path that leads to the edge of the boundary of 'o's
	for row := 1; row < len(board)-1; row++ { //only going to iterate from row 1 to row n - 2 (index)
		for col := 1; col < len(board[0])-1; col++ { //only going to search a boundary box, that is not including the outter bondary
			//these are indexes
			//if we found an 'O'
			if board[row][col] == byte('O') && !visited[row][col] {
				//dfsCircleFind(row, col, board, visited, directions)
				isConfined := false
				FindIslandsSurroundedByX(row, col, board, visited, directions, &isConfined)
				if isConfined {
					//means we found a region that is confined, lets go and change all the regions values to 'X'
					dfsChangeRegionValuesToX(row, col, board, directions)
				}
			}
		}
	}
}

func dfsChangeRegionValuesToX(row int, col int, board [][]byte, directions [][]int) {
	if board[row][col] == 'X' {
		return
	}
	//we assme its within the bondaries of X
	//if its not an X it must be an O
	board[row][col] = 'O'

	//now we must search
	for _, dir := range directions {
		dfsChangeRegionValuesToX(row+dir[0], col+dir[1], board, directions)
	}
}

//if it touches a boundary (OUTSIDE the outskirts) we can return true, signifying it touched outside, and is NOT
//contained within x's
func FindIslandsSurroundedByX(row int, col int, board [][]byte, visited [][]bool, directions [][]int, isConfined *bool) {
	//if the current cell is an 'X' lets return
	if board[row][col] == byte('X') {
		*isConfined = true
		return
	}
	//else since we know we are just searching from within a box, and not the boundaries
	//it must be a 'O'
	//we have to see if the 'o' lives on the bondary
	if row == 0 || row == len(board)-1 || col == 0 || col == len(board[0])-1 {
		//we are an 'o' AND on the bondary
		*isConfined = false
		return
	}

	//else lets search it since we are still within bondary
	//lets mark it as visited
	visited[row][col] = true
	for _, dir := range directions {
		if !visited[row+dir[0]][col+dir[1]] {
			FindIslandsSurroundedByX(row+dir[0], col+dir[1], board, visited, directions, isConfined)
		}

	}
}

func FindFuckingXsV2(grid [][]byte) {
	//actually its quite easy, we know that the boundary is crap, we honestly can simply DFS JUST like the pacific Ocean Rain problem
	//we can just DFS from the  OUTSIDE parameter, and DFS from OUTSIDe INWARDS MARKING all the o's that shouldn't becoime X's because they have a path
	//from the OUTSIDE/Boundary INWARDS (THINK Pacific Alantic Ocean problem) whats to say we cannot just dfs from outside in and mark the  nodes/cells that
	//are ineligable for becoming ex's, so basically i Mark lets say.. with '1' on cells that are ineligable to become x'S, and remember the ONLY time a O should
	//be an X is WHEN ITS NOT Connected to an O that is OUTSIDE to the boundary, so if i can find all the O's that are ineligable ALL O's left will be eligilible

	//so first traverse the outside boundary, using DFS when you see a O, dfs inwards, marking all nodes that are o's as '1's

	//then traverse normalls through the matrix, for all '1's change them back to O's and for O's encountered change them to 'X's easy, O(MN)

	//think in another way to solve the issue, think BACKWARDS, why solve an issue forward thinking, think about it backwards, solve from INSIDE out
	//relentlyless ask and attack the problem, can i approach it from outside?
	//can i approch it starting from th ecenter?

	//lets go around the outside permiter and dfs INTO the map, if we find an o
	directions := [][]int{
		//X,Y
		{0, -1}, //UP
		{1, 0},  //RIGHT
		{0, 1},  //DOWN
		{-1, 0}, //LEFT
	}
	//just iterate the top and bottom rows
	for col := 0; col < len(grid[0]); col++ { //we are iterating the first and last rows But that means basically the COLUMNS of these rows
		if grid[0][col] == byte('O') {
			//we have an entry point lets dfs into it and mark these cells
			dfsSet(0, col, grid, directions)
		}

		if grid[len(grid)-1][col] == byte('O') {
			//we have found an entry on the last row we should dfs into and see if we have paths
			dfsSet(len(grid)-1, col, grid, directions)
		}
	}

	//just iterate the left and right-most columns

	for row := range grid {
		//for each row, in the first and last columns
		//for the left-most first columnn
		if grid[row][0] == byte('O') {
			dfsSet(row, 0, grid, directions)
		}

		//for the last column (right-most)
		if grid[row][len(grid[0])-1] == byte('O') {
			dfsSet(row, len(grid[0])-1, grid, directions)
		}
	}

	//now that we have DFS's the outter boundary lets go ahead and
	//make the magic
	PrintMatrix(grid)
	for row := range grid {
		for col := range grid[0] {
			//if the cell is '1' turn it into an 'O'
			//if its 'o' turn it into a 'X'

			if grid[row][col] == byte('O') {
				grid[row][col] = 'X'
			}
			if grid[row][col] == byte('1') {
				grid[row][col] = 'O'
			}
		}
	}
	//we will have ------
	/*
		_____________
		|			|
		|			|
		|___________|
		iterated the boundary
	*/

}

func dfsSet(row int, col int, grid [][]byte, directions [][]int) {
	if row < 0 || row > len(grid)-1 || col < 0 || col > len(grid[0])-1 || grid[row][col] == byte('X') || grid[row][col] == byte('1') {
		//out of bounds
		//equals an 'X'
		//exit case, otherwise it equals an 'O'
		//or its somthing we've 'visited' already
		return
	}
	//we aren't out of bounds
	//so lets continue to mark and search
	grid[row][col] = '1'
	for _, dir := range directions {
		//dir[0] = X dir[1] = Y
		dfsSet(row+dir[0], col+dir[1], grid, directions)
	}

}

func main() {
	/*
		testData := [][]byte{
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
		}

		solve(testData)
		PrintMatrix(testData)
	*/

	/*
		testData2 := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}

		solve(testData2)
		PrintMatrix(testData2)
	*/
	/*
		testData3 := [][]byte{
			{'X', 'X', 'X', 'O', 'X', 'X'},
			{'X', 'O', 'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X', 'X', 'O'},
			{'X', 'O', 'O', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'X'},
		}

		FindFuckingXsV2(testData3)
		PrintMatrix(testData3)
	*/
	testData4 := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}

	FindFuckingXsV2(testData4)
	PrintMatrix(testData4)
}

func modB(bal *bool) {
	*bal = true
}

func PrintMatrix(mat [][]byte) {
	for i := range mat {
		for j := range mat[0] {
			fmt.Printf("%c ", mat[i][j])
		}
		fmt.Println("")
	}
}

//dfsCicleFind
