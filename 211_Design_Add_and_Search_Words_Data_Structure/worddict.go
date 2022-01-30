package main

type TrieNode struct {
	isWord   bool
	Children [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: [26]*TrieNode{},
	}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {

	return WordDictionary{
		root: NewTrieNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {

	if word != "" {
		cur := this.root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			if cur.Children[index] == nil {
				//initalize
				cur.Children[index] = NewTrieNode()
			}
			cur = cur.Children[index]
		}
		cur.isWord = true
	}
}

func (this *WordDictionary) Search(word string) bool {

	if word == "" {
		return false
	}
	return this.searchHelper(word, 0, this.root)
}

func (this *WordDictionary) searchHelper(word string, start int, ptr *TrieNode) bool {
	cur := ptr
	for i := start; i < len(word); i++ {
		if cur == nil {
			break
		} else {
			//not nil
			if word[i] != '.' {
				idx := word[i] - 'a'
				cur = cur.Children[idx]
			} else {
				//store the current node
				temp := cur
				//we found a wildcard, need to search all children
				for j := 0; j < 26; j++ {
					cur = temp.Children[j]
					if this.searchHelper(word, i+1, cur) {
						return true
					}

				}
				return false
			}
		}

	}
	return cur != nil && cur.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

//basically the tricky thing is when you have wildcard charactres, we basically only have to modify the search functionality, to run O(n^2), when we encounter a wildcard character skip we will have to run a search for the next character for EVERY next character

//so for search
//if char is not a wildcard, then look for child like normal
//if char is wildcard, for each child node does the next char exists in any of my children node
// ie if "b.ad" when we are on 'b' and we see that our next char is wildcard, for each child of b skip a char and ask for each of the children
