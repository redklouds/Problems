package main

import "testing"

func TestHouseRobber_Rob(t *testing.T) {
	testHouseMoney := []int{2, 7, 9, 3, 1}

	expected := 12

	if actual := rob(testHouseMoney); actual != expected {
		t.Errorf("Expected %d Got %d", expected, actual)
	}

	testHouseMoney = []int{1, 2, 3, 1}

	expected = 4

	if actual := rob(testHouseMoney); actual != expected {
		t.Errorf("Expected %d Got %d", expected, actual)
	}

	testHouseMoney = []int{2, 1, 1, 2}//rob house 1 then skip and rob last house = 4

	expected = 4

	if actual := rob(testHouseMoney); actual != expected {
		t.Errorf("Expected %d Got %d", expected, actual)
	}
}
