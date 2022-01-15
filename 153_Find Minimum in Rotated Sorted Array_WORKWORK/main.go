package main

import "fmt"

//traversing through the array , checking if the order is correctly made
func findMinBetter(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			//search the left partiton, if MID is smaller than the End that means 16(mid) ... 40(right)
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

//theres a kicker in this code in that the code does not work if you have 2 elements
//minus 0 index doesnt work
func findMin(nums []int) int {
	//the kicker hint is that the assumption that the array IS SORTED
	//IS Unique,
	//IS Ascending order
	//And we can check if its rotated

	if len(nums) == 1 {
		return nums[0] //return the first element
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	//check if its rotated
	//nums[0] < nums[n]
	if nums[0] < nums[len(nums)-1] {
		//is NOT rotated return the first element
		return nums[0]
	}

	//at this point er know the array is rotated
	//set values to track our left and rights
	left := 0
	right := len(nums) - 1

	//binary search , while left  idx doesn't equal and IS still smaller than our right index continue
	for right > left {

		//get the middle element, remember our right and left shift so we need to equalize our mid
		mid := left + (right-left)/2

		//two conditions to find the inflection point, or smallest
		//remeber assmption that we are goiung in ASCENDING order
		//if left of mid is GREATER than mid, then we have our smallest at mid i.e 7, 2 (mid), 3

		if nums[mid] < nums[mid-1] {
			//smallest is mid
			return nums[mid]
		}

		//if our right neighbor is SMALLER than us, we have our smallest, because order should be ascending order
		if nums[mid] > nums[mid+1] {
			return nums[mid+1] //ie 5,6(mid),2
		}

		//partitoning, key is that to assume num[0] is NOT smallest, but its a range, so if numbers are going up
		//than between nums[0] and mid
		//if mid is GREATER than nums[0] then we search the right mid+1 to end because we know that between mid and arr[0] that all the numbers are going to beascending from arr[0]
		//if mid is SMALLER than nums[0] then we search the left since if nums[0] is 30 ... mid is 16 somewhree in between 30 --> 16 will hold the lower, and going from 16 to the right will be larger numbers

		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			//search the left side
			right = mid - 1
		}
	}
	return -1
}

func main() {

	testData := []int{5, 1}

	testData1 := []int{5, 1, 2, 3, 4}
	res := findMin(testData)
	res2 := findMin(testData1)
	fmt.Println(res)

	fmt.Println(res2)
}
