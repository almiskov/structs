package list

type list[T any] struct {
	head, tail *element[T]
	len        int
}

func New[T any]() *list[T] {
	return new(list[T])
}

func (l *list[T]) add(v T) *element[T] {
	n := &element[T]{v: v}

	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}

	l.len++
	return n
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
		el.removeFrom(l)
		return true
	}
	return false
}

func (l *list[T]) removeMany(p func(v T) bool) []T {
	els := l.find(p)
	rm := make([]T, len(els))

	for i, e := range els {
		rm[i] = e.v
		e.removeFrom(l)
	}
	return rm
}

// TODO: implement methods: clear(), insertAt(idx int, v T), removeAt(idx int), indexOf(v T), forEach(func(v T)), map(???) (same as js Array.prototype.map())
// TODO: write descriptions
// TODO: try to implement some specific extension methods, e.g. Sum(), Min(), Max() fot list[int], if it's allowed
// TODO: always return affected elements, so everybody becomes able to see prev and next elements