package list

// element represents an element of doubly linked list
type element[T any] struct {
	prev, next *element[T]
	v          T
}

// Prev returns the previous element of the list or nil if element is tail
func (e *element[T]) Prev() *element[T] {
	if e == nil {
		return nil
	}
	return e.prev
}

// Prev returns the next element of nil if element is head
func (e *element[T]) Next() *element[T] {
	if e == nil {
		return nil
	}
	return e.next
}

// removeFrom removes element from given list and decreases length of the list
func (e *element[T]) removeFrom(l *list[T]) {
	if e == l.head {
		l.head = e.Next()
		l.head.prev = nil
	} else if e == l.tail {
		l.tail = e.Prev()
		l.tail.next = nil
	} else {
		e.prev.next, e.next.prev = e.Next(), e.Prev()
	}
	l.len--
}
