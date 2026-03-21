package heap

import "fmt"

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// This method allows us to make it maxHeap or minHeap
func New[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func NewHeapOnSlice[T any](data []T, less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		data: make([]T, len(data)),
		less: less,
	}
	copy(h.data, data)
	return h
}

func (h *Heap[T]) Len() int { return len(h.data) }

func (h *Heap[T]) IsEmpty() bool { return len(h.data) == 0 }

func (h *Heap[T]) Peek() (T, bool) {
	if h.IsEmpty() {
		fmt.Println("heap: Peek called on empty Heap")
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Push(val T) {
	h.data = append(h.data, val)
	// Adjust is pending
}

func (h *Heap[T]) adjustUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.less(h.data[i], h.data[parent]) {
			h.data[i], h.data[parent] = h.data[parent], h.data[i]
			i = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) adjustDown(i int) {
	n := len(h.data)
	for {
		better := i
		l, r := 2*i+1, 2*i+2
		if l < n && h.less(h.data[l], h.data[better]) {
			better = r
		}
		if better == i {
			break
		}
		h.data[i], h.data[better] = h.data[better], h.data[i]
		i = better
	}
}

func (h *Heap[T]) String() string {
	return fmt.Sprintf("%v", h.data)
}
