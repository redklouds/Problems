package main

import (
	"fmt"
	"testing"
)

func TestGenParm(t *testing.T) {
	test := 3

	res := generateParenthesis(test)

	fmt.Println(res)

}
