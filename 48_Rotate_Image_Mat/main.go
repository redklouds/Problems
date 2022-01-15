package main

func rotate(matrix [][]int) {

	//idea is to set up four corners, you NEED TO DO THIS SO THINGS ARE MORE CLEAR AND STRAIGHT FORWARD

	//you'll notice the matrix is nXn meaning that its a square, moving the elements will always move them to four corners

	//therefore , we will have four variables tracking the four corners

	//at each step we will shift the values to their corners

	//ALSO NOTICE when we work CLOCKWISE, we will NEED TWO variables to store the curernt cell value and the next cell valuie, which makes the code a little harder to follow

	//to fix this think backwards, if we do this we only need to store a single variable for the top left place,

	left, right := 0, len(matrix)-1

	//while the left is less than the right boundaries we continue
	for left < right {
		for i := 0; i < (right - left); i++ {

			//set the top and bottom boundaries
			top, bottom := left, right

			//four things we need to do
			//store the TOP LEFT value
			topLeft := matrix[top][left+i]

			//move the bottom left value into the top Left spot
			matrix[top][left+i] = matrix[bottom-i][left]

			//move the bottom right value into the bottom left spot
			matrix[bottom-i][left] = matrix[bottom][right-i]

			//move the top right value into the bottom right spot
			matrix[bottom][right-i] = matrix[top+i][right]

			//move the top left value into the top right spot
			matrix[top+i][right] = topLeft

			//finish the shift now we should update our boudaries to be inner

		}
		left += 1
		right -= 1

	}
}
