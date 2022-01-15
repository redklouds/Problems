package main

//brute force(dfs), trree top-down, all combinnations until we hit the value target
func CombinationSumIV(nums []int, target int) int {
	//dfs
	return dfsCombo(nums, target, 0)
}

//if you draw our this brute force approach you will see MULTIPLE dfsCombos(N) with the same
//N paramter called over and over, bedesign to get the best COMBO, taht's what dfs does
func dfsCombo(nums []int, target int, curVal int) int {
	numOfCombos := 0
	for i := 0; i < len(nums); i++ {
		if curVal+nums[i] == target {
			numOfCombos++ //we found that if we add the current val to this combo we get
			//a target
		} else if curVal+nums[i] < target {
			//if the current stack can still search we make another call
			numOfCombos += dfsCombo(nums, target, curVal+nums[i])
		}
	}
	return numOfCombos
}

//using memoization
func CombinationSumIV_WithCache(nums []int, target int) int {
	cache := map[int]int{}
	return dfsComboWithCache(nums, target, 0, cache)
}

func dfsComboWithCache(nums []int, target int, curVal int, cache map[int]int) int {
	if val, ok := cache[curVal]; ok {
		//exists, return the value associated with number of combos for this depth
		return val
	}
	//else it doesn exists lets go ahead and calculate it

	for i := 0; i < len(nums); i++ {
		if curVal+nums[i] == target {
			cache[curVal]++
		} else if curVal+nums[i] < target {
			//keep searching
			cache[curVal] += dfsComboWithCache(nums, target, curVal+nums[i], cache)
		}
	}
	return cache[curVal]
}

//this method implements a Dynamic Programming BOTTOM-UP-DP approach
//always ask and infer, so, how many ways can i repersne value 0, 1,2,3,4..., N
//dp[val] = how many wayts can i repersent this value
func CombinationSumIV_DP(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1 //if you write this logic on papaer you will see dp[0]=1 lets use reference ourselves
	//3-3 = 0 -> dp[0]=1 means we choose ourselves and the other coiun doesn't need to exists
	//realize that at each iteration
	//JUST LIKE COIN CHANGE -> Except coinchange was looking at dp[i] = how many minimal way to creat token
	for currentTargetSum := 1; currentTargetSum <= target; currentTargetSum++ {

		//for ever choice/COIN, we want to 'take this coin, subtract the choice we have now, get the 'othercoin'
		//and and understand that if i have this coint + the other coint i make the sum, BUT the question is
		//how many ways can i create the coin, so if i select myself, + the other coin how many ways can i create the other coin with?

		for coin_choice := range nums {
			otherCoin_choice := currentTargetSum - nums[coin_choice]
			if otherCoin_choice >= 0 {
				dp[currentTargetSum] += dp[otherCoin_choice]
			}

		}
	}
	return dp[target]
}
