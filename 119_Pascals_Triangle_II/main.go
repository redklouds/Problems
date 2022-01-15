package main

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}
	/*
		if rowIndex == 1 {
			return []int{1, 1}
		}
	*/
	//if we are oob then it counts as a 0

	//our solution matrix
	dpMat := make([][]int, rowIndex+1)
	//initalize the col

	for i := range dpMat {
		dpMat[i] = make([]int, rowIndex+1)
	}
	//inital state is the first (0,0) is 1
	dpMat[0][0] = 1
	for r := 1; r <= rowIndex; r++ {

		for c := 0; c <= r; c++ {
			//LEFT ROW -1, col -1
			var leftVal int //either zero or oob
			if c+-1 < 0 {
				leftVal = 0
			} else {
				leftVal = dpMat[r-1][c+-1]
			}
			dpMat[r][c] = leftVal + dpMat[r-1][c]
		}
	}

	return dpMat[rowIndex]
}
