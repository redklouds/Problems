package main


//O(n) runtime O(1) space
func findDuplicates(nums []int) []int {
	// what we know, we know that the values are between [1,n] meaning that zero index is not used

	//we know that the values are only once or twice appearing

	//we know that it constant space

	//we know that if we shift the current values -1 we will fit all values in the array, ie
	//val 8 fits in n-1 spot 8 values = 0 -> 7 spots len(8)-1 = 7 = 8 is in the last spot

	//we can walk the array and flip the negative signs of values we have seen, if we encounter a value
	//thats already negative that means we have seen this before therefore is a duplicate

	//we also need to caution that we will be overlapping some of our negatigve, so using the absoluite value will save us from modifying the ACTUAL underlying data

	//assumptions values are only seen TWICE or ONCE, and are ALL POSITIVE, between 1,n

	//1: walk the array, getting the new index of our current value to mark it as seen
	var result []int
	for _, val := range nums {

		//1 the vals 'marker' index is one less than self since we are working with [1,n]
		markerIdx := Abs(val) - 1
		//2 check the markerIdx if it is already marked record it, otherwise we will mark it
		if nums[markerIdx] < 0 {
			result = append(result, markerIdx+1) //remember the markerIdx REPERSENTS the value MINUS 1, SEE above, to inverse this we add 1 to get the actual value the markerIdx Repersents
		} else {
			nums[markerIdx] = -1 * Abs(nums[markerIdx]) //Abs helps us preserve our negative flag
		}
	}

	return result

}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
