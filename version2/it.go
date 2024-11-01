package version2

type List[T any] struct {
	data  []T
	index int
}
type Mapa[T comparable, R any] map[T]R

type Steam[T any] interface {
	Next() T // Posible optional
	HasNext() bool
	// con dos generics GetKey => index array o key map
}

func (l List[T]) Next() T {
	value := l.data[l.index]
	l.index += 1
	return value
}

func (l List[T]) HasNext() bool {
	return l.index < len(l.data)
}

func (l List[T]) Filter(f func(T) bool) Steam[T] {
	results := make([]T, 0)
	for l.HasNext() {
		value := l.Next()
		if f(value) {
			results = append(results, value)
		}
	}
	return List[T]{data: results}
}
