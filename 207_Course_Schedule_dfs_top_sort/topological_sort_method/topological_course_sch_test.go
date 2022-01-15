package TopicalOrderingCourseSchdule

import "testing"

func NumCourseOutOfBoundTests(T *testing.T) {

	testData := [][]int{
		{1, 0},
	}

	res := canFinish(2, testData)
	if !res {
		T.Error("ERROR")
	}
}
