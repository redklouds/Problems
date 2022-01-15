package main

//fatest you can do
func climbStairsDp(nums []int, targetVal int) int {

	dp := make([]int, targetVal+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {

		for _, steps := range nums {
			if i-steps >= 0 {
				dp[i] += dp[i-steps]
			}
		}
	}
	return dp[targetVal]
}
