package main

//26. Remove Duplicates from Sorted Array

/*
We can use two pointer method :
one idx pointer pointing to the latest unique element,
 another pointer to scan through the array. If the second pointer
 encountered values different than what idx is pointing, increase
  idx and add the new unique value. Finally return the idx + 1 as
  the length of the unique array.
*/
func removeDuplicates(nums []int) int {

	baseIdx := 0
	for i := range nums {
		if nums[i] != nums[baseIdx] {
			baseIdx++
			nums[baseIdx] = nums[i]

		}
	}
	return baseIdx + 1
}
