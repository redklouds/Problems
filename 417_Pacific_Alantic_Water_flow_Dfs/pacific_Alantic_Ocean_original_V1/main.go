package main

import (
	"fmt"
	"math"
)

func findWaterPaths(matrix [][]int) [][]int {
	//given a matrix that shows the continent, return coordinates that show when it rains that the water will run to both
	//the Paciaic ocean and alatanic ocean
	/*
					PACIFIC (X)
					__________
		PACIFIC(Y)  |		  |ALANTIC (Y)
					|		  |
					-----------
					ALANTIC(X)
	*/
	//the idea is that, we DFS from the edges(water) up into the continet
	//there exists two rules:
	//RUle !: water can travel N,S,E,W, of the current cell
	//RULES 2: water can travel IFF the direction cell is less than or equal to selected cell

	//we know that if its mountains that if you start from a cell and walk into deeper into the continent , that you will eventually
	//hit a peak of the cells, where current Cell > right Cell (walking from left to right), if we THINK IN REVERSE, water can travel DOWN from PEAK to ocean
	//if we hit a peak that peak is the higest point where water is able to travel from that peak down into the ocean,
	//that's also to say that we want to basically map out coordinates that water can travel from ocean into contient, thinking REVERSE, from Ocean to Continent
	//our formula follows -> if the Current Cell is Greater than or equal to the previous cell THAN we can also say that water CAN FLOW from the current cell to the ocean

	//create two matrixes, one to store the paths water can travel too from the pacific and one where the water can travel to alantic ocean

	//so water from alatic to Continent

	//water from Alantic to Continent

	direction := [][]int{
		{0, -1}, //left
		{0, 1},  //right
		{-1, 0}, //UP
		{1, 0},  //DOWN
	}
	pacific := make([][]bool, len(matrix))

	alantic := make([][]bool, len(matrix))

	initalizeBoolMatrix(pacific, len(matrix[0]))

	initalizeBoolMatrix(alantic, len(matrix[0]))

	//perform Modified DFS

	//REMEMEBR THAT the TOP (X) First row and the FIRST Column belong to the first ocean (Pacific) "Touchging the pacific"

	//REMEMEBR THAT THE BOTTTOM(X) LAST ROW and the LAST COLUMN (Y) Belong to the second ocean (Alantic) "Touching the alantic"

	//so we start our respectice search this way

	//set up and search the columns

	for col := 0; col < len(matrix[0]); col++ {
		//run dfsModified to search columns for respective oceans
		//we want to start at the first COLUMN, we use MININFITIY to denote the lowest part(ocean touchiung )
		dfsModified(matrix, 0, col, math.MinInt32, pacific, direction)

		//dfs for the alantic which COLUM IS THE LAST COLUMN
		dfsModified(matrix, len(matrix)-1, 0, math.MinInt32, alantic, direction)
	}

	//now we search along the X axis, the top and bottom rows

	for row := 0; row < len(matrix); row++ {
		//the pacific, 'controls' touches, the TOP ROW
		dfsModified(matrix, row, 0, math.MinInt32, pacific, direction)
		//the alantic , 'touches the BOTTOM' row, START at the bottom
		dfsModified(matrix, row, len(matrix[0])-1, math.MinInt32, alantic, direction)
	}

	//result array with the coordinates

	resultCoords := make([][]int, 0)

	//find the coordinates that both pacific and alantic share those are the peaks and places where water will both travel to both oceans

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if pacific[row][col] && alantic[row][col] {
				resultCoords = append(resultCoords, []int{row, col})
			}
		}
	}

	return resultCoords

}

//our rules are if the current item is larger the previous value, then we should continue searching the four corners
func dfsModified(matrix [][]int, row int, col int, prev int, ocean [][]bool, direction [][]int) {
	//BASE CASE: if the current row and col indexs are our of bounds then we should skip this iteration because, we have hit water or out of the matrix bounds
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		//we are out of boundary on search lets finish
		return
	}
	//else if the current cell has been checked, OR the current value is SMALLER than the previous value, return
	//since we want to keep climbing the continent it doesn't do us any good to find smaller values within the contineut
	if matrix[row][col] < prev || ocean[row][col] {
		return
	}

	//make this as TRUE since it IS GREATER THAN
	//now we know we are NOT out of bounds, and that our current value IS GREATER Or Equal to the previous value, lets search for that same condition on its
	ocean[row][col] = true

	//adjacent (N,S,E,W) cells (UP< DOWN< LEFT<RIGHT)
	for _, dir := range direction {
		//dir[0] ROW
		//dir[1] COL
		//search each adjacent direction in using DFS method
		//the direction index are made such that you will 'ADD' to that curent index up down left or right
		dfsModified(matrix, row+dir[0], col+dir[1], matrix[row][col], ocean, direction)
	}
}

func initalizeBoolMatrix(mat [][]bool, MLen int) {
	//matrix are pass by reference
	for i := range mat {
		//initalize i is index
		mat[i] = make([]bool, MLen)
	}
}

func main() {

	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	//mRows := len(matrix)
	//mCol := len(matrix[0])
	//mRows := 5
	//mCol := 5
	//pacific := [][]bool{}
	//2D arrays cannot be initalized with a variable they MUST BE CONSTANT
	//instead use zero-based idea

	pacific := make([][]bool, len(matrix))
	for i := range pacific {
		//i is index
		pacific[i] = make([]bool, 5)
	}
	fmt.Printf("%v", pacific)

	res2 := findWaterPaths(matrix)

	fmt.Printf("\n%v", res2)

	h := make([][]int, 5)
	h = nil

	fmt.Printf("%v", h)
}
