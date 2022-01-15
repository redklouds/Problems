package main

//O(n^3) because each character, will iterate the array again, and at each iteration
//we also check if its a palindrom looping aagain O(n^2) n* N * n
func longestPalindrome(s string) string {
	//start from the backkk
	//current palindrome
	resultPalindrome := ""
	for i := len(s) - 1; i >= 0; i-- {
		//for each char, see if you can find a palindrome and if you can see if its longer
		//than your current longest

		for nextChar := i; nextChar < len(s); nextChar++ {
			if isPalindrome(i, nextChar, s) {
				//found palindrome
				if len(s[i:nextChar+1]) > len(resultPalindrome) {
					resultPalindrome = s[i : nextChar+1]
				}
			}
		}
	}
	return resultPalindrome
}

//anothjer approach O(n^2)
//why don't we move outwards <--SelectedIdx -->, keep moving until two Lptr and Rptr are diff
//when they are same we have palindrome, move until they are different
func LongestPalindromeSubString(s string) string {

	currentLongest := "a"
	for i := 1; i <= len(s)-1; i++ {

		leftPtr := i - 1
		rightPtr := i + 1
		
		for s[leftPtr] == s[rightPtr]{
			//if they the same keep moving the ptrs outward until they are diff, and record
			//the length

		}
	}
	return currentLongest
}
func isPalindrome(l int, r int, s string) bool {

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
