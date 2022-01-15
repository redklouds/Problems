package main

import (
	"testing"
)

/*
A note on compiler optimisations
Before concluding I wanted to highlight that to be completely accurate,
 any benchmark should be careful to avoid compiler optimisations eliminating
 the function under test and artificially lowering the run time of the benchmark.
*/
var result int

func BenchmarkBruteeForce(b *testing.B) {
	//BenchmarkBruteeForce-16    	36923304	        32.41 ns/op	       0 B/op	       0 allocs/op
	testNums := []int{2, 1, 3}

	target := 35
	var r int
	for i := 0; i < b.N; i++ {
		r = CombinationSumIV(testNums, target)
	}
	result = r
}

func TestCombinationSumIV_BruteForce(t *testing.T) { //ok  	combination_sum	0.064s
	//O(2^n)
	testNums := []int{1, 2, 3}
	expected := 7
	target := 4
	if actual := CombinationSumIV(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}

	testNums = []int{2, 1, 3}
	expected = 1132436852
	target = 35
	if actual := CombinationSumIV(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func BenchmarkWithCache(b *testing.B) {
	testNums := []int{2, 1, 3}

	target := 35
	var r int
	for i := 0; i < b.N; i++ {
		r = CombinationSumIV_WithCache(testNums, target)
	}
	result = r //fgor compuiler optmization if youm dont make function equal to something
	//it will kill off if it has not reference
}
func TestCombinationSumIV_WithCache_Memoization(t *testing.T) {
	testNums := []int{1, 2, 3}
	expected := 7
	target := 4
	if actual := CombinationSumIV_WithCache(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}

	testNums = []int{2, 1, 3}
	expected = 1132436852
	target = 35
	if actual := CombinationSumIV_WithCache(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func BenchmarkWithDP(b *testing.B) {
	testNums := []int{2, 1, 3}

	target := 35
	var r int
	for i := 0; i < b.N; i++ {
		r = CombinationSumIV_DP(testNums, target)
	}
	result = r //fgor compuiler optmization if youm dont make function equal to something
	//it will kill off if it has not reference
}

//O(n) alot alot faster
func TestCombinationSumIV_With_DP(t *testing.T) {
	testNums := []int{1, 2, 3}
	expected := 7
	target := 4
	if actual := CombinationSumIV_DP(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}

	testNums = []int{2, 1, 3}
	expected = 1132436852
	target = 35
	if actual := CombinationSumIV_DP(testNums, target); actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}
