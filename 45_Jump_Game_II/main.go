package main

import "math"

func jump(nums []int) int {
	//o(n) solution, in which you must think in terms of throwing a rock,
	//you throw a rock as far as the current index allows, it, and see if you can reach the end,
	//you know that the first partition, if you cannot find a farther reach will be the first partion, and in this case
	//the farthestJumpEndIdx will be reached therefore meaning that in the first partion you couln't find a farther jump to the ned
	//so now you will need to pick up the rock and throw it again, doing the same process
	farthestJumpIdx := 0
	partitionJumpEndIdx := 0
	numJumps := 0
	for i := 0; i < len(nums)-1; i++ {

		farthestJumpIdx = Max(farthestJumpIdx, i+nums[i])
		if i == partitionJumpEndIdx {
			numJumps++
			partitionJumpEndIdx = farthestJumpIdx
		}
	}
	return numJumps
}

//this takes extra space
func jump_DP(nums []int) int {

	if len(nums) == 2 {
		return 1
	}
	//dp array will tell us
	//"At index I, what's the MIMIMAL amount of hops/jumps will it take me to get to the end"
	//we assume that there is always a path,

	//the process go at each I we KNOW that the end can be reached wtih ZERO jumps, that's because
	//we KNOW that if i land at the end I made it  NO jumps,

	//for all other values from the end up, we can compute the number of jumps because if you do the DFS
	//tree solution you will see even with memoization that ALLLLL the combinations computations starts heavy at the front
	//top and then tickles downwards, by then you optimiza a little, what if we started from the BACK we know our end case
	//our success case and our default state which is dp[end] = 0, what about dp[end-1]=? dp[end-2]=?
	dp := make([]int, len(nums))

	dp[len(dp)-1] = 0 //default case at the end there exists no need to jump

	//lets set our default states for all the DP values, (THIS IS ALWAYS IMPORTANT)

	for i := 0; i < len(dp)-1; i++ {
		dp[i] = math.MaxInt32 //since the boundary case is that values max is going to be 1000 so we can safely  say a max would be 1001
	}
	//for DP you need couple things:

	//1: a Question, What are you trying to ask? given the prompt, and what will the dp answer?
	//2: the BASE CASE, usually at the very beginning of the problem or at the very end. "what does a success case look like? (plain sight)"
	//3: DEFAULT STATE, either dp[end] = 0 or dp[start] = something to start then dp ARRAY default values. amd thjat's it

	//4; a rythm to stop making more compute if you've already computed the last answer

	//for current index we want to know how many minimual jumps it takes to get to the end
	for currentIndex := len(nums) - 1; currentIndex >= 0; currentIndex-- {

		//there won't be zero and always a path so we do not need to worry about his case
		for neighbor := 1; neighbor <= nums[currentIndex]; neighbor++ {
			//we iterate the amount the current value can jump
			//watch for out of bounds
			if currentIndex+neighbor < len(nums) {
				dp[currentIndex] = Min(dp[currentIndex], 1+dp[currentIndex+neighbor])
			}
		}
	}

	return dp[0]
}

//THINK IF THERE IS ALWAYS A PATH, WE can simple SEND IT and make sure we go the furthest everytime
//everytime we make it farther, its like throwing a rock if we throw a rock and then eventually walk to it,
//we can say that that is s ajump
func linearJump(nums []int) int {
	currentJumpEnd := 0
	farthestJump := 0
	numJumps := 0
	for i := 0; i < len(nums); i++ {
		farthestJump = Max(farthestJump, i+nums[i])
		if i == currentJumpEnd {
			numJumps++
			currentJumpEnd = farthestJump
		}
	}
	return numJumps
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
