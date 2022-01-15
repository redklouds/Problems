package main

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)

	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[0] = 0

	for i := 0; i <= k+1; i++ {
		temp_prices := make([]int, n)
		copy(temp_prices, prices)

		for _, val := range flights {
			if prices[val[0]] == math.MaxInt32 {
				continue
			}
			if prices[val[0]]+ val[2] < temp_prices[val[1]] {
				temp_prices[val[1]] = prices[val[0]] + val[2]
			}
		}
		prices = temp_prices
	}

	if prices[dst] == math.MaxInt32{
		return -1
	}
	return prices[dst]
}
