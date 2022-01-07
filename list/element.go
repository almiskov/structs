package list

// element represents an element of doubly linked list
type element[T any] struct {
	prev, next *element[T]
	v          T
}

// Prev returns the previous element of the list or nil if element is tail
func (n *element[T]) Prev() *element[T] {
	return n.prev
}

// Prev returns the next element of nil if element is head
func (n *element[T]) Next() *element[T] {
	return n.next
}
