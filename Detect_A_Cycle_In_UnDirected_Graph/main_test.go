package main

import "testing"

//given a GraphNode
//detect a cycle
/*
type GraphNode struct {
	Val byte
	AdjN []*GraphNode
}
*/

//test if that a cycle should Not exists returning true
func TestDetectACycle_Exists(t *testing.T) {

	nA := &GraphNode{
		Val: 'A',
	}

	nB := &GraphNode{
		Val: 'B',
	}

	nC := &GraphNode{
		Val: 'C',
	}

	nD := &GraphNode{
		Val: 'D',
	}

	nA.AdjN = []*GraphNode{
		nB,
	}

	nB.AdjN = []*GraphNode{
		nA,
		nD,
		nC,
	}

	nC.AdjN = []*GraphNode{
		nB,
		nD,
	}

	nD.AdjN = []*GraphNode{
		nB,
		nC,
	}
	//should return true if cycle exists
	if CycleExists := CycleExists(nA); CycleExists {
		t.Errorf("Cycle Should Exists but Got: %t", CycleExists)
	}
}
func Setup() {

}

//test if that a cycle should Not exists returning false
func TestDetect_Does_Not_Exists(t *testing.T) {
	nA := &GraphNode{
		Val: 'A',
	}

	nB := &GraphNode{
		Val: 'B',
	}

	nC := &GraphNode{
		Val: 'C',
	}

	nD := &GraphNode{
		Val: 'D',
	}

	nA.AdjN = []*GraphNode{
		nB,
	}

	nB.AdjN = []*GraphNode{
		nA,
		nD,
		nC,
	}

	nC.AdjN = []*GraphNode{
		nB,
	}

	nD.AdjN = []*GraphNode{
		nB,
	}

	expected := false
	actual := CycleExists(nA)
	//return true if cycle Exists, false if none
	if expected != actual {
		t.Errorf("Expected %t, Got: %t", expected, actual)
	}
}
