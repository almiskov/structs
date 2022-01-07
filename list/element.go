package list

// element represents an element of doubly linked list
type element[T any] struct {
	prev, next *element[T]
	v          T
}

// Prev returns the previous element of the list or nil if element is tail
func (e *element[T]) Prev() *element[T] {
	if e == nil { return nil }
	return e.prev
}

// Prev returns the next element of nil if element is head
func (e *element[T]) Next() *element[T] {
	if e == nil { return nil }
	return e.next
}

