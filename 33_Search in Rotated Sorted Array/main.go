package main

func Search(nums []int, target int) int {

	//need to find the sorted parituion and search there

	//set the left and right pointers

	//get the mid,

	//check if the target is mid and return it if true otherwise set your left and right
	//pointers

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] == target {
			return target
		}

		//otherwise lets check, which parition is sorted, because it will be easiest to find
		//the value there
		if nums[mid] > nums[left] {
			//the left is sorted search within here

			//if the right is sorted.. easyt to check does this target live within the boundaries?
			//if it doesn't no point on searching this one search the other partition
			if target >= nums[left] && target <= nums[mid-1] {
				//its between the left partiton, otherwise, its on the right size
				//we def want to search the left partition
				right = mid - 1
			} else {
				//otherwise we KNOW that the current target lies on the other partiton
				left = mid + 1
			}
		} else {
			//the right is sorted check here
			//same thing if the right is sorted we want to first ask.. does the target ly in this section
			//if it does then we want to make sure to see if it lies within these  boundaries
			//otherwise we know its in th eother side
			if target >= nums[mid+1] && target <= nums[right] {
				//we know it lives here
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
