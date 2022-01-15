package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if !CheckDimensions(len(mat), len(mat[0]), r, c) {
		return [][]int{}
	}
	resultSet := make([][]int, r)
	for i := range resultSet {
		resultSet[i] = make([]int, c)
	}

	mat_r := 0
	mat_c := 0
	for row := range mat {

		for col := range mat[0] {

			//basically keeping track of where we are
			//incrementing ever time we take a value for the current spot in result set
			resultSet[mat_r][mat_c] = mat[row][col]
			mat_c++
			if mat_c >= c {
				mat_r++
				mat_c = 0
			}
		}
	}
	return resultSet
}

func CheckDimensions(matR, matC, r, c int) bool {
	numElement := matR * matC
	otherNumElement := r * c
	return numElement == otherNumElement
}
