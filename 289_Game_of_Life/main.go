package main

func gameOfLife(board [][]int) {

	mat := make([][]int, len(board))

	//popualte th board

	for i := range mat {
		mat[i] = make([]int, len(board[0]))
	}

	//basicallcol we just asking all directions

	directions := [][]int{
		//8 directions
		//row(Y), col(X)
		{-1, 0},  //TOP
		{-1, 1},  //TOP-RIGHT
		{0, 1},   //RIGHT
		{1, 1},   //RIGHT-BOTTOM
		{1, 0},   //BOTTOM
		{1, -1},  //LEFT-BOTTOM
		{0, -1},  //LEFT
		{-1, -1}, //TOP-LEFT
	}

	//for each direction we check

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			//we check each directionm
			//num neibors
			numNeighbors := 0

			for _, dir := range directions {
				//for each direction dir[0] = row, dir[1] = col
				//check OOB
				if row+dir[0] >= 0 && row+dir[0] <= len(board)-1 && col+dir[1] >= 0 && col+dir[1] <= len(board[0])-1 {
					//we are within boundaries lets update
					//our rules

					//rule 1 if we are 1, and we have less than 2 neighbors , we become 0

					//rule 2 if we are 1, and we have 2 or 3 neighbors, we get to stacol/skip/become 1

					//rule 3 if we are 1 and we have more than 3 neibhor, we turn to 0

					//rule 4 if we are 0, and we have erowactlcol 3 neighbpor we become 1
					if board[row+dir[0]][col+dir[1]] == 1 || board[row+dir[0]][col+dir[1]] == 2 {
						numNeighbors++
					}
				}
			}

			if board[row][col] == 1 || board[row][col] == 2 { //means I WAS a 1  now im suppoused to be a zero afterwards
				if numNeighbors < 2 {
					board[row][col] = 2 //mat[row][col] = 2//0
				} else if numNeighbors == 2 || numNeighbors == 3 {
					board[row][col] = 1 //mat[row][col] = 1//1 can this case go away? if im a one i stay a one
				} else if numNeighbors > 3 {
					board[row][col] = 2 //mat[row][col] = 2//0
				}
			}
			if board[row][col] == 0 {
				if numNeighbors == 3 {
					board[row][col] = -1
					//mat[row][col] = -1//1
				}
			}
		}
	}

	//in place

	//2 says "I was a 1 now ima Zero!!!"
	//-1 Says "I was a zero Now I'ma One!"
	//for all 2s i find turn them into ZERO
	//for all -1 i find turn them into a 1
	for r := range mat {
		for c := range mat[0] {
			if board[r][c] == 2 {
				board[r][c] = 0
			}
			if board[r][c] == -1 {
				board[r][c] = 1
			}

		}
	}
}
