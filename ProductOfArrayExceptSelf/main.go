package main

import "fmt"

//https://dev.to/akhilpokle/product-of-array-except-self-a-mind-boggling-google-interview-question-1en1
//its all about generating the left most products
//breaking this problem into a smaller problem, we know we need to get the products
//of the left of i so 0 to i -1 and right i +1 to n
//so the res array first pass for left, stores the continuing product
//of every preceeding product
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	left := 1
	right := 1

	for i := 0; i < len(nums); i++ {
		res[i] = left
		left = left * nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = right * res[i]
		right = right * nums[i]
	}
	return res
}

func ProductExceptSelfLeft(nums []int) []int {

	left := make([]int, len(nums))
	mul := 1

	for i := 0; i < len(nums); i++ {
		left[i] = mul
		mul = mul * nums[i]
	}

	return left
}


func main() {
	testData := []int{-1, 1, 0, -3, 0}
	testData2 := []int{1, 2, 3, 4}
	res1 := ProductExceptSelf(testData2)
	fmt.Printf("%+v", res1)

	res2 := ProductExceptSelfLeft(testData)
	fmt.Printf("%+v", res2)
}
