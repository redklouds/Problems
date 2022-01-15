package main

func rob(nums []int) int {
	dp := make([]int, len(nums)) //repersents, if i 'hitup' this house, what's the MAXIMUM amount of money i can
	//get if i 'hit up this particular house', since we have to see which is the max house we can go with

	//set up the default values in dp as self, its always comparing
	//SELF OR self + house im aboutta hit up, which is larger
	//also setting to self solves the oob at the end
	for house := range dp {
		dp[house] = nums[house]
	}
	for house := len(nums) - 1; house >= 0; house-- {

		for otherHouseMax := house + 2; otherHouseMax < len(nums); otherHouseMax++ {
			//if we out of bounds this will not call and the default house price will stay the default we set earlier! yaya

			//how calculate our dp of HOW MUCH MONEY can i rob, if i take this current house DP meaning that if i jump to this index,
			//that means that this dp[index/house] i have already ALSO calculated before how much THIS index can also produce if i land here
			dp[house] = Max(dp[house], nums[house]+dp[otherHouseMax])
		}
	}

	curMax := 0
	for _, MaxHitVal := range dp {
		curMax = Max(MaxHitVal, curMax)
	}

	return curMax
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
