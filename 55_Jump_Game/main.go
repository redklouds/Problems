package main

func canJump_BruteForce(nums []int) bool {

	//using DFS to do dis, we know that when we jump our objective is that our index equals the last
	//index

	return dfs(nums, 0)
}

func canJump_memoization(nums []int) bool {
	cache := map[int]bool{}
	return dfs_memozation(cache, nums, 0)
}

func dfs_memozation(cache map[int]bool, nums []int, i int) bool {
	if _, ok := cache[i]; ok {
		return cache[i]
	}

	if i == len(nums)-1 {
		cache[i] = true
		return cache[i]
	}
	if nums[i] == 0 {
		cache[i] = false
		return cache[i]
	}

	//for each jump we want to get ALL combinations of jumps of the current
	//maximum amounts
	for jumpAmount := 1; jumpAmount <= nums[i]; jumpAmount++ {
		//for each jump amount lets see if we can reach the end
		if dfs_memozation(cache, nums, i+jumpAmount) {
			return true
		}
	}

	return false
}

//lets use memoization
func dfs(nums []int, i int) bool {

	if i == len(nums)-1 {
		return true
	}
	if nums[i] == 0 {
		return false
	}

	//for each jump we want to get ALL combinations of jumps of the current
	//maximum amounts
	for jumpAmount := 1; jumpAmount <= nums[i]; jumpAmount++ {
		//for each jump amount lets see if we can reach the end
		if dfs(nums, i+jumpAmount) {
			return true
		}
	}

	return false
}

//uising DP - bottom -up Approach
func canJump_DP(nums []int) bool {

	if len(nums) == 2 {
		if nums[0] != 0 {
			return true
		}
		return false
	}
	//the idea again is tthe subproblem, whats our sub problem
	//what camn we infer? whats our goal?
	//our goal is to reach the end, so.. what about n -1 index, can they reach the end? is it possible for them to reach the end?
	//meaning if i land on that spot can i reach the end? our dp array will ask and answer thes questions, so that, our dp will ask, if i land on this index
	//can it reach the end?

	dp := make([]bool, len(nums))

	//make the inital state, INITAL STATE always information
	dp[len(dp)-1] = true // the last will always be able to reach the last

	//previousIdx := len(nums) - 1
	//go backwards
	for currentIdx := len(nums) - 2; currentIdx >= 0; currentIdx-- {
		//we want to ask if the current index, can reach the next one, because the nesxt index also know if IT can reach the end that will tell you, if you end
		//on this index, you will or will not be able to reach the end

		//for every eligable Jumper, if from where i stand jump to one of them, can any of them make it to the end??, because original i only checked the next door nieghbor but we know that they
		//might not be able too

		if nums[currentIdx] == 0 {
			dp[currentIdx] = false
		} else {
			for nextJump := 1; nextJump <= nums[currentIdx]; nextJump++ {
				//does there exists a neighbor who can make it
				if nextJump+currentIdx < len(nums) && dp[nextJump+currentIdx] {
					dp[currentIdx] = dp[nextJump+currentIdx]
					break;
				}
			}
		}

		/*
			if nums[currentIdx] >= previousIdx-currentIdx {
				//we can reach the next index
				dp[currentIdx] = dp[previousIdx]
			} else {
				//we cannot
				dp[currentIdx] = false
			}

			previousIdx = currentIdx

		*/

	}

	return dp[0]
}
