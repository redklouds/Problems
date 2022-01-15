package main

import "sort"

/*
Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func ThreeSome(nums []int) [][]int {

	//couple conditions we can DERIVE, Extract
	//the three numbers CANNOT be the same index
	// the result CANNOT contain Duplicates

	//we select a index (choice)

	resultArray := [][]int{}

	//choice is a,

	//sort the array, so this works our method that is, sot aht we can increment or
	//decrement our b and c to match zero
	sort.Ints(nums)
	for a := 0; a < len(nums); a++ {
		//if duplicate to what we've processed then skip this number
		if a > 0 && nums[a] == nums[a-1] {
			//we match the previous number skip
			continue
		}

		l, r := a+1, len(nums)-1
		for l < r {
			//aagain we don't want to iterate until they touch becase then L == R is not apart of the condition
			threeSom := nums[a] + nums[l] + nums[r]
			if threeSom < 0 {
				//our sum is less than zero increse it by 1 (left) its sortted ascending
				l++
			} else if threeSom > 0 {
				//tooo 'big' decrease it by moving right down
				r--
			} else {
				//equals zero lets go ahead and add it to results
				resultArray = append(resultArray, []int{nums[a], nums[l], nums[r]})
				//now that we added this value, we also still need to SHIFT one of our pointers
				//which one? just choose one, BUT MAKE SURE that if you encounter the same value as
				//before, you kleep moving it forward
				l++
				for nums[l] == nums[l-1] && l < r {
					l++ //make sure that l is not moving out of bounds, AND that l is not the same l as before

				}
			}

		}

	}
	return resultArray
}
