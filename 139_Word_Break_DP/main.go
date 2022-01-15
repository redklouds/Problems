package main

import "fmt"

//given a string and a dictionary of words, can the word be broken
//up using only the whole words inside dic

//ex S := "Leetcode" dic =["Leet", "code"] => true
func wordBreak(word string, dic []string) bool {
	cache := make([]bool, len(word))
	for ch := len(word) - 1; ch >= 0; ch-- {
		//going backwards
		containsSubWord := false
		for _, dicWord := range dic {
			res := containsSubString(word, dicWord, ch)
			if res {
				containsSubWord = true

				break
			}
		}
		cache[ch] = containsSubWord
	}
	return true
}

//given a word can it be broken down using just the exact words in wordBank
func wordBreak2(word string, wordBank []string) bool {

	cache := make([]bool, len(word)+1)
	cache[len(cache)-1] = true //set the last item in the cache to true, its our base case
	//the OBVIOUS CASE, where the end of the string will be try and true ONLY AND ONLY IF WE are able to
	//parse the entire string all the way throgh THINK**** if we move from top-down approach we will go through the entire string
	//from start(idx=0) and ccheck the dic to see if each of these sub index words exists, moving our cursor up, UNTIL WE HIT the END of the STRING
	//signifying that we successfully parsed the entire tring and word bank contains words that can be broken into

	for selectedIdx := len(word); selectedIdx >= 0; selectedIdx-- {
		//starting from the back
		//if the current substring of word is less than the test word
		for _, dicWord := range wordBank {
			//if the current Dictionary word is longer than the substring word (obviously if its shorter than there is NO MATCH)
			if selectedIdx+len(dicWord) <= len(word) && word[selectedIdx:selectedIdx+len(dicWord)] == dicWord { //we do this because idx +len(dicWrod) becaue we are COMPAREING the SUBSTRING of the current index -->> to the end //ex dicWord ="CAR" and selected INdex is just "S" we know already not gonna work
				//and the substring matches the current word
				//so if its smaller substring AND the current substring equals the dictionary word, update the cache
				//signify  "If you start at this index, does there exists a matching dictionary word"
				cache[selectedIdx] = cache[selectedIdx+len(dicWord)]
			}
			if cache[selectedIdx] {
				break
			}
			//then we want tos et the cache value
		}
	}
	return cache[0]
}

func dfs(cache map[string]bool, word string) bool {
	return false
}

func containsSubString(word string, dicWord string, wordStartIndex int) bool {
	if len(dicWord) > len(word) { //obv. if the word we are comparing is larger than ourself
		//there is no match
		return false
	}
	for chIdx := range dicWord {
		if dicWord[chIdx] != word[chIdx+wordStartIndex] {
			return false
		}
	}
	return true
}

func main() {
	/*
		testWord := "cars"                     //"catsandog"                                  //"applepenapple"                              //"Leetcode"
		testDic := []string{"car", "ca", "rs"} //"cats", "dog", "sand", "and", "cat"} //"apple", "pen"} //"Leet", "code"}
		fmt.Printf("Word can be broken down? : %t \n", wordBreak(testWord, testDic))

		fmt.Printf("Word2 can '%s' be Broken Down with '%v' : %t\n", testWord, testDic, wordBreak2(testWord, testDic))
	*/
	test2Word := "Leetcode"
	testDic2 := []string{"Leet", "code"}
	fmt.Printf("Word2 can '%s' be Broken Down with '%v' : %t\n", test2Word, testDic2, wordBreak2(test2Word, testDic2))
	fmt.Printf("%t", test2Word[4:len(test2Word)])
	fmt.Printf("%t", test2Word[4:len(test2Word)] == "code")
}
