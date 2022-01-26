package main

import "testing"

func TestMain(t *testing.T) {
	//TODO:

	text := "loonbalxballpoon"
	exp := 2
	if act := maxNumberOfBalloons(text); act != exp {
		t.Error()
	}
}
