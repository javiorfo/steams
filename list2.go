package steams

import (
	"iter"
	"slices"

	"github.com/javiorfo/nilo"
)

type it[T any] iter.Seq[T]

// Of creates a Steam from a variadic list of elements of type T.
func Of2[T any](args ...T) it[T] {
	return OfSlice2(args)
}

// OfSlice creates a List (iterator) from a slice.
func OfSlice2[T any](slice []T) it[T] {
	return it[T](slices.Values(slice))
}

func (i it[T]) Filter(predicate func(T) bool) it[T] {
	return func(yield func(T) bool) {
		for v := range i {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (i it[T]) Map(mapper func(T) T) it[T] {
	return func(yield func(T) bool) {
		for v := range i {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

func (i it[T]) MapToString(mapper func(T) string) it[string] {
	return func(yield func(string) bool) {
		for v := range i {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

func (i it[T]) MapToInt(mapper func(T) int) it[int] {
	return func(yield func(int) bool) {
		for v := range i {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

func (i it[T]) FilterMap(mapper func(T) nilo.Option[T]) it[T] {
	return func(yield func(T) bool) {
		for v := range i {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

func (i it[T]) FilterMapToString(mapper func(T) nilo.Option[string]) it[string] {
	return func(yield func(string) bool) {
		for v := range i {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

func (i it[T]) FilterMapToInt(mapper func(T) nilo.Option[int]) it[int] {
	return func(yield func(int) bool) {
		for v := range i {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

func (i it[T]) FlatMap(mapper func(T) it[T]) it[T] {
	return func(yield func(T) bool) {
		for v := range i {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

func (i it[T]) FlatMapToString(mapper func(T) it[string]) it[string] {
	return func(yield func(string) bool) {
		for v := range i {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

func (i it[T]) FlatMapToInt(mapper func(T) it[int]) it[int] {
	return func(yield func(int) bool) {
		for v := range i {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

func (i it[T]) Take(n int) it[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}

		count := 0
		for v := range i {
			if !yield(v) {
				return
			}
			count++
			if count >= n {
				break
			}
		}
	}
}

func (i it[T]) Count() int {
	return len(i.Collect())
}

func (i it[T]) ForEach(consumer func(T)) {
	for v := range i {
		consumer(v)
	}
}

func (i it[T]) ForEachIdx(consumer func(int, T)) {
	index := 0
	for v := range i {
		consumer(index, v)
		index++
	}
}

func (i it[T]) Inspect(inspector func(T)) it[T] {
	for v := range i {
		inspector(v)
	}
	return i
}

func (i it[T]) All(predicate func(T) bool) bool {
	for v := range i {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (i it[T]) Any(predicate func(T) bool) bool {
	for v := range i {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (i it[T]) None(predicate func(T) bool) bool {
	return !i.Any(predicate)
}

func (i it[T]) First() nilo.Option[T] {
	for v := range i {
		return nilo.Value(v)
	}
	return nilo.Nil[T]()
}

func (i it[T]) Find(predicate func(T) bool) nilo.Option[T] {
	for v := range i {
		if predicate(v) {
			return nilo.Value(v)
		}
	}
	return nilo.Nil[T]()
}

func (i it[T]) TakeWhile(predicate func(T) bool) it[T] {
	return func(yield func(T) bool) {
		for v := range i {
			if !predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

func (i it[T]) SkipWhile(predicate func(T) bool) it[T] {
	return func(yield func(T) bool) {
		dropping := true
		for v := range i {
			if dropping {
				if predicate(v) {
					continue
				}
				dropping = false
			}

			if !yield(v) {
				return
			}
		}
	}
}

func (i it[T]) Fold(initValue T, acc func(T, T) T) T {
	result := initValue
	for v := range i {
		result = acc(result, v)
	}
	return result
}

func (i it[T]) RFold(initValue T, acc func(T, T) T) T {
	var elements []T = i.Collect()
	result := initValue
	for i := len(elements) - 1; i >= 0; i-- {
		result = acc(elements[i], result)
	}

	return result
}

func (i it[T]) Reverse() it[T] {
	return func(yield func(T) bool) {
		var buf []T = i.Collect()
		for index := len(buf) - 1; index >= 0; index-- {
			if !yield(buf[index]) {
				return
			}
		}
	}
}

func (i it[T]) Position(predicate func(T) bool) nilo.Option[int] {
	index := 0
	for v := range i {
		if predicate(v) {
			return nilo.Value(index)
		}
		index++
	}
	return nilo.Nil[int]()
}

func (i it[T]) RPosition(predicate func(T) bool) nilo.Option[int] {
	list := i.Collect()
	length := len(list)
	for index := length - 1; index >= 0; index-- {
		if predicate(list[index]) {
			return nilo.Value(index)
		}
	}
	return nilo.Nil[int]()
}

func (i it[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for v := range i {
			if !yield(index, v) {
				return
			}
			index++
		}
	}
}

func (i it[T]) Last() nilo.Option[T] {
	slice := i.Collect()
	length := len(slice)
	if length > 0 {
		return nilo.Value(slice[length-1])
	}
	return nilo.Nil[T]()
}

func (i it[T]) Skip(n int) it[T] {
	return func(yield func(T) bool) {
		count := 0
		for v := range i {
			if count < n {
				count++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func (i it[T]) Sorted(cmp func(T, T) int) it[T] {
	return func(yield func(T) bool) {
		buf := i.Collect()
		slices.SortFunc(buf, cmp)

		for _, v := range buf {
			if !yield(v) {
				return
			}
		}
	}
}

func (i it[T]) Compare(cmp func(T, T) bool) nilo.Option[T] {
	count := i.Count()
	if count == 0 {
		return nilo.Nil[T]()
	}
	list := i.Collect()
	item := list[0]
	for i := 1; i < count; i++ {
		if cmp(list[i], item) {
			item = list[i]
		}
	}
	return nilo.Value(item)
}

func (i it[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](i))
}
