package main

type Coords struct {
	Row int
	Col int
}

//static Directions object
var Directions = [][]int{
	{0, -1}, //LEFT
	{-1, 0}, //UP
	{0, 1},  //RIGHT
	{1, 0},  //DOWN
}

//in byte
//'#' = 13
//'*' = ??
func getFood(grid [][]byte) int {

	//0 create the queue
	var q []*Coords

	var startingCoords *Coords
	//1
	//first we want to set the state of the grid
	//when we first see that the cell is a '*' we want to save the starting coords for later
	//when we see that a space is 'O' open, we want to change this to a infiniy/255
	//when we see a '#' a food cell change it to ZERO and enqueue it

	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == '#' {
				//a food cell, may be multiple
				q = append(q, &Coords{r, c})
				grid[r][c] = 0 // MAX for a byte is 255, make sure you knw what the symbol values are
			} else if grid[r][c] == 'O' {
				grid[r][c] = 255 //infinity, do not need to queue
			} else if grid[r][c] == '*' {
				//the start we want to record the starting position and make it inifinity
				startingCoords = &Coords{r, c}
				grid[r][c] = 255
			}
		}
	}

	//2
	//run BFS on the grid, checking weather the current adjacent cells
	//are larger than the current cell, if it is update the adjcent cells
	//and enqueue them
	//however if the cell is a star/start we want to skip enqueueing it

	for len(q) > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			coord := q[i]
			for _, dir := range Directions {
				adjRow := coord.Row + dir[0]
				adjCol := coord.Col + dir[1]

				//if not oob, and not a 'X' lets update its distance
				if adjRow > -1 && adjRow < len(grid) &&
					adjCol > -1 && adjCol < len(grid[0]) &&
					grid[adjRow][adjCol] != 'X' && grid[adjRow][adjCol] > grid[coord.Row][coord.Col]+1 {

					grid[adjRow][adjCol] = grid[coord.Row][coord.Col] + 1
					if adjRow != startingCoords.Row || adjCol != startingCoords.Col {
						q = append(q, &Coords{adjRow, adjCol})
					}

				}

			}

		}
		q = q[qLen:]
	}
	//3
	//if the start coords are still infinity it means we were not able to reach them
	//therefore returning -1, else return whatever is in the start coor cell

	if grid[startingCoords.Row][startingCoords.Col] == 255 {
		return -1
	}
	return int(grid[startingCoords.Row][startingCoords.Col])
}
