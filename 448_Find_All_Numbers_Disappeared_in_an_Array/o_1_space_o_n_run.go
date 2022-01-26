package main

func findDisappearedNumbers(nums []int) []int {
	//BRUTE FORCE

	//idea -> since we know that the range is from 1 -> n, we can hash the values as a set
	//then walk from 1 -> n, checking if the next number in the sequence exists or not if it does not add it to the result set, if it does skip it

	//O(1) saopce

	//first we know that the result doesn count

	//we also know that the index themselve son the result array can tell us which values exists or not
	//we can also it erate and mark values that do not eixsts

	//1: create the result array, and initalize values as -1
	//result must be N+1

	for _, val := range nums {
		index := Abs(val) - 1
		nums[index] = -1 * Abs(nums[index])
	}
	//2: walk the nums array, and for the particular number set the index of that value to its value

	var result []int
	for i, val := range nums {
		if val > 0 {
			result = append(result, i+1)
		}
	}
	//3: iterate over the result array, from 1 to end if we see a value other than -1 remove it, if we see -1 keep it and put its index value in its place

	return result
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
