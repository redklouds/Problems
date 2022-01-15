package main

import "fmt"

//lexicograhpically sorted means to be dictionary sorted
func verifyLexicographicallySorted(wordlist []string) bool {

	if len(wordlist) < 2 {
		return true
	}

	//to verify we don't need to ierate through all words
	//we just need to check if the next adjacent word is sorted from us
	//that way we know that the current sequence of words are sorted
	//therefore taking a stepwise approach
	for i := 0; i < len(wordlist)-1; i++ {
		//don't iterate to the end iterate until n-2, index
		word1 := wordlist[i]
		word2 := wordlist[i+1]
		shortedLength := len(word1)
		if len(word2) < shortedLength {
			shortedLength = len(word2)
		}

		//iterate over collection
		for w := 0; w < shortedLength; w++ {
			//if the two values differ we need to check that they
			//are in alphabetical order
			if word1[w] != word2[w] && word1[w] > word2[w]{
				if word1[w] > word2[w] {
					return false
				} else {
					//if the above rule stands that word1's diff char is SMALLER than word2's
					//then we break check the next word
					break
				}
			}
		}
		//finished we have confirmed that the prefix are the same
		//check the order
		if len(word1) > len(word2) {
			return false
		}

	}
	return true
}

func main() {
	WL := []string{
		"ABODY",
		"ABOY",

		"MEN",
	}
	isLexiSorted := verifyLexicographicallySorted(WL)
	fmt.Printf("Is the Collection %v, Lexi Sorted? : %t", WL, isLexiSorted)
}
