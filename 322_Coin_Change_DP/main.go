package main

import "math"

//O(n)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//init array
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	//the understanding of working from BOTTOM up
	//THINK how many number of coins can we make up
	//the DP array starting from the MINIMUM = WHAT WE KNOW. WHAT WE ALREADY HAVE lets build onto that, we know that with a target of ZERO there is ZERO number of coins that can fit this

	//so the idea is at each interval, we can check, given a target value, if i choose THIS coin which means i've picked ONE of the coins out of the bunch,

	//take its value.. lets get the current target and subtract this coin i've selected, now... if the sum is greater than 1 meaning that there is a solution somwhere, we can ask what is the difference? becuase the different will tell us what is the minium number of coins that can be made from that
	for sum := 1; sum < amount+1; sum++ {
		for _, c := range coins {
			diff := sum - c
			if diff >= 0 {
				//positive
				dp[sum] = Min(1+dp[diff], dp[sum])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1 //meaning there isn't a solution
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
