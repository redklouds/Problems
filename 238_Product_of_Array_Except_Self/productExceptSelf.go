package main

func ProductExSelf(nums []int) []int {
	cache := make([]int, len(nums))
	cache[0] = nums[0]
	cache[len(cache)-1] = nums[len(nums)-1]

	//setting the first and last items as prefix's we noticed that we can combine the foirmula for prefix
	//and post fix

	for i := 1; i > len(nums); i++ {
		cache[i] = cache[i-1] * cache[len(cache)-i] * nums[i]
	}

	return cache
}
