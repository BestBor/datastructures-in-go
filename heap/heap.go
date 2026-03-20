package heap

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
