package main

import "testing"

func TestUniquePathsWithObsticles(t *testing.T) {

	obstacleGrid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	expected := 7

	if actual := minPathSum(obstacleGrid); actual != expected {
		t.Errorf("Got : %v, wanted: %v", actual, expected)
	}

	obstacleGrid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	expected = 7

	if actual := minPathSum(obstacleGrid); actual != expected {
		t.Errorf("Got : %v, wanted: %v", actual, expected)
	}

	/*
		obstacleGrid = [][]int{{0, 1}, {0, 0}}
		expected = 1

		if actual := uniquePathsWithObstacles(obstacleGrid); actual != expected {
			t.Errorf("Got : %v, wanted: %v", actual, expected)
		}

		obstacleGrid = [][]int{{1}}
		expected = 0

		if actual := uniquePathsWithObstacles(obstacleGrid); actual != expected {
			t.Errorf("Got : %v, wanted: %v", actual, expected)
		}

		obstacleGrid = [][]int{{0, 1}}
		expected = 0

		if actual := uniquePathsWithObstacles(obstacleGrid); actual != expected {
			t.Errorf("Got : %v, wanted: %v", actual, expected)
		}
	*/

}
