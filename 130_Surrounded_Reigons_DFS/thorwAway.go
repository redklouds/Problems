package main

func dfsCircleFind(row int, col int, board [][]byte, visited [][]bool, directions [][]int) int {
	//couple cases, if we out of bounds lets return
	/*if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0]) - 1{
	    //remember we are playting with lengths here SIZES, and indexes don't forget the MINUS 1
	    return //we are out of bounds
	}*/
	//actually do we need to go so far to search the boarder?
	//if we are here we know we are within bounds

	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0]) {
		return 1 //the ONLY reason that i would gone out of range is if i encountered a '0' on the edge
	}

	if row == 0 || row == len(board)-1 || col == 0 || col == len(board[0]) && board[row][col] == byte('O') {
		//we are on the boarder and we found an O, we return false
		return 1
	}

	//now we know that we are not on the boarder, and within bounds, lets check if its an X
	if board[row][col] == byte('X') {
		//we found an X, return true, we signafy our search has found a surrounding
		return -1
	}

	//if we have not found an X and landed ourselves on an 'O' lets mark it as visited
	visited[row][col] = true

	//and search its adjacent nodes

	flip := -1
	for _, dir := range directions {
		flip *= dfsCircleFind(row+dir[0], col+dir[1], board, visited, directions)
		/*
		   if flip {
		       board[row][col] = 'X'
		   }
		*/
	}
	return -1
}
