package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//TODO:
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(test)
	fmt.Println(res)
}
