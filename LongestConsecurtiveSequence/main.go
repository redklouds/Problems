package main

import (
	"fmt"
)


//idea is to use a map to store the values, then for each elemetn in the array
//check wether there exists a LOWER bound/number in the seqeunce that exists in the array
//since we will start out counting/sequence at that LOWER number instead of the current number
//which is  a waste of processing power
func longestConsecutive(nums []int) int {

	if len(nums) < 1 {
		return 0
	}
	solMax := 0
	numMap := make(map[int]bool) //default zero-based (false) if doesn't exists
	ArrayToMap(nums, numMap)

	for _, num := range nums {
		//check if this current number is the LOWEST number for starting our seqeunce 
		//if there exists a LOWER number that we've seen before skip this guy, since it means
		//we have reached somewhere in the middle
		if !numMap[num - 1] {
			//if false, it means that a LOWER number DOES NOT EXISTS, therefore we should NOT skip this guy
			//and start our number counting from here
			local_max := 0
			
			//loop as long as the next sequence exists
			for numMap[num] { //grab the current element to increment correctly
				local_max++
				num++

			}
			solMax = Max(solMax, local_max)
		}
	}
	return solMax
}
//brute force, idea is to have O(n) memory
//to store all the items into a dict, and then look for them in sequence

func longestConsecutive1(nums []int) int {

	if len(nums) < 1 {
		return 0
	}
	//basically put everyhting into the map
	//iterate over the array
	//checking if Arr[i] - 1 exists, because if it does, we KNOW that Arr[i] is the start of the lowest consecutive number
	//ie [ 100, 200, 1,3,2] if i == 1 which is 200, we check if arr[i] - 1 exists which is 199, if 199 exist we know that this is the lower
	//starting popint, so we should check from 199, 200 ..., to last, otherwise we KNOW NO LESSER NUMBER exists at arr[i] therefore inferring that
	//the current arr[i] is the lowest/starting point
	curMax := 0
	numMap := make(map[int]bool)
	ArrayToMap(nums, numMap)

	for idx := range nums {
		//iterations basically the same as for i:0; i < len(nums);i++ //this goes from 0 to Len(arr)-1 or the collection index -1
		start := nums[idx] - 1 //we start with checking wether our current element is the lowest element or does a LOWER eelement exist

		local_max := 0
		if _, ok := numMap[start]; !ok {
			start = nums[idx]
		}

		for {
			if _, ok := numMap[start]; ok {
				local_max++
				start = start + 1
				continue
			}
			break
		}
		curMax = Max(local_max, curMax)

	}
	return curMax
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func ArrayToMap(nums []int, numsMap map[int]bool) {
	for _, val := range nums {
		numsMap[val] = true
	}
}
func main() {
	start := 0
	start++
	testData := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(testData)
	fmt.Printf("Longest Consecurtive: %d", res)

	tmp := make(map[int]bool)

	if !tmp[6969] {
		fmt.Printf("Hello world bby %v", tmp)
	}

	fmt.Println("\nExists?", tmp[32423324])
}
