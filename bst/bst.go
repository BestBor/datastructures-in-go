package bst

import "cmp"

// Node: Represents a BST Node
type Node[T cmp.Ordered] struct {
	Key  T
	LSon *Node[T]
	RSon *Node[T]
}

// BST: Represents the base structure that references the root node
type BST[T cmp.Ordered] struct {
	root *Node[T]
}

// Insert: adds node to de BST
func (t *BST[T]) Insert(k T) {
	if t.root == nil {
		t.root = &Node[T]{Key: k}
		return
	}
	t.root.Insert(k)
}

// Insert (recursive method)
func (n *Node[T]) Insert(k T) {
	if k < n.Key {
		if n.LSon == nil {
			n.LSon = &Node[T]{Key: k}
		} else {
			n.LSon.Insert(k)
		}
	} else if k > n.Key {
		if n.RSon == nil {
			n.RSon = &Node[T]{Key: k}
		} else {
			n.RSon.Insert(k)
		}
	}
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
	if k == n.Key {
		return n
	} else if k < n.Key {
		return searchNode(n.LSon, k)
	} else {
		return searchNode(n.RSon, k)
	}
}

// Exists is a boolean validator, T or F if element exists on BST
func (t *BST[T]) Exists(k T) bool {
	return existsNode(t.root, k)
}

func existsNode[T cmp.Ordered](n *Node[T], k T) bool {
	if n == nil {
		return false
	}
	if k == n.Key {
		return true
	} else if k < n.Key {
		return existsNode(n.LSon, k)
	} else {
		return existsNode(n.RSon, k)
	}
}

// findMin: returns next superior value in right son subtree
func findMin[T cmp.Ordered](n *Node[T]) *Node[T] {
	for n.LSon != nil {
		n = n.LSon
	}
	return n
}

// Delete
func (t *BST[T]) Delete(key T) {
	t.root = deleteNode(t.root, key)
}

func deleteNode[T cmp.Ordered](n *Node[T], key T) *Node[T] {
	if n == nil {
		return nil
	}
	if key < n.Key {
		n.LSon = deleteNode(n.LSon, key)
	} else if key > n.Key {
		n.RSon = deleteNode(n.RSon, key)
	} else {
		// No children
		if n.LSon == nil && n.RSon == nil {
			return nil
		}
		// One child
		if n.LSon == nil {
			return n.RSon
		}
		if n.RSon == nil {
			return n.LSon
		}
		// Two children: Replace with min RSon in subtree
		replacement := findMin(n.RSon)
		n.Key = replacement.Key
		n.RSon = deleteNode(n.RSon, replacement.Key)
	}
	return n
}
