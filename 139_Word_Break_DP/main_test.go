package main

import (
	"testing"
)

//if a substring exists in the string, return true
func TestContainsSubstring(t *testing.T) {

	testWord := "Leet"

	TestDicWords := []string{"Leet1", "code"}

	expected := true

	if actual := containsSubString(testWord, TestDicWords[0], 0); actual != expected {
		t.Errorf("Expected %t, Got %t", expected, actual)
	}
}
