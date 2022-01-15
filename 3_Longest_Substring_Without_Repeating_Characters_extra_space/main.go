package main

//the idea is you can use an extra space such as a charSet -> you iterate through each char,
//when you haven't see this char before, we want to use what's called a 'SLIDING WINDOW'
//basically when we encounter a char that has been seen before, what we do is we remove the all values up to that repeated character, in search of a longer
//common string, 'imagine a sliding window' ['|ge'eksforg'|e|eks']
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return len(s)
	}
	charSet := make(map[byte]int)

	i, j := 0, 1
	charSet[s[i]] = 0
	maxLength := 1
	curLen := 1
	for j < len(s)-1 {

		if index, ok := charSet[s[j]]; ok {
			//we found something we have seen before
			//we have to move i up where this value is +1 as to 'remove it' slide our window forward, in search of greater things
			for idx := i; idx <= index; idx++ {
				//erase the values from the window
				delete(charSet, s[idx])
			}
			i = index + 1 //store stores the index we last seen these values at moving it over 1 means we erase it from our boundaries
			charSet[s[j]] = j
			maxLength = Max(maxLength, curLen)
			curLen = 1
		} else {
			//we haven't seen this char yet
			//place it in the store
			charSet[s[j]] = j
			curLen++
		}
		j++
	}
	return maxLength
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
