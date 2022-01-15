package main

import "testing"

func TestMain(t *testing.T) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	exp := "wertf"
	if actual := alienOrder(words); exp != actual {
		t.Error()
	}

	words = []string{"z", "x"}
	exp = "zx"
	if actual := alienOrder(words); exp != actual {
		t.Error()
	}

	words = []string{"z", "x", "z"}
	exp = ""
	if actual := alienOrder(words); exp != actual {
		t.Error()
	}

	
	words = []string{"z", "z"}
	exp = "z"
	if actual := alienOrder(words); exp != actual {
		t.Error()
	}

}
