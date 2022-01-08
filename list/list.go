package list

type list[T any] struct {
	head, tail *element[T]
	len        int
}

func New[T any]() *list[T] {
	return new(list[T])
}

// add adds value v to the end of list and returns added element
func (l *list[T]) add(v T) (el *element[T]) {
	el = &element[T]{v: v}

	if l.head == nil {
		l.head, l.tail = el, el
	} else {
		l.tail.next = el
		el.prev = l.tail
		l.tail = el
	}

	l.len++
	return el
}

// first finds the first element in the list that matches a predicate p
// and returns its index and found element. If there is no matched element
// the method returns -1, nil
func (l *list[T]) first(p func(v T) bool) (idx int, el *element[T]) {
	for el = l.head; el != nil; el = el.Next() {
		if p(el.v) {
			return idx, el
		}
		idx++ 
	}
	return -1, nil
}

// find finds all elements in the list that matches a predicate p
// and returns it as a slice of elements
func (l *list[T]) find(p func(v T) bool) (els []*element[T]) {
	els = make([]*element[T], 0)
	for el := l.head; el != nil; el = el.Next() {
		if p(el.v) {
			els = append(els, el)
		}
	}
	return els
}

// findAtIdx finds element at index idx and returns it. If idx < 0 or idx >= length of the list returns nil
func (l *list[T]) findAtIdx(idx int) (el *element[T]) {
	if idx < 0 || idx >= l.len {
		return nil
	}

	// TODO: there must be a way to start searching from tail if idx > l.len/2

	el = l.head
	
	for i := 0; i < idx; i++ {
		el = el.Next()
	}
	return el
}

// remove removes the first element that matches a predicate p
// and returns it as an element
func (l *list[T]) remove(p func(v T) bool) *element[T] {
	if _, el := l.first(p); el != nil {
		return el.removeFrom(l)
	}
	return nil
}

// removeMany removes the elements that matches a predicate p
// and returns removed elements
func (l *list[T]) removeMany(p func(v T) bool) (els []*element[T]) {
	els = l.find(p)
	for i, e := range els {
		els[i] = e.removeFrom(l)
	}
	return els
}

// clear clears the list
func (l *list[T]) clear() {
	l = New[T]()
}

// insertAt inserts value v in the list at the given index idx.
// Returns inserted element or nil if there is no element at the index
func (l *list[T]) insertAt(idx int, v T) (el *element[T]) {
	cur := l.findAtIdx(idx)
	if cur == nil {
		return
	}

	el = &element[T]{v: v}

	if cur == l.head {
		l.head.prev = el
		l.head = el
	} else {
		cur.Prev().next = el
		cur.prev = el
	}
	l.len++
	return el
}



// TODO: implement methods: insertAt(idx int, v T), removeAt(idx int), forEach(func(v T)), map(???) (same as js Array.prototype.map())
// TODO: write descriptions
// TODO: try to implement some specific extension methods, e.g. Sum(), Min(), Max() fot list[int], if it's allowed
// TODO: always return affected elements, so everybody becomes able to see prev and next elements