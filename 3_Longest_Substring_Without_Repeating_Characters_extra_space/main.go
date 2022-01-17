package main

//o(n) o(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	charStore := make(map[byte]bool)
	//"dvdf" like d,v,d,f you if your second pointer is on the second d(idx 2) then you want
	//to erase the first d, by moving the first pointer forward UNTIL it encounters the char between the two windows
	//removing each one, once we find the char, we will move ptr1 over 1 forward
	ptr1 := 0
	ptr2 := 1
	maxLength := 1
	charStore[s[0]] = true
	for ptr2 != len(s) {
		if _, ok := charStore[s[ptr2]]; ok {
			//if ptr2 i found, we want to move ptr up until it hits it, reseting all other values along the way
			for i := ptr1; i < ptr2; i++ {
				//we will go ex [0(a),1(b),2(c),3(a)]
				delete(charStore, s[i]) //remember each char as we move out ptr1 forward
				if s[i] == s[ptr2] {
					ptr1 = i + 1
					break //when we find our index that has our character that is our dup we want to exit and leave
					//out ptr1 on the NEXT index
				}

			}
			//we removed
			//ptr1 = ptr2
		} else {
			//the value is different and we have not seen before,
			//lets go ahead and get max substring
			charStore[s[ptr2]] = true
			maxLength = Max(maxLength, len(charStore))
			ptr2++
		}

	}

	//we finish iterating through the enture striung
	return maxLength

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
