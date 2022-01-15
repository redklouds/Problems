package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "cbbd"
	expected := "bb"

	if actual := longestPalindrome(s); actual != expected {
		t.Errorf("Got: %s wanted: %s", actual, expected)
	}

	s = "ac"

	expected = "c"

	if actual := longestPalindrome(s); actual != expected {
		t.Errorf("Got: %s wanted: %s", actual, expected)
	}

	s = "a"
	expected = "a"

	if actual := longestPalindrome(s); actual != expected {
		t.Errorf("Got: %s wanted: %s", actual, expected)
	}
}
func TestIsPalindrom(t *testing.T) {
	testString := "RACAR"
	expected := true

	if actual := isPalindrome(0, len(testString)-1, testString); actual != expected {
		t.Errorf("Got: %t, Expected: %t", actual, expected)
	}

	testString = "BOOB"
	expected = true

	if actual := isPalindrome(0, len(testString)-1, testString); actual != expected {
		t.Errorf("Got: %t, Expected: %t", actual, expected)
	}

	testString = "BABAD"
	expected = false

	if actual := isPalindrome(0, len(testString)-1, testString); actual != expected {
		t.Errorf("Got: %t, Expected: %t", actual, expected)
	}
}
