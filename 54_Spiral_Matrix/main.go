package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	//so the idea is that we search right, then down, then left, then up

	//we can cycle this directions loop, our two edge cases, if
	//if we hit OOB, or if we hit a cell that we visited

	//to exist we can have a flag called DONE

	//if by the end of a cycle IE, (UP) direction and DONE is still true, means we exist because at each interval , if we hit a edge, or visited, DONE is set, else Done is flipped back

	//for each directions

	//FOR VISITED we also know that the max a cell will be is 100, therefore we can make a visited cell 101
	DONE := false

	//this ONLY works in the set order of Right, Down, Left, Up
	directions := [][]int{
		//ROW = UP/DOWN, COL = LEFT/RIGHT"
		{0, 1},  //RIGHT
		{1, 0},  //DOWN
		{0, -1}, //LEFT
		{-1, 0}, //UP
	}

	var results []int
	row := 0
	col := 0
	results = append(results, matrix[0][0])

	matrix[0][0] = 101
	di := 1
	for !DONE {
		for _, dir := range directions {
			//Right first

			testRow := row + dir[0]
			testCol := col + dir[1]
			if testRow < 0 || testRow > len(matrix)-1 || testCol < 0 || testCol > len(matrix[0])-1 || matrix[testRow][testCol] == 101 {
				//DONE = true
				//break
				di += 1
				if di == 4 {
					DONE = true
					break
				}
				continue
			}
			//we want to keep going this direction until 2 cases, OOB or Visited
			for {

				row = row + dir[0]
				col = col + dir[1]
				if row < 0 || row > len(matrix)-1 || col < 0 || col > len(matrix[0])-1 || matrix[row][col] == 101 {
					//we have a edge or visited cell, change directions
					//lets result our row and col back to the state it was
					row = row - dir[0]
					col = col - dir[1]

					break
				}
				//we are inboundaries, lets add this value to the results
				results = append(results, matrix[row][col])
				matrix[row][col] = 101

			}
			di += 1
			//have to check at each end of direction if the next direction has any elements left

			//if the direction equals
		}
		di = 1
		//GIVEN TEH KNOWLEGE that all PRECEDING ROW and post(top and bottom rows have been proceeded)
		//we can say that as long as the next RIGHT cell has NOT been visited or OOB then we are finished

	}

	return results
}
