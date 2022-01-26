package main

//BRUTE FORCE, easiest

func maxNumberOfBalloons(text string) int {

	//sub problem break down

	//1: we want to go over the string,and collect all the count of characters
	charCountMap := make(map[byte]int) //consist of english char only
	for i := 0; i < len(text); i++ {
		charCountMap[text[i]]++
	}

	//2 have a function that given the hashmap and a word see if it can build that work from it

	//3 have a loop that will check as long as a b exist in the map run the helper function

	wordCount := 0
	for charCountMap['b'] != 0 {
		if wordExists("balloon", charCountMap) {
			wordCount++
		} else {
			break
		}
	}
	//4 remove the char from the map as we check

	return wordCount
	//O(n) space to store the char in the map
	//O(n) runtime, because the worst case we are just running through each char once

}

//given a word to search for, return if the word can be made using the string texts
func wordExists(word string, charCountMap map[byte]int) bool {
	for i := 0; i < len(word); i++ {
		//this char doesn't exists or we are out of the chars aval
		if charCountMap[word[i]] == 0 {
			return false
		}
		//remove a count
		charCountMap[word[i]]--
	}
	return true
}
