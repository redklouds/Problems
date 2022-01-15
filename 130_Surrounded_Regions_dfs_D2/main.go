package main

//very slow
//Runtime it seems, is O(n^2) because we will go over it a couple times
//Space ... is going to be O(n) recursion
func solve(board [][]byte) {
	//idea mark all the edges as islands we cannot touch, using DFS we can mark all peice that we cannot touch

	//then run through it again, this time we know that all the O's that we see are to be marked to X since they aren't an edge or connected to an edge side

	//ASSUMPTION ANY O that is NOT toucing edges will be surrounded by X's

	//so DFS the edges , marking the infected islands
	if len(board) < 3 {
		//less than three forget even considering it you cannot have 2 rows that are surrounded
		return
	}

	directions := [][]int{
		//ROW, COL , ROW=UP/DOWN, CoL = Right/Left
		{-1, 0}, //up
		{1, 0},  //down
		{0, -1}, //left
		{0, 1},  //right
	}
	//we have at least 3 rows

	//run dfs on the top and bottom rows
	for colIdx := range board[0] {
		//running for top and bottom ROWS meaning we are checking each column of the first and last rows
		dfs(board, directions, 0, colIdx)
		dfs(board, directions, len(board)-1, colIdx)
	}

	//run dfs on the left and right columns
	for rowIdx := range board {
		dfs(board, directions, rowIdx, 0) //first column
		dfs(board, directions, rowIdx, len(board[0])-1)
	}

	//then finally we want to just convert our findings

	//ledger if we encounter a 'A' -> 'O'
	//if we see an 'O' -> 'X'
	for row := range board {
		for col := range board[0] {
			if board[row][col] == 'A' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}

		}
	}
}

func dfs(board [][]byte, directions [][]int, row int, col int) {

	//condition -  we want to use this method to 'mark' infected O's that are connected to an edge, meaning these are islands that are NOT 4 side surrounded by X's

	//run this dfs on the edges, the outter edges

	//edge cases,

	//if its a 2 rows meaning it only has 2 rows and N columns, it DOESN"T matter we won't flip shit, since, you cannot have a field surrounded by X's with 2 rows

	//MUST be at least 3 rows  for someting to be surrounded

	//incase its 3 or more rows, after we dfs the edges ,we then iterate the inner starting r +1 (inn) to len(r-1) same for columns

	//if out of bounds we do not continue the recursive calls

	//ledger
	//if we encounter an O -> mark it an 'A'
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 || board[row][col] == 'A' || board[row][col] == 'X' {
		return
	}
	//at this point we landed on a cell that is WITHIN boundaries, has NOT been visited that was once an CIRCLE AND is NOT an 'X' must be an 'O'
	//search neighbors for O

	board[row][col] = 'A' //mark the current
	for _, dir := range directions {
		dfs(board, directions, row+dir[0], col+dir[1])
	}

	//we are within BOUNDS and landed on a X or a O

}
