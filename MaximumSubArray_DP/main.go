package main

import (
	"fmt"
	"math"
)

func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	//we can optimzie since we know what the current max is we can
	//update that max whenever we choose a new subarray to start

	global_max := -math.MaxInt32
	local_max := nums[0]
	for i := 1; i < len(nums); i++ {
		local_max = Max(local_max, local_max+nums[i])
		if local_max > global_max {
			global_max = local_max
		}
	}

	return global_max
}

//using memoization its kinda like the product of self but only doing one side
//you are  to calculate all the sums of the left most subarray
func MaxSubArray(nums []int) int {
	localMax := make([]int, len(nums))
	localMax[0] = nums[0]
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		localMax[i] = Max(nums[i], nums[i]+localMax[i-1])
	}
	fmt.Println(localMax)

	global_max := -math.MaxInt32
	for _, v := range localMax {
		if v > global_max {
			global_max = v
		}
	}
	return global_max

}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//Brute Force to finding the maximum sub array is to simply go through every element,
//and build the sub arrays keeping a count of the maxmium amount, building on top of that
func BruteForceMaxSubArray(nums []int) int {
	max := -math.MaxInt32
	var val int
	for i := 0; i < len(nums); i++ {
		//check the inital value
		val = nums[i]
		if val > max {
			max = val
		}
		for j := i + 1; j < len(nums); j++ {
			val += nums[j]
			if val > max {
				max = val
			}
		}
	}

	return max
}

func main() {
	max := -math.MaxInt32
	fmt.Println(max)

	testData := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	res := BruteForceMaxSubArray(testData)

	fmt.Printf("Maximum Subarray %d\n", res)

	res1 := MaxSubArray(testData)

	fmt.Printf("Maximum Subarray %d\n", res1)

	testData3 := []int{-1, -2}
	res2 := maxSubArray2(testData3)

	fmt.Printf("Maximum Subarray %d\n", res2)
}
