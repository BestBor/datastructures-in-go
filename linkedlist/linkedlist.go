package linkedlist

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	data T
	next *node[T]
}

type LinkedList[T comparable] struct {
	head   *node[T]
	length int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Prepend(value T) {
	n := &node[T]{data: value}
	n.next = l.head
	l.head = n
	l.length++
}

func (l *LinkedList[T]) Length() int {
	return l.length
}

func (l *LinkedList[T]) Head() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.data, true
}

func (l LinkedList[T]) String() string {
	var b strings.Builder
	for n := l.head; n != nil; n = n.next {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(&b, n.data)
	}
	return b.String()
}

func (l *LinkedList[T]) DeleteWithValue(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return true
	}

	prev := l.head
	for curr := l.head.next; curr != nil; curr = curr.next {
		if curr.data == value {
			prev.next = curr.next
			l.length--
			return true
		}
		prev = curr
	}

	return false
}
