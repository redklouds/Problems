package main

import (
	"fmt"
	"testing"
)

func TestPascalTriangle(t *testing.T) {

	i := 3
	v := getRow(i)
	fmt.Println(v)
}
