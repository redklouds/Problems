package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	s := "abcabcbb"

	expected := 3
	if actual := lengthOfLongestSubstring(s); actual != expected {
		t.Error()
	}

	s = "bbbbb"
	expected = 1
	if actual := lengthOfLongestSubstring(s); actual != expected {
		t.Error()
	}

	s = "pwwkew"
	expected = 3
	if actual := lengthOfLongestSubstring(s); actual != expected {
		t.Error()
	}

}
