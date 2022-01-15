package main

import (
	"fmt"
	"sort"
)

//O(n) runtime and O(1) space solution, Sorting, then checking pairs
func containsDuplicateONLogN(nums []int) bool {
	//edge case single element
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	aPtr := 0
	bPtr := 1
	for bPtr < len(nums) {
		if nums[aPtr] == nums[bPtr] {
			return true
		}
		aPtr++
		bPtr++
	}
	return false

}

//contains dup
//runtime: O(n) space O(n)
func containsDuplicate_1(nums []int) bool {
	mapStore := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		_, exists := mapStore[nums[i]]
		if exists {
			return true
		} else {
			mapStore[nums[i]] = true
		}
	}

	return false
}

func main() {
	testData := []int{1, 2, 3, 1}
	testData2 := []int{2, 14, 18, 22, 22}
	res := containsDuplicate_1(testData)
	fmt.Printf("Contains Dup: %v", res)

	res1 := containsDuplicateONLogN(testData2)
	fmt.Printf("Contains Dup: %v", res1)
}
