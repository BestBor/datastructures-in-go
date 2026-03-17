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

func (t *Trie) traverse(s string) *trieNode {
	current := t.root
	for _, ch := range s {
		if _, exists := current.children[ch]; !exists {
			return nil
		}
		current = current.children[ch]
	}
	return current
}

func (t *Trie) Search(word string) bool {
	n := t.traverse(word)
	return n != nil && n.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.traverse(prefix) != nil
}

func wordsNode(n *trieNode, prefix string, result *[]string) {
	if n.isEnd {
		*result = append(*result, prefix)
	}
	for ch, child := range n.children {
		wordsNode(child, prefix+string(ch), result)
	}
}

func (t *Trie) Words() []string {
	var result []string
	wordsNode(t.root, "", &result)
	return result
}

func deleteNode(n *trieNode, word string, depth int) *trieNode {
	if n == nil {
		return nil
	}
	if depth == len([]rune(word)) {
		n.isEnd = false
		if len(n.children) == 0 {
			return nil
		}
		return n
	}

	ch := []rune(word)[depth]
	n.children[ch] = deleteNode(n.children[ch], word, depth+1)
	if n.children[ch] == nil {
		delete(n.children, ch)
	}
	if len(n.children) == 0 && !n.isEnd {
		return nil
	}
	return n
}

func (t *Trie) Delete(word string) bool {
	if !t.Search(word) {
		return false
	}
	t.root = deleteNode(t.root, word, 0)
	t.size--
	return true
}
