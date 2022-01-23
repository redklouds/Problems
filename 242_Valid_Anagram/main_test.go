package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:

	s := "anagram"
	s2 := "nagaram"

	expected := true
	if actual := isAnagram(s, s2); actual != expected {
		t.Error()
	}
}
