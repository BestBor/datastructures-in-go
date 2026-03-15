package bst

import "cmp"

// Node: Represents a BST Node
type Node[T cmp.Ordered] struct {
	key  T
	lSon *Node[T]
	rSon *Node[T]
}

// BST: Represents the base structure that references the root node
type BST[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func (t *BST[T]) Size() int {
	return t.size
}

// Insert: adds node with k value to de BST
func (t *BST[T]) Insert(k T) {
	if t.root == nil {
		t.root = &Node[T]{key: k}
		t.size++
		return
	}
	if t.root.insert(k) {
		t.size++
	}
}

// Insert (recursive method), returns true if inserted
func (n *Node[T]) insert(k T) bool {
	if k < n.key {
		if n.lSon == nil {
			n.lSon = &Node[T]{key: k}
			return true
		}
		return n.lSon.insert(k)
	} else if k > n.key {
		if n.rSon == nil {
			n.rSon = &Node[T]{key: k}
			return true
		}
		return n.rSon.insert(k)
	}
	return false
}

// Search find value on BST
func (t *BST[T]) Search(k T) *Node[T] {
	return searchNode(t.root, k)
}

// searchNode:
func searchNode[T cmp.Ordered](n *Node[T], k T) *Node[T] {
	if n == nil {
		return nil // no encontrado
	}
	if k == n.key {
		return n
	} else if k < n.key {
		return searchNode(n.lSon, k)
	} else {
		return searchNode(n.rSon, k)
	}
}

func (t *BST[T]) Exists(k T) bool {
	return t.Search(k) != nil
}

// findMin: returns the node with the minimum key in the subtree rooted at n
func findMin[T cmp.Ordered](n *Node[T]) *Node[T] {
	for n.lSon != nil {
		n = n.lSon
	}
	return n
}

// Delete
func (t *BST[T]) Delete(key T) {
	var deleted bool
	t.root, deleted = deleteNode(t.root, key)
	if deleted {
		t.size--
	}
}

func deleteNode[T cmp.Ordered](n *Node[T], key T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}
	var deleted bool
	if key < n.key {
		n.lSon, deleted = deleteNode(n.lSon, key)
	} else if key > n.key {
		n.rSon, deleted = deleteNode(n.rSon, key)
	} else {
		// No children
		deleted = true
		if n.lSon == nil && n.rSon == nil {
			return nil, deleted
		}
		// One child
		if n.lSon == nil {
			return n.rSon, deleted
		}
		if n.rSon == nil {
			return n.lSon, deleted
		}
		// Two children: Replace with min RSon in subtree
		replacement := findMin(n.rSon)
		n.key = replacement.key
		n.rSon, _ = deleteNode(n.rSon, replacement.key) // ignore boolean (this one always exists)
	}
	return n, deleted
}

func heightNode[T cmp.Ordered](n *Node[T]) int {
	if n == nil {
		return 0
	}
	lHeight := heightNode(n.lSon)
	rHeight := heightNode(n.rSon)
	if lHeight > rHeight {
		return 1 + lHeight
	}
	return 1 + rHeight
}

func (t *BST[T]) Height() int {
	return heightNode(t.root)
}

// Traversals

type TraversalOrder int

const (
	InOrder TraversalOrder = iota
	PreOrder
	PostOrder
)

func walkNode[T cmp.Ordered](n *Node[T], order TraversalOrder, fn func(T)) {
	if n == nil {
		return
	}
	switch order {
	case InOrder:
		walkNode(n.lSon, order, fn)
		fn(n.key)
		walkNode(n.rSon, order, fn)
	case PreOrder:
		fn(n.key)
		walkNode(n.lSon, order, fn)
		walkNode(n.rSon, order, fn)
	case PostOrder:
		walkNode(n.lSon, order, fn)
		walkNode(n.rSon, order, fn)
		fn(n.key)
	}
}

func (t *BST[T]) Walk(order TraversalOrder, fn func(T)) {
	walkNode(t.root, order, fn)
}
