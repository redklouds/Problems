package main

import "fmt"

func main() {

	num := 0011
	if num&1 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	fmt.Println(num)
}
