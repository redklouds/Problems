package main

//given a array that rep, heights of two bars, return the maxmimum
//amount of WATER, IE Area that you can contain

//idea is simple, Area = Length X Width, you know the index are width
//2 points left and right, get diff, then find the smaller of two heights and
//compute, moving the smaller heigh in search of a larger area
func MaxRaimWater_Area(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	left, right := 0, len(heights)-1
	maxArea := 0

	for left < right {
		//here I DO NOT want to compute when Left and Right are on 'TOP' of one another
		//it wont give me anything, we just want the boundary
		//compute the width

		width := right - left

		//now find the lower height of the two, since if you think of a box
		/*
			|~~~~~~~~~
			|	|
			|	|
		*/
		//water 'spills' over the lower side, so if you try to compute max area of the higher bar
		//then you will inaccuratly coimpute the wrong area since water spills over

		if heights[left] < heights[right] {
			//left is the smaller height of the two
			maxArea = Max(maxArea, width*heights[left])
			//move the left (smaller forward in search of a higher heigh since right is already highj)
			left++
		} else {
			//the heights are the same or the right is greater
			maxArea = Max(maxArea, width*heights[right])
			right--
		}
	}
	return maxArea
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
