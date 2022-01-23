package main

//2 methods, youi can sort the string, and this will run you O(n * log k) where k is the length of the longest stribng
//n is the number of strings,

//we can do this in O(m * n) meaning that we can actually runo(n for each string) for M for each length

//we want to group by frequency STRING
//basically if we can build a hashmap key, using a frequency string as unique key
//ex "a1b2c2" = abbcc , so when you look at bcbca also equals "a1b2c2" key
func groupAnagrams(strs []string) [][]string {
	//brute force, write a anagram function,
	//then check for each word, which group is which belong to which anagram?

	//if anagrames are anagrams of each other they are seperate keys, their own group
	//store := make(map[string][]string)
	anagramGroupingStore := make(map[[26]int][]string)
	for _, str := range strs {
		freqCharKey := [26]int{} //26 characters in albephet
		for i := 0; i < len(str); i++ {
			//do the bottom to EQUALIZE the indexs
			//basically lower case 'a' is 97 in ASCII(byte)
			//WE KNOW the values are ALL LOWER CASE ALPHABET
			//so for instance if we want to start at 0(a) -> z()
			//if str[i] == 97 (a) if we subtract 'a' from that 97 - 97 == 0 YES then
			//same thing if we encounter 'b' str[i] == 98 (b) minus 97 = 1 98-97 = 1(b) that's how we have
			//we can have an array of 0 -> 26 for all characters of alphanet in an array
			freqCharKey[str[i]-'a']++

		}
		//after finishing this current charayter our freqCharKey has
		//THE FREQUENCYT KEY , lets look if our hash map contains this freq key, if so add
		//the current str as a anagram for this grouping
		anagramGroupingStore[freqCharKey] = append(anagramGroupingStore[freqCharKey], str)
	}

	//after processing all the valuhes into their respective freqeuncies, we need to extrct all the values

	res := make([][]string, len(anagramGroupingStore))
	i := 0
	for _, val := range anagramGroupingStore {
		res[i] = val
		i++
	}
	return res
}

func IsAnagram(s string, s1 string) bool {
	if len(s) != len(s1) {
		return false
	}
	store := make(map[rune]int)
	for _, rn := range s {
		store[rn]++
	}
	for _, rn := range s1 {
		store[rn]--
		if store[rn] == 0 {
			delete(store, rn)
		}
	}

	return len(store) == 0
}
