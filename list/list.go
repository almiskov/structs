package list

type list[T any] struct {
	head, tail *element[T]
	len        int
}

func New[T any]() *list[T] {
	return new(list[T])
}

func (l *list[T]) add(v T) *element[T] {
	n := element[T]{v: v}

	if l.head == nil {
		l.head, l.tail = &n, &n
	} else {
		l.tail.next = &n
		n.prev = l.tail
		l.tail = &n
	}

	l.len++
	return &n
}

func (l *list[T]) first(p func(v T) bool) *element[T] {
	for e := l.head; e != nil; e = e.Next() {
		if p(e.v) {
			return e
		}
	}
	return nil
}

func (l *list[T]) find(p func(v T) bool) []*element[T] {
	els := make([]*element[T], 0)
	for e := l.head; e != nil; e = e.Next() {
		if p(e.v) {
			els = append(els, e)
		}
	}
	return els
}

func (l *list[T]) remove(p func(v T) bool) bool {
	if el := l.first(p); el != nil {
		if el == l.head {
			l.head = el.next
			l.head.prev = nil
		} else if el == l.tail {
			l.tail = el.prev
			l.tail.next = nil
		} else {
			el.prev.next, el.next.prev = el.next, el.prev
		}
		l.len--
		return true
	}
	return false
}