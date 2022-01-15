package main

import "fmt"

func setZeroes_Best(matrix [][]int) {

	//i need some way to set the rows and columns to zero
	//i cannot use a DFS on four directions since that will traverse the other directions and not jsut the row
	//what i can do, is track, since if i find a zero at a cell its ENTIRE row and Colum will be zero,
	//therefore basically finding a zero will mean that i mark that row and colum as being zero

	//i can use two arrayts one for row and one for columns to track.. however ot make it contant space..

	//can't i also use one f the rows and columns from the grid itself?

	//we will need to store information as a variable O(1) to indicatre if the row is zero,

	//then at each iteration , starting at the top left corner, we check if there exists a zero, if there does, we will

	//we need an extra space to store information about the COLUMN ZERO
	//since it will overlap for the ROW ZERO we don't want that, since the row of 0 might not need to zero out

	//the other assumption about 'keeping informatnion of row and coluimn zero' we need to realize that ANYTIME a cell in THAT ROW or COL is ZERO , it doesn't matter what's in index zero of the first row or column that ENTIRE ROW and COL will be zero, so we can overrite what is already there

	rowZeroIdx := false //does row zero have a ZERO
	//the rowZeroIdx variable tells us IF the zerro row needs to be zero out
	//the first ROW will tell us which COLUMNS need to be zerod out

	for row := range matrix {
		for col := range matrix[0] {

			//if we find a cell that is zero we mark it

			if matrix[row][col] == 0 {
				//found a zero mark its row first
				matrix[0][col] = 0 //means that the entire column needs to be zero out
				if row == 0 {
					rowZeroIdx = true
				} else {
					matrix[row][0] = 0
				}

			}
		}
	}

	//then we will update the coluis and rows

	//change the rows
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	//for the first row
	if matrix[0][0] == 0 {
		for row := range matrix {
			matrix[row][0] = 0
		}
	}

	if rowZeroIdx {
		for col := range matrix[0] {
			matrix[0][col] = 0
		}
	}

}

func SetZeros(matrix [][]int) {
	zero_col := make([]bool, len(matrix[0]))
	zero_row := make([]bool, len(matrix))

	//find the zeros
	for row := 0; row < len(matrix); row++ {

		for col := 0; col < len(matrix[0]); col++ {

			if matrix[row][col] == 0 {
				//set the row and col
				zero_col[col] = true
				zero_row[row] = true
			}
		}
	}

	//now modify the matrix
	for row := 0; row < len(matrix); row++ {

		for col := 0; col < len(matrix[0]); col++ {

			if zero_row[row] || zero_col[col] {
				matrix[row][col] = 0
			}
		}
	}

}

func PrintMatrix(mat [][]int) {
	for _, row := range mat {
		fmt.Printf("%v\n", row)
	}
}

type pQ []int

func (q *pQ) someting() {
	old := *q //dereference the pointer and get the actual array

	n := len(old) //getting the actualy iunderlying array len

	//item := old[n-1] //getting the last item in the array

	*q = old[0 : n-1] //removing the last item and setting as the current array

	fmt.Printf("\n%v", *q)
}

func (q *pQ) Push(val interface{}) {
	//n := len(*q)      //gettting the actualy underlying array length
	item := val.(int) //cast the input object as its actual object
	*q = append(*q, item)
}
func main() {

	testData := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	PrintMatrix(testData)
	fmt.Println("")
	SetZeros(testData)
	PrintMatrix(testData)

	arr1 := make([]string, 5)

	fmt.Printf("%v", &arr1)

	a := pQ{}

	a.someting()
}
