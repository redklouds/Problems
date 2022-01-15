package main

import "fmt"

//alien dictionary, lexicographical sorts AKA dictionary sort
//we are given a list of words that are SORTS by thes Alien rules
//we know that the alien language is using LEXICOGRAPHICAL sort to sort their alient
//langauge
//reverse engineeer..
// if we know how LEXICOGRAPHICAL algo works, we can undo it to find
//what is the rules/order that each alphabet goes in
//in this alien language

//TOPLOGICAL sort, DAG because each letter MUST COME BEFORE another, aka
//DEPENDENCY RESOLVER, if given a list of words, all words staring with 'T' MUST COME before 'D'
//do we have a T -> D in a graph relationship, we have to process all words
//that start with 'T' before we can start processing words wtih letter 'D'

//given a lexicographical sorted array of words in an alien language, retunr
//the rules/order of the language
func AlienDictionary(wordList []string) []byte {

	//O(n) for space, since i need to store O(N ), N nodes in an adjacent List
	//O(n) for space for the innode counter
	//as for runtime, O(N) worst case TOPLOGICAL traversal using a BFS/Queue approcah worst case i go through the entire thing (good)

	//first populate our innode degree map
	//and our adjList

	adjList := map[byte][]byte{} //golang doesn't have char datatype instead
	//strings are immutable byte array wrappers, whre in go each char is a byte sequence
	//a stirn gis a sequence of bytes , therefore chars are bytes

	innodeCount := map[byte]int{}

	//lets populate the two maps

	for i := 0; i < len(wordList)-1; i++ {
		//we want to not iterate to the end we want to stop n-2 index before
		word1 := wordList[i]
		word2 := wordList[i+1]

		//now we have the two words we need to populate
		//find the shorter word
		x := len(word2)
		if len(word1) < len(word2) {
			x = len(word1)
		}

		for a := 0; a < x; a++ {
			//we compare the indexs of the chars if they are different we want to create a node relationship to edge and vertices
			if word1[a] != word2[a] {
				if _, ok := adjList[word1[a]]; !ok {
					adjList[word1[a]] = []byte{}
				}

				if _, ok := adjList[word2[a]]; !ok {
					adjList[word2[a]] = []byte{}
				}

				if _, ok := innodeCount[word1[a]]; !ok {
					innodeCount[word1[a]] = 0
				}
				if _, ok := innodeCount[word2[a]]; !ok {
					innodeCount[word2[a]] = 0
				}

				//they are different at this point word2[a] DEPENDS on word1[a] to process before it can be proccceseed
				//in other words WORD1[a] has be come BEFORE word2[a]
				adjList[word1[a]] = append(adjList[word1[a]], word2[a]) //word1[a] is the vertex word2[a] is the edge/adjacent node
				//increment the innode indegreee count for adjacent node
				innodeCount[word2[a]]++
				break
			}
		}
	}

	fmt.Printf("Here is the data : %v : \n %v", adjList, innodeCount)
	//now we have populate remember accessing a map with a key doesn't exists return szero-value
	//we can perform a TOPOLOGICAL SORT/Traversal, if its a DAG it will have a solution otherwise a cycle exists therefore a contridition on the rules exists

	//lets use BFS, to do this

	var q []byte
	//queue in the zero indegree nodes first

	for key, val := range innodeCount {
		if val == 0 {
			q = append(q, key)
		}
	}

	//we have zero in degree nodes lets process them
	var resultArr []byte

	for len(q) > 0 {

		//remove this element,

		elm := q[0]
		q[0] = 0
		q = q[1:] //reslice/size
		resultArr = append(resultArr, elm)
		//add it to the result array

		//now look at its adjacent nodes
		for _, adjN := range adjList[elm] {
			innodeCount[adjN]--
			if innodeCount[adjN] == 0 {
				q = append(q, adjN)
			}
		}
		//decrement its adjacent nodes, and check if those are new zero degreee nodes
	}

	if len(resultArr) != len(adjList) {
		return []byte{}
	}
	return resultArr

}

func main() {
	testData := []string{
		"wrt",
		"wrf",
		"er",
		"ett",
		"rftt",
	}
	//literal def always needs ending comma ** in go
	res := AlienDictionary(testData)
	fmt.Printf("\nresult: %c", res) //%c is character from byte to char

}
