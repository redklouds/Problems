package main

//IDEA , we want to traverse each direciton (8) until we either hit a Queen OR
//hit out of boundary

//so start at the king,
//and have a map of each index,
//map Key = is row value is colummn, so when se

//optimizes the searching if a value exists in this particular field
//USING A ROW AND COL array
//each index of the array repersents if the current row and col exists a queen, this prevents us from iterating to find
//a pair of coords and helps optmize look up speed to jsut map[curRow][curCol]

func queensAttacktheKing(queens [][]int, king []int) [][]int {

	//first populate the row and column arrays

	//we KNOW tghat its a 8X8 chessboard, so the most we have is index 7

	locationMap := make([][]bool, 8)
	for i := range locationMap {
		locationMap[i] = make([]bool, 8)
	}
	//populate the rows and columns that contain somthing

	for _, queenPair := range queens {
		locationMap[queenPair[0]][queenPair[1]] = true
	}

	//make the directions map so we can call our run method, 8 directions

	directions := [][]int{
		//ROW, COL = ROW="up/Down", Col = "Left/Right"
		{0, -1},  //LEFT
		{-1, -1}, //LEFT/UP
		{-1, 0},  //UP
		{-1, 1},  //UP/RIGHT
		{0, 1},   //RIGHT
		{1, 1},   //RIGHT/DOWN
		{1, 0},   //DOWN
		{1, -1},  //DOWN/LEFT
	}

	var result [][]int
	for _, dir := range directions {
		//check if a queen is found in this direction, if so then we want to record it
		if found, coords := run(dir, locationMap, king); found {
			result = append(result, coords)
		}
	}

	return result
}

func run(direction []int, queenLocationmat [][]bool, kingCoords []int) (isFound bool, coords []int) {
	//our boundaries, we
	//either we are out of boundaries, OR we HIT a queen
	nextRow := kingCoords[0]
	nextCol := kingCoords[1]
	for {
		//start at kings coords calculate the next points
		nextRow += direction[0] // ROW
		nextCol += direction[1] // COL

		//check boundaries and exist case
		if nextRow < 0 || nextRow > 7 || nextCol < 0 || nextCol > 7 {
			return false, nil //exit we are out of boundaries
		}

		//else we are within boundaries
		//is it a queen?
		if queenLocationmat[nextRow][nextCol] {
			return true, []int{nextRow, nextCol}
		}

	}

}
