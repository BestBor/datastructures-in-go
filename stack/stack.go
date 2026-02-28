package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(element T) {
	s.items = append(s.items, element)
}

func (s *Stack[T]) Pop() (T, bool) {

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove, true
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Stack[T]) String() string {
	var b strings.Builder
	for i, item := range s.items {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(&b, item)
	}
	return b.String()
}
