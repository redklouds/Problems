package main

import "strings"

func alienOrder(words []string) string {
	// using a topological sort

	//first create the indegree map
	indegreMap := make(map[byte]int)
	//initalize it
	//the adj list
	adjList := make(map[byte][]byte)
	//iterate through the words
	//and each word we need to

	//go until i -2 since we want word1 tobe second to last

	for i := 0; i < len(words)-1; i++ {

		word1 := i
		word2 := i + 1
		//assume that word1 is smaler, but check to make sure we get the smaller length
		wLenToCompare := len(words[word1])
		if len(words[word2]) < wLenToCompare {
			wLenToCompare = len(words[word2])
		}

		w1 := words[i]
		w2 := words[i+1]

		for c := 0; c < len(w1); c++ {
			if _, found := adjList[w1[c]]; !found {
				adjList[w1[c]] = []byte{}
			}
			if _, found := indegreMap[w1[c]]; !found {
				//initalize it
				indegreMap[w1[c]] = 0
			}

		}
		for c := 0; c < len(w2); c++ {

			if _, found := indegreMap[w2[c]]; !found {
				//initalize it
				indegreMap[w2[c]] = 0
			}

			if _, found := adjList[w2[c]]; !found {
				adjList[w2[c]] = []byte{}
			}
		}

		//iterate through each char
		for cIdx := 0; cIdx < wLenToCompare; cIdx++ {
			//ensure that the value exists, if it does not INITALIZE IT else our topological sort wont work
			if _, found := adjList[w1[cIdx]]; !found {
				adjList[w1[cIdx]] = []byte{}
			}

			if _, found := adjList[w2[cIdx]]; !found {
				adjList[w2[cIdx]] = []byte{}
			}

			//now add the indegree update
			if _, found := indegreMap[w1[cIdx]]; !found {
				//initalize it
				indegreMap[w1[cIdx]] = 0
			}

			if _, found := indegreMap[w2[cIdx]]; !found {
				//initalize it
				indegreMap[w2[cIdx]] = 0
			}
			//since doing this will result in comparing BYTES instaed of the actual unicode char value its OK
			//to compare just the bytes of the char
			//now we have each char look for the first difference
			if w1[cIdx] != w2[cIdx] {
				//we found a difference
				//we know that word1 char Comes BEFORE word2

				//add the neighbor
				adjList[w1[cIdx]] = append(adjList[w1[cIdx]], w2[cIdx])
				indegreMap[w2[cIdx]]++
				break
			}

		}

	}

	result := TolopogicalSort(indegreMap, adjList)
	return result
	//perform the topolicail sort

}

func TolopogicalSort(indegreeMap map[byte]int, adjList map[byte][]byte) string {
	var sb strings.Builder

	//lets first find all nodes with zero indgree and enqueue them
	var q []byte
	for k, v := range indegreeMap {
		if v == 0 {
			q = append(q, k)
		}
	}

	for len(q) != 0 {
		//while que is not empty lets process
		c := q[0]
		q = q[1:]
		sb.WriteByte(c)
		//add to the sort
		//know search the neighbors
		//and decrement their innode degree, since its dependent has been processed

		for _, adjNode := range adjList[c] {
			indegreeMap[adjNode]--
			//check if the current node indegree has reached zero if it has enqueue it
			if indegreeMap[adjNode] == 0 {
				q = append(q, adjNode)
			}
		}
	}
	///now we check if there still exists
	//basically the adjList tells you all the NODES/chars existsant, if our results length does not have ALL the nodes in that list as keys, then we know that there existed a cycle somewhere, and one of the nodes still have indgree pointing to it, meanin a cycle exists somewhere
	resultString := sb.String()
	if len(resultString) != len(adjList) {
		return ""
	}
	return resultString
}
