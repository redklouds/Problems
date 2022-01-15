package main

import (
	"fmt"
	"sort"
)

//there are 4 solutions
/*
	1: brute force nested for loops to find all pairs
	2: hashing having the values stored inside the map O(n) run and O(n) space
	3: sort and Inward step, two pointers one at the top and one at the end
		we will assume that if i and j are higher than target, move the j lower
		otherwise if i + j is lower we move the i higher, since all solutions exist
		in the range O(n log (n) ) for sorting, O(1) for storage
	4: Sort and Binary Search (O(n log (n))) run
*/

//we will allocate an map that will store all the values and their index
//the duplicates problem say [3,3], t=6, if we add at the beginning we overwrite
//the map with our 3's index at 1, if we check and add later, the firs time we see the value
//we wills earch immediately
//O(n) runtime for running through entire Nums and O(n) space for storeing inside the map
func TwoSumsHash(nums []int, target int) []int {
	valStore := make(map[int]int)
	solution := []int{0, 0}
	//check each value
	for i, v := range nums {
		valToFind := target - v
		val, exists := valStore[valToFind]
		if exists {

			//return []int{val, i}
			solution[0] = val
			solution[1] = i
			break
		} else {
			valStore[v] = i
		}
	}
	return solution
}

//this solution involves assuming that you understand
//if you sort numbers, and put two pointers one on the tail and one on the head
//adding head and tail equals a number, if you want a number below that number
//move the tail pointer down, else move it up if you want a higher number
//the combination of the pairs will exsists if a sum machines within the middle
func TwoSumInwardWalkAndSort(nums []int, target int) []int {
	//THIS DOES NOT WORK ANYMORE< ESP IF YOU NEED THE INDEXS OF THE VALUES< SINCE YOU SORT YOU MOVE THE INDEXS FROM THEIR ORIGINAL PLACE
	lowPtr := 0
	highPtr := len(nums) - 1

	result := make([]int, 2)
	//NEED TO SORT FIRST ELSE THIS DOESN;T WORK BECAUSE ASSUMING UPPER HALF IS GREATER
	//assume array is editable
	sort.Ints(nums)
	for lowPtr < highPtr {
		sum := nums[lowPtr] + nums[highPtr]
		if sum == target {
			result[0] = lowPtr
			result[1] = highPtr
			break
		} else if sum > target {
			//move high ptr down since the sum is over the target
			highPtr--
		} else {
			lowPtr++
		}
	}

	return result
}

func main() {
	nums := []int{3, 3}
	nums2 := []int{3, 2, 3}
	res := TwoSumsHash(nums, 6)
	res2 := TwoSumInwardWalkAndSort(nums2, 6)
	fmt.Printf("%+v \n", res)
	fmt.Printf("%+v \n", res2)

}
