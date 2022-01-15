package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:
	nums := []int{1, 2, 3}
	actual := permute(nums)


	nums = []int{1}
	actual = permute(nums)

	nums = []int{0,1}
	actual = permute(nums)
	
	fmt.Println(actual)
}
