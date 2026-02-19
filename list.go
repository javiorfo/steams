package steams

import (
	"slices"
	"sort"

	"github.com/javiorfo/nilo"
)

// List is a generic type that represents a slice of elements of type T.
// It provides methods to perform various operations on the list, following a functional programming style.
type List[T any] []T

// Of creates a Steam from a variadic list of elements of type T.
func Of[T any](args ...T) Steam[T] {
	return OfSlice(args)
}

// OfSlice creates a Steam from a slice of elements of type T.
func OfSlice[T any](slice []T) Steam[T] {
	return List[T](slice)
}

// Filter returns a new List containing only the elements that match the provided predicate function.
func (i List[T]) Filter(predicate func(T) bool) Steam[T] {
	index := 0
	for _, v := range i {
		if predicate(v) {
			i[index] = v
			index++
		}
	}
	return i[:index]
}

// Map applies the provided mapper function to each element in the List and returns the List modified of type T.
// If result to specific type is needed, use integration function Mapping[T, R](s Steam[T], mapper func(T) R)
func (list List[T]) Map(mapper func(T) T) Steam[T] {
	for i, v := range list {
		list[i] = mapper(v)
	}
	return list
}

// MapToString applies the provided mapper function to each element in the List and returns a new List of strings.
func (i List[T]) MapToString(mapper func(T) string) Steam[string] {
	results := make(List[string], i.Count())
	for i, v := range i {
		results[i] = mapper(v)
	}
	return results
}

// MapToInt applies the provided mapper function to each element in the List and returns a new List of integers.
func (i List[T]) MapToInt(mapper func(T) int) Steam[int] {
	results := make(List[int], i.Count())
	for i, v := range i {
		results[i] = mapper(v)
	}
	return results
}

// FilterMap filters the elements based on the provided predicate and then maps the remaining elements
// using the provided mapper function, returning the List modified of type T.
func (i List[T]) FilterMap(mapper func(T) nilo.Option[T]) Steam[T] {
	index := 0
	for _, v := range i {
		mapper(v).Consume(func(t T) {
			i[index] = t
			index++
		})
	}
	return i[:index]
}

