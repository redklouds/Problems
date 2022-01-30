package main

import "testing"

func TestMain(t *testing.T) {

	var wordDictionary WordDictionary = Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.AddWord("")
	if act := wordDictionary.Search("pad"); act != false {
		t.Error()
	} // return False
	if act := wordDictionary.Search("bad"); act != true {
		t.Error()
	} // return True
	if act := wordDictionary.Search(".ad"); act != true{
		t.Error()
	} // return True
	if act := wordDictionary.Search("b.."); act != true{
		t.Error()
	} // return True
}
