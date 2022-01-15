package main

//O(n) solutono moving the left and right pioiters, because we know its sorted
//and twe know that we can move it around
//however there is a faster soluton LogN
func twoSum(nums []int, target int) []int {

	if len(nums) < 2 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1

	for left < right {

		val := nums[left] + nums[right]
		if val == target {
			return []int{left + 1, right + 1}
		} else if val > target {
			//move right down
			right = right - 1
		} else {
			left = left + 1 //move left up since the val is less than target
		}
	}
	return make([]int, 2)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func main() {

}
