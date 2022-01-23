package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	s := "bab"
	s2 := "aba"

	expected := 1
	if actual := minSteps(s, s2); actual != expected {
		t.Error()
	}

	s = "leetcode"
	s2 = "practice"

	expected = 5
	if actual := minSteps(s, s2); actual != expected {
		t.Error()
	}

	s = "anagram"
	s2 = "mangaar"

	expected = 0
	if actual := minSteps(s, s2); actual != expected {
		t.Error()
	}
}
