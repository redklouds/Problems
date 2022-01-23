package main

func minSteps(s string, t string) int {

	//BASICALLY O(N) space
	//o(n) runtime, we will just ask if there are differences and if there are
	//we can move things around and count those as stepss
	//What exists in t that doesn't exists in s, because we know that those are steps we need to take
	//to solve it

	//build the map, we can also try using a constant array to store the count of each character

	charCountStore := make(map[byte]int)

	for i := range s {
		charCountStore[s[i]]++
	}

	//we need to increment the step if we find that t has a value that DOESN"T exist in S,
	//that means that we need to go ahead and add that character over to s
	steps := 0
	for c := range t {
		if charCountStore[t[c]] == 0 {
			//we know that the current character DOES NOT exists in the first string s,
			//so we need to do something here, that somthing
			steps++
		} else {
			//we know that there exists a character in string t and s, so lets decrement the count
			charCountStore[t[c]]--
		}
	}
	//essentially number of differences
	return steps
}
