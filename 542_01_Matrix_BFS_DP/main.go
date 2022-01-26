package main

import "math"

type Coords struct {
	Row int
	Col int
}

var DIRECTIONS = [][]int{
	//row, col,
	{0, -1}, //LEFT
	{-1, 0}, //UP
	{0, 1},  //RIGHT
	{1, 0},  //DOWN
}

func updateMatrix(mat [][]int) [][]int {

	//breaking the problem statement down

	//what we know

	//subproblems,
	//we KNOW we can perform BFS on each cell to find its distance  from each zero
	//to the nearest 1 and place that value in place

	//first thing to do, set the inital state, since we know that our goal for each step

	//is to ask, if i add 1 to myself(current cell) are you(adjacent cell) still greater than myself?
	//if so, this means that we can make a shorter path/distance, since our inital state will

	//set the grid to all 1's turn into infinity, and all zeros stay the same, a zero next to another zero,
	//will be the same, since if you are zero distance away from a zero , you cannot possibly be any closer

	//set the inital state of the board to 1's become infinity, and if we encounter a zero

	//we want to enqueu that into the bfs queue to process its paths

	//so step 1: set the inital state of the map such that we say all 1's are infinitly furthest away from all zeros
	//if we encounter a zero we enqueue it so that we can start our bfs from those zeros

	var q []*Coords

	for row := range mat {

		for col := range mat[0] {

			//inital state,
			if mat[row][col] == 1 {
				mat[row][col] = math.MaxInt32
			} else {
				//its a zero enqueue
				q = append(q, &Coords{row, col})
			}
		}
	}

	//Step2: we need to now take these zero coords and bfs ssearch to update their
	//paths to their respecitive distances

	for len(q) != 0 {
		qLen := len(q)

		for i := 0; i < qLen; i++ {
			//we now search
			coord := q[i]
			for _, dir := range DIRECTIONS {
				adjRow := coord.Row + dir[0]
				adjCol := coord.Col + dir[1]
				if adjRow > -1 && adjRow < len(mat) && adjCol > -1 && adjCol < len(mat[0]) && mat[adjRow][adjCol] > mat[coord.Row][coord.Col]+1 {
					//we know we can do better, since the
					mat[adjRow][adjCol] = mat[coord.Row][coord.Col] + 1
					q = append(q, &Coords{adjRow, adjCol})
				}
			}
		}

		//lesser slices needed, optmizations not at each interval
		//remove elements
		q = q[qLen:] //[inclusiveIdx:]
	}
	return mat
	//how can we break this down into parts
}
