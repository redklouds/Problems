package main

import (
	"fmt"
)

//given 7 keep subtracking
//keep track of the recursive steps, because everytime i recursively go deeper
//there exists a POSSIBLE solution,
//if there exists a solution ie, targetVal==0 then return the Num Stack back
//get the minmum stack value returned
func coinChange(coins []int, targetVal int) int {
	if targetVal == 0 {
		fmt.Println("Found a solution")
		return 0

	}
	if targetVal < 0 {
		return 9999
	}
	//i know im here, so 1 stack or 1 type of coin can be added here

	curDep := 1
	for i := 0; i < len(coins); i++ {
		//for each coin
		curDep = 1 + coinChange(coins, targetVal-coins[i])

	}

	return curDep

}

func CountStacks(n int) int {
	curDepth := 1
	if n != 1 {
		n-- //10 9 8 7 6 5 4 3 2 1 0
		curDepth += CountStacks(n)
	}
	return curDepth
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

//given a array of coins, and a target sum you want
//generate a map that contains the MINIMUM number of coins
//you would need to make the respective sum
//ie map[sum] = min # of coins(from our array)
func MinmumNumberOfCountMap(coins []int, targetSum int) int {
	minCoinCache := map[int]int{}
	//set the first value, however go is zero based, so even if cache[0] is not explicitly insterted,
	//if access to cache[0] does arise it would just return 0 //minCoinCache[0] = 0
	for sum := 0; sum <= targetSum; sum++ {
		//we will calculate BOTTOM UP, starting from how minimum coins it takes to
		//make a sum of 0, 1, 2,.... to Target, this way, we are able to tell so.. how many real coins does it take
		//to get to targetSum

		//for each targetSum Value
		//currentSumMin := math.MaxInt32
		for _, coin := range coins {

			if sum-coin == 0 {
				minCoinCache[sum] = 1
			} else if sum-coin > 0 {
				//there's a remainder
				minCoinCache[sum] = 1 + minCoinCache[sum-coin] // get the num of remainder
			}
			/*
				var localMinSum int
				//for each coin
				//calculate the difference between the target Sum and the current coin

				diff := sum - coin
				//if the difference is ZERO, ie targetSum is 5 and current coin is 5, this means
				//that there is minimum of ONE coin we need to make a sum of 5, perfectly
				if diff == 0 {
					//we have our winning pair, the current coin perfectly fits the current target sum
					localMinSum = 1 //we have found a pair with just 1 coin
					break
				} else if diff < 0 {
					//ie there isnt such pairing exists targetSum is 0 and we have [1,2,4] , there exists no such pairing
					localMinSum = 0

				} else {
					//this is interesting, we have found a postive remainder, this is where we want to ask
					//"WHATS THE MINIMUM # OF COINS TO MAKE UP THE REMAINDER", since we've used 1 of the current coin already
					//we want to 1 + store[remainder], the store will always hold the minimum number of coins for the respective target Sum
					if val, ok := minCoinCache[diff]; ok {
						//if this exists in the store, we can retrieve it here
						localMinSum = 1 + val //val is the min number of the other coin such like if targetSum is 2 - Choosen coin (1)
						//with remainder of 1, our remainder store will have cache[1]=1 there is minimum of 1 coin to get to sum of 1
					} else {
						localMinSum = 0
					}
				}
				if localMinSum >0
			*/
		}
		/*
			currentSumMin = Min(currentSumMin, localMinSum)
			minCoinCache[sum] = currentSumMin
		*/
	}
	fmt.Printf("\nMaps Values :\n %v", minCoinCache)
	return minCoinCache[targetSum]

}
func main() {

	testCoins := []int{1, 2, 5}
	tagetVal := 4
	r := coinChange(testCoins, tagetVal)
	fmt.Println(r)

	r1 := CountStacks(10)
	fmt.Println(r1)

	testCoins2 := []int{1, 2, 5}
	MinmumNumberOfCountMap(testCoins2, 999)
}
