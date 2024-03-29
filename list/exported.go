package list

func (l *list[T]) Len() int {
	return l.len
}

func (l *list[T]) Add(v T) *element[T] {
	return l.add(v)
}

func (l *list[T]) AddMany(s ...T) []*element[T] {
	rv := make([]*element[T], len(s))
	for i, v := range s {
		rv[i] = l.add(v)
	}
	return rv
}

func (l *list[T]) Remove(p func(v T) bool) bool {
	return l.remove(p)
}

func (l *list[T]) RemoveMany(p func(v T) bool) []T {
	return l.removeMany(p)
}

func (l *list[T]) First(p func(v T) bool) *element[T] {
	return l.first(p)
}

func (l *list[T]) ForEach(f func(v T)) {
	for e := l.head; e != nil; e = e.Next() {
		f(e.v)
	}
}
