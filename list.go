package steams

import (
	"iter"
	"slices"

	"github.com/javiorfo/nilo"
)

// It is a wrapper around iter.Seq[T] that provides a fluent API
// for functional-style iterator transformations.
type It[T any] iter.Seq[T]

// From creates a It from a variadic list of elements of type T.
func From[T any](args ...T) It[T] {
	return FromSlice(args)
}

// FromSlice creates a It (iterator) from a slice.
func FromSlice[T any](slice []T) It[T] {
	return It[T](slices.Values(slice))
}

// AsSeq returns the underlying iter.Seq[T].
func (it It[T]) AsSeq() iter.Seq[T] {
	return iter.Seq[T](it)
}

// Filter returns an iterator containing only elements that satisfy the predicate.
func (it It[T]) Filter(predicate func(T) bool) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Map returns an iterator that applies the mapper function to each element.
func (it It[T]) Map(mapper func(T) T) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// MapToString applies a mapper that transforms each element into a string.
func (it It[T]) MapToString(mapper func(T) string) It[string] {
	return func(yield func(string) bool) {
		for v := range it {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// MapToInt applies a mapper that transforms each element into an int.
func (it It[T]) MapToInt(mapper func(T) int) It[int] {
	return func(yield func(int) bool) {
		for v := range it {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// FilterMap applies a mapper that returns an Option. Only "Value" options
// are yielded, effectively filtering and transforming in one step.
func (it It[T]) FilterMap(mapper func(T) nilo.Option[T]) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

// FilterMapToString is a FilterMap variant that yields a sequence of strings.
func (it It[T]) FilterMapToString(mapper func(T) nilo.Option[string]) It[string] {
	return func(yield func(string) bool) {
		for v := range it {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

// FilterMapToInt is a FilterMap variant that yields a sequence of ints.
func (it It[T]) FilterMapToInt(mapper func(T) nilo.Option[int]) It[int] {
	return func(yield func(int) bool) {
		for v := range it {
			if m := mapper(v); m.IsValue() {
				if !yield(m.AsValue()) {
					return
				}
			}
		}
	}
}

// FlatMap applies a mapper that returns an iterator for each element,
// then flattens all resulting iterators into a single sequence.
func (it It[T]) FlatMap(mapper func(T) It[T]) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

// FlatMapToString is a FlatMap variant resulting in a sequence of strings.
func (it It[T]) FlatMapToString(mapper func(T) It[string]) It[string] {
	return func(yield func(string) bool) {
		for v := range it {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

// FlatMapToInt is a FlatMap variant resulting in a sequence of ints.
func (it It[T]) FlatMapToInt(mapper func(T) It[int]) It[int] {
	return func(yield func(int) bool) {
		for v := range it {
			for inner := range mapper(v) {
				if !yield(inner) {
					return
				}
			}
		}
	}
}

// Take returns an iterator that yields at most the first n elements.
func (it It[T]) Take(n int) It[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}

		count := 0
		for v := range it {
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

// Count consumes the iterator and returns the total number of elements.
// Note: This collects the iterator into memory to determine length.
func (it It[T]) Count() int {
	return len(it.Collect())
}

// ForEach executes the consumer function for every element in the iterator.
// This is a terminal operation.
func (it It[T]) ForEach(consumer func(T)) {
	for v := range it {
		consumer(v)
	}
}

// ForEachIdx executes the consumer function for every element,
// providing the current 0-based index. This is a terminal operation.
func (it It[T]) ForEachIdx(consumer func(int, T)) {
	index := 0
	for v := range it {
		consumer(index, v)
		index++
	}
}

// Inspect applies a function to each element without modifying the sequence.
// Useful for debugging or side effects during iteration.
func (it It[T]) Inspect(inspector func(T)) It[T] {
	for v := range it {
		inspector(v)
	}
	return it
}

// All returns true if every element satisfies the predicate.
// It short-circuits on the first false result.
func (it It[T]) All(predicate func(T) bool) bool {
	for v := range it {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element satisfies the predicate.
// It short-circuits on the first true result.
func (it It[T]) Any(predicate func(T) bool) bool {
	for v := range it {
		if predicate(v) {
			return true
		}
	}
	return false
}

// None returns true if no elements satisfy the predicate.
func (it It[T]) None(predicate func(T) bool) bool {
	return !it.Any(predicate)
}

// First returns the first element as an Option, or an empty Option
// if the iterator is empty.
func (it It[T]) First() nilo.Option[T] {
	for v := range it {
		return nilo.Value(v)
	}
	return nilo.Nil[T]()
}

// Find returns the first element that satisfies the predicate.
func (it It[T]) Find(predicate func(T) bool) nilo.Option[T] {
	for v := range it {
		if predicate(v) {
			return nilo.Value(v)
		}
	}
	return nilo.Nil[T]()
}

// TakeWhile yields elements as long as the predicate returns true.
func (it It[T]) TakeWhile(predicate func(T) bool) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// SkipWhile discards elements until the predicate returns false,
// then yields all remaining elements.
func (it It[T]) SkipWhile(predicate func(T) bool) It[T] {
	return func(yield func(T) bool) {
		dropping := true
		for v := range it {
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

// Fold reduces the sequence to a single value using an accumulator,
// starting with initValue and processing from left to right.
func (it It[T]) Fold(initValue T, acc func(T, T) T) T {
	result := initValue
	for v := range it {
		result = acc(result, v)
	}
	return result
}

// RFold reduces the sequence to a single value processing from right to left.
// Note: This collects the entire sequence into memory first.
func (it It[T]) RFold(initValue T, acc func(T, T) T) T {
	var elements []T = it.Collect()
	result := initValue
	for i := len(elements) - 1; i >= 0; i-- {
		result = acc(elements[i], result)
	}

	return result
}

// Reverse returns an iterator that yields elements in reverse order.
// Note: This collects the entire sequence into memory first.
func (it It[T]) Reverse() It[T] {
	return func(yield func(T) bool) {
		var buf []T = it.Collect()
		for index := len(buf) - 1; index >= 0; index-- {
			if !yield(buf[index]) {
				return
			}
		}
	}
}

// Position returns the index of the first element satisfying the predicate.
func (it It[T]) Position(predicate func(T) bool) nilo.Option[int] {
	index := 0
	for v := range it {
		if predicate(v) {
			return nilo.Value(index)
		}
		index++
	}
	return nilo.Nil[int]()
}

// RPosition returns the index of the last element satisfying the predicate.
// Note: This collects the entire sequence into memory first.
func (it It[T]) RPosition(predicate func(T) bool) nilo.Option[int] {
	list := it.Collect()
	length := len(list)
	for index := length - 1; index >= 0; index-- {
		if predicate(list[index]) {
			return nilo.Value(index)
		}
	}
	return nilo.Nil[int]()
}

// Enumerate returns a 2-variable iterator yielding (index, value) pairs.
func (it It[T]) Enumerate() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		index := 0
		for v := range it {
			if !yield(index, v) {
				return
			}
			index++
		}
	}
}

// Last returns the final element of the iterator.
// Note: This collects the entire sequence into memory first.
func (it It[T]) Last() nilo.Option[T] {
	slice := it.Collect()
	length := len(slice)
	if length > 0 {
		return nilo.Value(slice[length-1])
	}
	return nilo.Nil[T]()
}

// Skip returns an iterator that ignores the first n elements.
func (it It[T]) Skip(n int) It[T] {
	return func(yield func(T) bool) {
		count := 0
		for v := range it {
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

// SortBy returns an iterator yielding elements in the order defined
// by the comparison function. Note: This collects and sorts the
// entire sequence in memory.
func (it It[T]) SortBy(cmp func(T, T) int) It[T] {
	return func(yield func(T) bool) {
		buf := it.Collect()
		slices.SortFunc(buf, cmp)

		for _, v := range buf {
			if !yield(v) {
				return
			}
		}
	}
}

// Compare finds the "best" element based on the provided comparison function
// (e.g., to find Min or Max).
// Note: Use helper functions like steams.Min, steams.Max
func (it It[T]) Compare(cmp func(T, T) bool) nilo.Option[T] {
	count := it.Count()
	if count == 0 {
		return nilo.Nil[T]()
	}
	list := it.Collect()
	item := list[0]
	for i := 1; i < count; i++ {
		if cmp(list[i], item) {
			item = list[i]
		}
	}
	return nilo.Value(item)
}

// Collect consumes the iterator and returns a slice of all elements.
func (it It[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](it))
}

// Chain appends a second iterator to the current one.
func (it It[T]) Chain(i2 It[T]) It[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				return
			}
		}
		for v := range i2 {
			if !yield(v) {
				return
			}
		}
	}
}

// Nth returns the element at the given index, if it exists.
func (it It[T]) Nth(n int) nilo.Option[T] {
	if n < 0 {
		return nilo.Nil[T]()
	}

	count := 0
	for v := range it {
		if count == n {
			return nilo.Value(v)
		}
		count++
	}

	return nilo.Nil[T]()
}

// Partition splits the iterator into two collections: those that satisfy
// the predicate and those that do not.
func (it It[T]) Partition(politer func(T) bool) (It[T], It[T]) {
	var pos []T
	var neg []T
	for v := range it {
		if politer(v) {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
	}
	return FromSlice(pos), FromSlice(neg)
}
