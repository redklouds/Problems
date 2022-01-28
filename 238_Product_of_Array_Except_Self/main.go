package main

//O(n) run  time and O(n) space time
func productExceptSelf(nums []int) []int {
	//idea have a prefix array and a post fix array

	//with the prefix and post fix values stored within them

	//we want to essentially create two stores, one where we compute the prefix's at index i and one whre we compute the postfix at index i

	//given these values we can multiple at selected I, what is the prefix and post fix and combine those results to place at I

	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	//our array will store the prefixs given the nums array, now lets combine/computte the postfixes

	postfix := 1
	for i := len(nums) - 1; i > -1; i-- {
		result[i] *= postfix
		postfix *= nums[i] //multiple by whats already in nums
	}

	return result
}

//https://dev.to/akhilpokle/product-of-array-except-self-a-mind-boggling-google-interview-question-1en1
//its all about generating the left most products
//breaking this problem into a smaller problem, we know we need to get the products
//of the left of i so 0 to i -1 and right i +1 to n
//so the res array first pass for left, stores the continuing product
//of every preceeding product
func ProductExceptSelf_v1(nums []int) []int {
	res := make([]int, len(nums))
	left := 1
	right := 1

	for i := 0; i < len(nums); i++ {
		res[i] = left
		left = left * nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = right * res[i]
		right = right * nums[i]
	}
	return res
}

func ProductExceptSelfLeft(nums []int) []int {

	left := make([]int, len(nums))
	mul := 1

	for i := 0; i < len(nums); i++ {
		left[i] = mul
		mul = mul * nums[i]
	}

	return left
}
