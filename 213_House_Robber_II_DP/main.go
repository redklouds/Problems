package main

func rob(nums []int) int {

	rob1 := robHelper(nums[1:])
	rob2 := robHelper(nums[:len(nums)-1])

	curMax := Max(rob1, rob2)
	return Max(curMax, nums[0])
}

//because the first and last house cannot be robbed together, we need to ommit them from each rob scenario so TWO dp arrays one for robbing the 1 -> n-1 houes
//and one to rob the 0 -> n -2

func robHelper(nums []int) int {

	//only added rule if im the first house i just cannot touch the last house
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = nums[i]
	}

	//Just like House Robber I

	for house := len(nums) - 1; house >= 0; house-- {

		//we need to calculate for each house we land on what's their MAX profit landing here with ONE rule
		for otherHouse := house + 2; otherHouse < len(nums); otherHouse++ {

			//i cannot rob the first house and last house trogether, so think... THINK
			//we go backwards, finind the max for each jump, i can run the rob function for the entire array from 2 partitions
			//o -> len(nums)-2

			dp[house] = Max(dp[house], nums[house]+dp[otherHouse])
			//acception
			/*if house != 0 && otherHouse != len(nums)-1 {
				//then we gucci to check because if we are on the current first house and want to find out max to the last house nah

				dp[house] = Max(dp[house], nums[house]+dp[otherHouse])
			}
			*/
		}
	}

	maxProfitRobbing := 0
	for i := range dp {
		maxProfitRobbing = Max(maxProfitRobbing, dp[i])
	}

	return maxProfitRobbing
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
