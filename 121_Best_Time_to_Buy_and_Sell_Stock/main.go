package main

import (
	"fmt"
	"math"
	"time"
)

//NOTE THERE EXIST NO FUNCTION 'Max' in Go for Ints ONLY FOR FLOATS
//OR MIN for INts

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return -1
	}
	maxProfit := 0
	minLocal := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if prices[i] < minLocal {
			minLocal = prices[i]
		} else {
			maxProfit = Max(maxProfit, prices[i]-minLocal)
		}
	}
	return maxProfit

}

func MaxProfitBruteForce(prices []int) int {
	start := time.Now()
	if len(prices) == 0 {
		return -1
	}
	maxProfit := 0
	//localMin := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}
	if maxProfit <= 0 {
		return 0
	}
	els := time.Now().Sub(start)
	fmt.Println(els)
	return maxProfit
}
func main() {
	a := math.Max(5, 10)
	fmt.Println(a)

	prices := []int{7,6,4,3,1}
	res1 := MaxProfitBruteForce(prices)

	res2 := MaxProfit(prices)

	fmt.Printf("Brute Force >> : %d\n", res1)

	fmt.Printf("Brute Force >> : %d\n", res2)
}
