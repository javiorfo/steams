package steams

import (
	"slices"
	"sort"

	"github.com/javiorfo/steams/opt"
)

// List is a generic type that represents a slice of elements of type T.
// It provides methods to perform various operations on the list, following a functional programming style.
type List[T any] []T

// ListOf creates a List from a variadic list of elements of type T and returns it as a Steam.
func ListOf[T any](args ...T) Steam[T] {
	return List[T](args)
}

// Filter returns a new List containing only the elements that match the provided predicate function.
func (list List[T]) Filter(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

// MapToAny applies the provided mapper function to each element in the List and returns a new List of type any.
// If result to specific type is needed, use integration function Mapping[T, R](s Steam[T], mapper func(T) R)
func (list List[T]) MapToAny(mapper func(T) any) Steam[any] {
	results := make(List[any], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

// MapToString applies the provided mapper function to each element in the List and returns a new List of strings.
func (list List[T]) MapToString(mapper func(T) string) Steam[string] {
	results := make(List[string], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

// MapToInt applies the provided mapper function to each element in the List and returns a new List of integers.
func (list List[T]) MapToInt(mapper func(T) int) Steam[int] {
	results := make(List[int], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

// FilterMapToAny filters the elements based on the provided predicate and then maps the remaining elements
// using the provided mapper function, returning a new List of type any.
func (list List[T]) FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any] {
	results := make(List[any], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return results
}

// FlatMapToAny applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type any.
func (list List[T]) FlatMapToAny(mapper func(T) Steam[any]) Steam[any] {
	results := make(List[any], 0, list.Count())
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[any]))
	}
	return results
}

// FlatMapToInt applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type int.
func (list List[T]) FlatMapToInt(mapper func(T) Steam[int]) Steam[int] {
	results := make(List[int], 0, list.Count())
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[int]))
	}
	return results
}

// FlatMapToString applies the provided mapper function to each element in the List, which returns a Steam,
// and concatenates the results into a single List of type string.
func (list List[T]) FlatMapToString(mapper func(T) Steam[string]) Steam[string] {
	results := make(List[string], 0, list.Count())
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[string]))
	}
	return results
}

// Limit restricts the number of elements in the List to the specified limit and returns a new List.
func (list List[T]) Limit(limit int) Steam[T] {
	if limit > list.Count() {
		limit = list.Count()
	}
	results := make(List[T], limit)
	copy(results, list[:limit])
	return results
}

// Count returns the number of elements in the List.
func (list List[T]) Count() int {
	return len(list)
}

// ForEach applies the provided consumer function to each element in the List.
func (list List[T]) ForEach(consumer func(T)) {
	for _, v := range list {
		consumer(v)
	}
}

// Peek applies the provided consumer function to each element in the List without modifying it,
// and returns the original List.
func (list List[T]) Peek(consumer func(T)) Steam[T] {
	for _, v := range list {
		consumer(v)
	}
	return list
}

// AllMatch checks if all elements in the List match the provided predicate function.
// It returns true if all elements match, false otherwise.
func (list List[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// AnyMatch checks if any element in the List matches the provided predicate function.
// It returns true if at least one element matches, false otherwise.
func (list List[T]) AnyMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return true
		}
	}
	return false
}

// NoneMatch checks if no elements in the List match the provided predicate function.
// It returns true if no elements match, false otherwise.
func (list List[T]) NoneMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return false
		}
	}
	return true
}

// FindFirst returns an Optional containing the first element of the List if it is present;
// otherwise, it returns an empty Optional.
func (list List[T]) FindFirst() opt.Optional[T] {
	if len(list) > 0 {
		return opt.Of(list[0])
	}
	return opt.Empty[T]()
}

// TakeWhile returns a new List containing elements from the start of the List
// as long as they match the provided predicate function.
// It stops including elements as soon as an element does not match.
func (list List[T]) TakeWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		} else {
			break
		}
	}
	return results
}

// DropWhile returns a new List that skips elements from the start of the List
// as long as they match the provided predicate function.
// It includes all subsequent elements after the first non-matching element.
func (list List[T]) DropWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if !predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

// Reduce applies an accumulator function to the elements of the List,
// starting with the provided initial value. It returns the final accumulated value.
func (list List[T]) Reduce(initValue T, acc func(T, T) T) T {
	result := initValue
	for _, v := range list {
		result = acc(result, v)
	}
	return result
}

// Reverse returns a new List containing the elements of the original List in reverse order.
func (list List[T]) Reverse() Steam[T] {
	length := len(list)
	results := make(List[T], length)
	index := length - 1
	for i := range list {
		results[i] = list[index]
		index--
	}
	return results
}

// Position returns an Optional containing the index of the first element that matches the provided predicate function;
// otherwise, it returns an empty Optional.
func (list List[T]) Position(predicate func(T) bool) opt.Optional[int] {
	for i, v := range list {
		if predicate(v) {
			return opt.Of(i)
		}
	}
	return opt.Empty[int]()
}

// Last returns an Optional containing the last element of the List if it is present;
// otherwise, it returns an empty Optional.
func (list List[T]) Last() opt.Optional[T] {
	length := list.Count()
	if length > 0 {
		return opt.Of(list[length-1])
	}
	return opt.Empty[T]()
}

// Skip returns a new List that skips the first n elements of the original List.
// If n is greater than or equal to the length of the List, it returns an empty List.
func (list List[T]) Skip(n int) Steam[T] {
	length := len(list)
	if length > n {
		length = length - n
	} else {
		return List[T]{}
	}

	results := make(List[T], length)
	for i := 0; i < length; i++ {
		results[i] = list[i+n]
	}
	return results
}

// Sorted returns a new List containing the elements of the original List sorted
// according to the provided comparison function.
func (list List[T]) Sorted(cmp func(T, T) bool) Steam[T] {
	slice := list.Collect()
	results := make(List[T], len(slice))
	copy(results, slice)
	sort.Slice(results, func(i, j int) bool {
		return cmp(results[i], results[j])
	})
	return results
}

// GetCompared returns an Optional containing the element that is compared
// according to the provided comparison function. If the List is empty, it returns an empty Optional.
func (list List[T]) GetCompared(cmp func(T, T) bool) opt.Optional[T] {
	if len(list) == 0 {
		return opt.Empty[T]()
	}
	item := list[0]
	for i := 1; i < len(list); i++ {
		if cmp(list[i], item) {
			item = list[i]
		}
	}
	return opt.Of(item)
}

// Collect returns the underlying slice of the List.
func (list List[T]) Collect() []T {
	return list
}
