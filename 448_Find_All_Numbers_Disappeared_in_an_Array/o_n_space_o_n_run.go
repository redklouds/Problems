package main

//O(n) space , just have a hash for the values,
func findDisappearedNumbers_v2(nums []int) []int {

	//1: run through nums collect the numbers
	numCollection := make(map[int]bool)
	for _, val := range nums {
		numCollection[val] = true
	}

	//2 : run through the 1 -> n asking if the hash contains the particular number
	var result []int
	for i := 1; i < len(nums); i++ {
		if _, ok := numCollection[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}
