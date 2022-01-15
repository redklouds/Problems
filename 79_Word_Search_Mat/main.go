package main

func exist(board [][]byte, word string) bool {

	//RUN DFS like when seraching for somethjing in a matrix,

	//idea,

	//search the board for the first letter, if found, we want to run DFS on the board

	//if we find the first cahrater run DFS on that starting board, if we hit visited BOX OR we run OOB stop that recurse, if we find a char on the next we continue
	directions := [][]int{
		//ROW, COL | ROW="UPDOWN" | COL="LEFT/RIGHT"
		{0, -1}, //LEFT
		{-1, 0}, //UP
		{0, 1},  //RIGHT
		{1, 0},  //DONW
	}

	found := false
	//search the entire grid to look for an entry point

	//convert the string word into a rune, so we access the word as a run array

	//wordArr := []rune(word)
	wordArr := []byte(word)
	for row := range board {
		for col := range board[0] {
			if board[row][col] == wordArr[0] {
				vis := CreateVisitedMat(len(board), len(board[0]))
				//we have a match of the first char, lets run dfs on this entry point and see if we can find a match for the sequence
				found = dfs(board, wordArr, 0, row, col, directions, vis)
				if found {
					return found
				}
			}
		}
	}

	return found
}

func dfs(board [][]byte, word []byte, curCharIdx int, row int, col int, directions [][]int, visited [][]bool) bool {
	//edge cases if we hit a edge or visited this before, then we exit the search
	//if we iterate untiil we are over the current index
	if curCharIdx > len(word)-1 {
		return true
	}

	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 || visited[row][col] || board[row][col] != word[curCharIdx] {
		return false
	}

	isFound := false
	visited[row][col] = true
	curCharIdx++

	for _, dir := range directions {
		//WE NEED TO USE THE OR statement here, because its true OR false, if we use a && we will chain the true and false which becomes the false
		f := dfs(board, word, curCharIdx, row+dir[0], col+dir[1], directions, visited)
		isFound = isFound || f
	}
	/*
		//check if this cell is one that is apart of the seqeuence
		if word[curCharIdx] == board[row][col] {
			//meaning that the current cell equals what we are looking for search its neighbors
			//if we are not OOB or visited, and we are not at the end
			isFound := true
			visited[row][col] = true
			curCharIdx++

		}
	*/

	//else
	//we need to

	//if we are in boundaries, is this cell what we are searching for?

	//if it is we want to search for the next char in its containing neighbors

	//
	return isFound
}

func CreateVisitedMat(n, m int) [][]bool {
	v := make([][]bool, n)
	for i := range v {
		v[i] = make([]bool, m)
	}
	return v
}
