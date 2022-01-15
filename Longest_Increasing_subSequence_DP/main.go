package main

import "fmt"

func LongestSub(arr []int) int {

	return dfsSubSq(arr, 0, 1)
}

func dfsSubSq(arr []int, l int, seqAmt int) int {
	seqAmt++
	if l == len(arr)-1 {
		return seqAmt
	}

	/*
		if arr[l] < arr[parentIdx] {
			//we are increasing order so if my choice is LESS than my parent.. BOO that's an end
			return seqAmt
		}
	*/
	
	for cIdx := l + 1; cIdx < len(arr); cIdx++ {
		if arr[cIdx] >= arr[l] {
			seqAmt = Max(seqAmt, dfsSubSq(arr, cIdx, seqAmt))
		}

	}

	return seqAmt

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testData := []int{0, 1, 0, 3, 2, 3}
	longestSubSeq := LongestSub(testData)

	fmt.Printf("Longest Sub Seq: %d", longestSubSeq)
}