// FilterMapToInt filters the elements based on the provided predicate and then maps the remaining elements
// using the provided mapper function, returning a new List of type int.
func (i List[T]) FilterMapToInt(predicate func(T) bool, mapper func(T) int) Steam[int] {
	var results List[int]
	for _, v := range i {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return results
}

// FilterMapToString filters the elements based on the provided predicate and then maps the remaining elements
// using the provided mapper function, returning a new List of type string.
func (i List[T]) FilterMapToString(predicate func(T) bool, mapper func(T) string) Steam[string] {
	var results List[string]
	for _, v := range i {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return results
}

// FlatMap applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type T.
func (i List[T]) FlatMap(mapper func(T) Steam[T]) Steam[T] {
	results := make(List[T], 0, i.Count())
	for _, v := range i {
		results = slices.Concat(results, mapper(v).(List[T]))
	}
	return results
}

// FlatMapToInt applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type int.
func (i List[T]) FlatMapToInt(mapper func(T) Steam[int]) Steam[int] {
	results := make(List[int], 0, i.Count())
	for _, v := range i {
		results = slices.Concat(results, mapper(v).(List[int]))
	}
	return results
}

// FlatMapToString applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type string.
func (i List[T]) FlatMapToString(mapper func(T) Steam[string]) Steam[string] {
	results := make(List[string], 0, i.Count())
	for _, v := range i {
		results = slices.Concat(results, mapper(v).(List[string]))
	}
	return results
}

// Limit restricts the number of elements in the List to the specified limit and returns a slice of List.
func (i List[T]) Limit(limit int) Steam[T] {
	count := i.Count()
	if limit > count {
		limit = count
	}
	return i[:limit]
}

// Count returns the number of elements in the List.
func (i List[T]) Count() int {
	return len(i)
}

// ForEach applies the provided consumer function to each element in the List.
func (i List[T]) ForEach(consumer func(T)) {
	for _, v := range i {
		consumer(v)
	}
}

// ForEachWithIndex applies the provided index and consumer function to each element in the List.
func (i List[T]) ForEachWithIndex(consumer func(int, T)) {
	for i, v := range i {
		consumer(i, v)
	}
}

// Peek applies the provided consumer function to each element in the List without modifying it,
// and returns the original List.
func (i List[T]) Peek(consumer func(T)) Steam[T] {
	for _, v := range i {
		consumer(v)
	}
	return i
}

// AllMatch checks if all elements in the List match the provided predicate function.
// It returns true if all elements match, false otherwise.
func (i List[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range i {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// AnyMatch checks if any element in the List matches the provided predicate function.
// It returns true if at least one element matches, false otherwise.
func (i List[T]) AnyMatch(predicate func(T) bool) bool {
	return slices.ContainsFunc(i, predicate)
}

// NoneMatch checks if no elements in the List match the provided predicate function.
// It returns true if no elements match, false otherwise.
func (i List[T]) NoneMatch(predicate func(T) bool) bool {
	return !slices.ContainsFunc(i, predicate)
}

// FindFirst returns an Option containing the first element of the List if it is present;
// otherwise, it returns an empty Option.
func (i List[T]) FindFirst() nilo.Option[T] {
	if i.Count() > 0 {
		return nilo.Value(i[0])
	}
	return nilo.Nil[T]()
}

// FindOne returns a nilo.Option[T] that match the given predicate function.
func (i List[T]) FindOne(predicate func(T) bool) nilo.Option[T] {
	for _, v := range i {
		if predicate(v) {
			return nilo.Value(v)
		}
	}
	return nilo.Nil[T]()
}

// TakeWhile returns a slice of List containing elements from the start of the List
// as long as they match the provided predicate function.
// It stops including elements as soon as an element does not match.
func (i List[T]) TakeWhile(predicate func(T) bool) Steam[T] {
	index := 0
	for _, v := range i {
		if predicate(v) {
			i[index] = v
			index++
		} else {
			break
		}
	}
	return i[:index]
}

// DropWhile returns a slice of List that skips elements from the start of the List
// as long as they match the provided predicate function.
// It includes all subsequent elements after the first non-matching element.
func (i List[T]) DropWhile(predicate func(T) bool) Steam[T] {
	index := 0
	for _, v := range i {
		if !predicate(v) {
			i[index] = v
			index++
		}
	}
	return i[:index]
}

// Reduce applies an accumulator function to the elements of the List,
// starting with the provided initial value. It returns the final accumulated value.
func (i List[T]) Reduce(initValue T, acc func(T, T) T) T {
	result := initValue
	for _, v := range i {
		result = acc(result, v)
	}
	return result
}

// Reverse returns the List in reverse order.
func (list List[T]) Reverse() Steam[T] {
	for i, j := 0, list.Count()-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// Position returns an Option containing the index of the first element that matches the provided predicate function;
// otherwise, it returns an empty Option.
func (i List[T]) Position(predicate func(T) bool) nilo.Option[int] {
	for i, v := range i {
		if predicate(v) {
			return nilo.Value(i)
		}
	}
	return nilo.Nil[int]()
}

// Last returns an Option containing the last element of the List if it is present;
// otherwise, it returns an empty Option.
func (i List[T]) Last() nilo.Option[T] {
	length := i.Count()
	if length > 0 {
		return nilo.Value(i[length-1])
	}
	return nilo.Nil[T]()
}

// Skip returns a slice of List that skips the first n elements of the original List.
// If n is greater than or equal to the length of the List, it returns an empty List.
func (i List[T]) Skip(n int) Steam[T] {
	if i.Count() <= n {
		return List[T]{}
	}

	return i[n:]
}

// Sorted returns a new List containing the elements of the original List sorted
// according to the provided comparison function.
func (i List[T]) Sorted(cmp func(T, T) bool) Steam[T] {
	slice := i.Collect()
	results := make(List[T], len(slice))
	copy(results, slice)
	sort.Slice(results, func(i, j int) bool {
		return cmp(results[i], results[j])
	})
	return results
}

// GetCompared returns an Option containing the element that is compared
// according to the provided comparison function. If the List is empty, it returns an empty Option.
func (list List[T]) GetCompared(cmp func(T, T) bool) nilo.Option[T] {
	count := list.Count()
	if count == 0 {
		return nilo.Nil[T]()
	}
	item := list[0]
	for i := 1; i < count; i++ {
		if cmp(list[i], item) {
			item = list[i]
		}
	}
	return nilo.Value(item)
}

// Collect returns the underlying slice of the List.
func (i List[T]) Collect() []T {
	return i
}
