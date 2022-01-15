package main

import "fmt"

func longestCommonSubsequence(text2 string, text1 string) int {

	//make our dyanmic programming 2D matrix cache, then at the ends
	//makesure we signafy out of bounds cases and how to handle those

	cache := make([][]int, len(text1)+1)

	//initlize the rows
	for i := range cache {
		//each index set the rows to length of text2 +1
		cache[i] = make([]int, len(text2)+1)
	}

	//zero value based normal

	//so work from the backend of the matrix
	for row := len(cache) - 2; row >= 0; row-- {
		for col := len(cache[0]) - 2; col >= 0; col-- {

			//if text1[i] == text2[j]
			//then we have a match, we want to 1 diagonal
			//else
			//if they differ, we want to check the right pair and the bottom pair
			//get the Max ( text2[j + 1], text1[i+1])
			if text1[row] == text2[col] {
				//we have a match lets get the diagnoal
				cache[row][col] = 1 + cache[row+1][col+1]
			} else {
				//they are different we watch the current cell to be equal to the Larger Common SubSeqence from the two compares
				//ex bbcde
				//ex ace
				//both index[0] are different, we need to compare which is later, the word1 b compare with word2 (c,e) or word2 a compare with word1 (bbcde)
				cache[row][col] = Max(cache[row][col+1], cache[row+1][col])
			}

		}
	}
	return cache[0][0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	word1 := "abcde"
	word2 := "ace"
	fmt.Println(longestCommonSubsequence(word1, word2))
	testString := "Helloworld"

	for i := range testString {
		fmt.Printf("%c\n", testString[i])
	}
}
