package main

//O(n) runtime, O(n) space
func isAnagram(s string, t string) bool {
	store := make(map[rune]int)
	for _, ch := range s {
		store[ch]++
	}
	for _, rn := range t {
		//for each char in this, we need to make sure we are able to remove all the chars
		//in the store
		store[rn]--
		if store[rn] == 0 {
			delete(store, rn)
		}
	}
	return len(store) == 0
}
