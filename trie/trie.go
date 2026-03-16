package trie

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
}

type Trie struct {
	root *trieNode
	size int
}

func newNode() *trieNode {
	return &trieNode{children: make(map[rune]*trieNode)}
}

func (t *Trie) Insert(newWord string) {
	current := t.root
	for _, ch := range newWord {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = newNode()
		}
		current = current.children[ch]
	}
	if !current.isEnd {
		current.isEnd = true
		t.size++
	}
}
