package sets

import "fmt"

//https://www.probabilitycourse.com/chapter1/1_2_2_set_operations.php
/*
Sets are basically hash map but with ONLY a single eleemnt unique,
they do not repeat the element EX:
			the word “application”
			would equal		Set A= {a, p, l, i, c, t, o, n}, with 'p''a','i' not repeated



	Operations of a set

		Intersect -> Intersect(setA, SetB) intersect_set :produces a new set that has the elements that exists in BOTH sets
			of two sets A and B, denoted by A∩B, consists of all elements that are both in A and−−− B. For example, {1,2}∩{2,3}={2}
			means creating a set that only contains elements that exists in set A AND set B

			Impl: dictats that when checking if they both exists, we check the smaller set of the two.

		Union -> Union(SetA, SetB) union_set: produces a set that contains elements that are present in set A
			OR set B, the word "OR" is used to rep a union set, two list that contain elements in each set
			the union would EX; {1,2}∪{2,3}={1,2,3} basically makes a new SET with DISTINT values from both, a new set
			that has values from a or B


*/

type ByteSet struct {
	set map[byte]bool
}

func (set *ByteSet) Add(val byte) bool {
	_, found := set.set[val] //basically if the val is found its true if its not its fallse
	set.set[val] = true      // setthe value reguardless, its a map so nothing is broken just overwritten
	return !found            //False if already exists, meaning found is true !true = false, means that its already existsed
	//if found is false !false = true, meaning that its found
}

//returns a new set that contains values that exists in both SetA and SetB
func (setA *ByteSet) Intersect(setB *ByteSet) ByteSet {
	intersect_set := ByteSet{}
	//performing an intersect
	//its best to get the smaller of the sets
	if len(setA.set) > len(setB.set) {
		//setB is smaller, so we swap since we assumed that setA would be smaller
		//swap variables
		setA, setB = setB, setA
	}
	for k, _ := range setA.set {
		if setB.set[k] {
			//we iterate setA as the SMALLER of the two, and find values in SETB that also exists in setA and setB
			//if they do we add those to the new intersect set
			intersect_set.set[k] = true
		}

	}
	return intersect_set
}

func (setA *ByteSet) Union(setB *ByteSet) (unionSet ByteSet) {
	//values that exists both in A OR B
	union_set := ByteSet{}

	for k, _ := range setA.set {
		union_set.set[k] = true
	}
	for k, _ := range setB.set {
		union_set.set[k] = true
	}
	return union_set
}

func (set *ByteSet) Equal(setB ByteSet) bool {
	if len(set.set) != len(setB.set) {
		return false
	}
	for k, _ := range set.set {
		if !setB.set[k] {
			return false
		}
	}
	return true
}

func (set *ByteSet) Print() {
	for k, _ := range set.set {
		fmt.Printf("%c", k)
	}
	fmt.Println("")
}

//initalizaer

func NewByteSet() ByteSet {
	byteSet := ByteSet{
		set: make(map[byte]bool),
	}
	return byteSet
}
