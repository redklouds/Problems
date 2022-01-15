package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:
	word := "genetic"

	obj := Constructor()
	obj.Insert(word)
	if param_2 := obj.Search(word); param_2 != true {
		t.Error()
	}
	prefix := "gen"
	if param_3 := obj.StartsWith(prefix); param_3 != true {
		t.Error()
	}
}

func TestMain2(t *testing.T) {
	//TODO:
	word := "apple"

	obj := Constructor()
	obj.Insert(word)

	if param_2 := obj.Search(word); param_2 != true {
		t.Error()
	}
	word = "app"
	if param_2 := obj.Search(word); param_2 != false {
		t.Error()
	}

	prefix := "app"
	if param_3 := obj.StartsWith(prefix); param_3 != true {
		t.Error()
	}


	word = "app"
	obj.Insert(word)
	if param_2 := obj.Search(word); param_2 != true {
		t.Error()
	}
}
