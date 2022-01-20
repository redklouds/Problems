package main

func main() {
}

func longestConsecutive(nums []int) int {

	//BRUTE FORCE worst case i will have to go through each element O(n^2) times

	//idea brute force, dicttats me to think it like this

	//using a hashtable store each value
	//O(N)) O(n) space
	//then for each value, see if the current value + 1 exists ? if it does, then continue to ask if the next value exists until it does not
	//cehck if that length of values superceded teh current lenght, this is O(n)

	numberSet := make(map[int]bool)

	//populated our number set
	for _, num := range nums {
		numberSet[num] = true
	}

	//the idea is the iterate to find if the current number we are on exists a even LOWER boundary, if that is true, such as 100 -> 100 -1 = 99  if 99 exists in the map
	//we can safely SKIP this 100 value, and contniue iterating until we hit 99 SINCE WE KNOW that 99 is a  lower number we don't need to start our counting here, the same thing goes for any value such as 4, if we reach 4, we can ask the numsMap does 3 (1 lower exsits?) if it does, then we can ski 4 and wait for 3 EVEN if 1 exists because eventually, we will also run into 2 and if we run into 1 and 0 doesn not 4exist we know that there isn't a lower nnumer that itself now so we can start counting from here'

	//start our iteration over nums
	consec_max := 0
	for _, num := range nums {
		if _, found := numberSet[num-1]; !found {
			//means a lower number DOES NOT exists, we want to start counting now..

			local_max := 0
			for numberSet[num] {
				//since it exists now, lets update the local max
				//update the num incremented by 1 (consectutive)
				local_max++
				num++
			}
			//until we cannot find the next
			consec_max = Max(consec_max, local_max)
		}
	}

	return consec_max

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
