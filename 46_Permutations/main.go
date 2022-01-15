package main

import "fmt"

//basically understanding backtracking AKA reverting state
//we know that if we do this by hand,
//We know we have X amount of SPOTS -> for each spot we want to CHOOSE a value at that spot
//EX if we have [1,2,3] = we HAVE 3m SPOTS, for the first spot we have 3 choices, 1 or 2 or 3, once we choose a value
//there, then what we have left is we have 2 choices left, ex if i choose 3 for the first spot, i only have 1 or 2 as choices (ie 2 choices)
//now if i choose one of those say 2 the only choice i have is a 1 left (1 choice) thinking about this in an array
//we have CHOICES, 1 or 2 or 3, and a SPOT IE index, starting at ZERO, we know that its our CHOOSEN SPOT
//we choose to use this spot and cycle through putting all of our AVALIBLE CHOICES HERE, which will loop all values
//then we see that for EACH SPOT THAT WE CHOOSE, we need to move onto choosing the NEXT SPOT IE index 1, and do the same thing
//repeat until we have no more choices to do
//THE KEY IS HOWEVER for this to work we need a 'baseline' ideal data, meaning when we swap , we want to REVERT our changes
//so swapping back after the recursive call will put the values we expect to be there back to where they are meant ot be
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	resultSet := [][]int{}
	permuteHelper(nums, 0, &resultSet)
	return resultSet

}

func permuteHelper(nums []int, i int, resultSet *[][]int) {

	//base case
	if i >= len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		fmt.Println(nums)
		*resultSet = append(*resultSet, temp)
	} else {
		for a := i; a < len(nums); a++ {
			nums[i], nums[a] = nums[a], nums[i]
			permuteHelper(nums, i+1, resultSet) //make sure not to change i in the stack frame else
			//when we revert our state it will change the index we revert to
			nums[i], nums[a] = nums[a], nums[i]
		}
	}

}
