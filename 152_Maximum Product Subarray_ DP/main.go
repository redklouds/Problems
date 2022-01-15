package main

import (
	"fmt"
	"math"
)

func MaxProductLinearSpace(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	min_dp := nums[0]
	max_dp := nums[0]
	local_min := nums[0]
	local_max := nums[0]
	gmax := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			min_dp = Min(nums[i], nums[i]*local_min) //local_min/max is the PERVIOUS value in the sequence
			max_dp = Max(nums[i], nums[i]*local_max)
		} else {
			//we encounterd a negative , so inorder to get the MAX product
			//we need could get it by multiplying the current number, with the MINIMUM proudct
			//because if the min was negative and the current num is negative we get a positiove as the max
			//product

			min_dp = Min(nums[i], nums[i]*local_max)
			max_dp = Max(nums[i], nums[i]*local_min)
		}
		local_min = min_dp
		local_max = max_dp

		if local_max > gmax {
			gmax = max_dp
		}

	}
	return gmax
}

func MaxProductDPMaxArrays(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	min_dp := make([]int, len(nums))
	max_dp := make([]int, len(nums))

	min_dp[0] = nums[0]
	max_dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			min_dp[i] = Min(nums[i], nums[i]*min_dp[i-1])
			max_dp[i] = Max(nums[i], nums[i]*max_dp[i-1])
		} else {
			//we encounterd a negative , so inorder to get the MAX product
			//we need could get it by multiplying the current number, with the MINIMUM proudct
			//because if the min was negative and the current num is negative we get a positiove as the max
			//product

			min_dp[i] = Min(nums[i], nums[i]*max_dp[i-1])
			max_dp[i] = Max(nums[i], nums[i]*min_dp[i-1])
		}

	}

	fmt.Println("Max array ", max_dp)
	gmax := -math.MaxInt32
	for _, v := range max_dp {
		if v > gmax {
			gmax = v
		}
	}
	return gmax
}

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	//create the array
	res := make([]int, len(nums))

	//since we want to calculate the last seen/current largest sub array and each index IS a subarray
	//when we start we want to assign the largest as the first index of our result index
	res[0] = nums[0]

	//starting at 1 so we can 'look back' at our history store and compare
	for i := 1; i < len(nums); i++ {
		//is the current nums[i] Larger THAN our Current Subarray put together?
		//if it is then we want to start a new sub array
		res[i] = Max(nums[i], res[i-1]*nums[i])
	}
	//now we have our array of maxProduct Results at each index we can find the largest one

	global_maxProduct := -math.MaxInt32

	for _, v := range res {
		if v > global_maxProduct {
			global_maxProduct = v
		}
	}

	return global_maxProduct

}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}




func MaxProductSubArrayV3(nums []int) int {
	//EDGE case what if we encounter a ZERO? X * Zero is always ZERo
	if len(nums) < 2 {
		return 1
	}
	if len(nums) == 0 {
		return 0
	}
	current_min := nums[0]
	current_max := nums[0]
	max_product_subarray := nums[0]

	for i := 1; i < len(nums); i++ {
		//EDGE case
		if nums[i] > 0 {
			//if its a positive we can just multiply the number as if we are including this value
			//in our subarray total (product)
			current_min = Min(nums[i], current_min*nums[i]) //track what's our lowest number
			current_max = Max(nums[i], current_max*nums[i]) //track whats our current MAx (THIS ONE WE WANT)

		} else {

			//but we cannot it needs to be a seperate value
			cm_x := current_max
			current_max = Max(nums[i], current_min*nums[i]) //ex: cm=-1 and num[i] = -5 = 5, that's a max, if you took -5 and multipled by positive you go futther to negative
			//and our min can de dericived from our mAX  * current value which is negative, to FURTHER drive a MIN value
			current_min = Min(nums[i], cm_x*nums[i])
		}

		//after al the compute we check if our current max is greater than our overall total subarray max product

		max_product_subarray = Max(max_product_subarray, current_max)
	}

	return max_product_subarray

}
