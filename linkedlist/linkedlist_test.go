package linkedlist

import "testing"

func TestNewLinkedList_IsEmpty(t *testing.T) {
	list := New[int]()

	if list.Length() != 0 {
		t.Fatalf("expected empty list length 0, got %d", list.Length())
	}

	if _, ok := list.Head(); ok {
		t.Fatal("expected no head value in empty list")
	}
}

func TestPrepend(t *testing.T) {
	list := New[int]()

	list.Prepend(10)
	list.Prepend(20)

	if list.Length() != 2 {
		t.Fatalf("expected length 2, got %d", list.Length())
	}

	value, ok := list.Head()
	if !ok {
		t.Fatal("expected head value, got none")
	}

	if value != 20 {
		t.Fatalf("expected head value 20, got %d", value)
	}
}
