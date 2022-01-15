package main

type TrieNode struct {
	Val      byte
	Key      bool // 'end of a word'
	Children map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[byte]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	// n := newTrieNode{}
	return Trie{
		root: NewTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	this.insert([]byte(word), this.root)
}

func (this *Trie) insert(word []byte, root *TrieNode) {
	node := root

	//we want to insert
	for _, ch := range word {
		if childNode, ok := node.Children[ch]; ok {
			//exists
			node = childNode
		} else {
			//the key does not exists so we should create it
			new_node := NewTrieNode()
			new_node.Val = ch
			node.Children[ch] = new_node
			node = new_node
		}
	}

	//we finished processing all the chars, lets mark the node as valid node
	//because its the LAST node
	node.Key = true
}

func (this *Trie) Search(word string) bool {
	wordArr := []byte(word)
	res := this.search(wordArr, this.root)
	return res
}

func (this *Trie) search(word []byte, root *TrieNode) bool {

	node := root

	for _, ch := range word {
		//fmt.Printf("Searching ch -> %c\n", ch)
		if val, ok := node.Children[ch]; !ok {
			return false
		} else {
			node = val
		}
	}
	if node.Key {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	wordArr := []byte(prefix)
	return this.seartsWith(wordArr, this.root)
}

func (this *Trie) seartsWith(word []byte, root *TrieNode) bool {

	node := root

	for _, ch := range word {
		//fmt.Printf("Searching ch -> %c\n", ch)
		if val, ok := node.Children[ch]; !ok {
			return false
		} else {
			node = val
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
