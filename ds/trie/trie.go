package trie

// size of possible characters in the trie
const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(w string) {
	currNode := t.root
	for _, c := range w {
		charIdx := c - 'a'
		if currNode.children[charIdx] == nil {
			currNode.children[charIdx] = &Node{}
		}
		currNode = currNode.children[charIdx]
	}
	currNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	currNode := t.root
	for _, c := range w {
		charIdx := c - 'a'
		if currNode.children[charIdx] == nil {
			return false
		}
		currNode = currNode.children[charIdx]
	}
	return currNode.isEnd
}
