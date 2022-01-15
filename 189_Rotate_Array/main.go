package main

import "fmt"

func rotate(nums []int, k int) {
	//go to K, swap/flip the first and second parition,

	//K is called the inflection point
	//the point at which you will partiutoion the array into two parts

	//you want to reverse, the first partion and reverse the secopnd partioon
	//THEN you want to reverse the entire partion afterwards

	//swap the first partition
	for i := 0; i < k; i++ {
		nums[i], nums[k-i] = nums[k-i], nums[i]
	}

	fmt.Println(nums)
}
