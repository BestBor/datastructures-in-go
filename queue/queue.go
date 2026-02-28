package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue: Adds a value to the queue
func (q *Queue[T]) Enqueue(element T) {
	q.items = append(q.items, element)
}

// Dequeue
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove, true
}

// Peek
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	return q.items[0], true
}

// Size
func (q *Queue[T]) Size() int {
	return len(q.items)
}
