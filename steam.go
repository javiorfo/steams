package steams

import "github.com/javiorfo/steams/opt"

// Steam[T] is an interface for a collection of elements of type T,
// providing various methods for functional-style processing.
type Steam[T any] interface {
	// Filter returns a new Steam containing only the elements that match the
	// given predicate function.
	Filter(predicate func(T) bool) Steam[T]

	// MapToAny transforms each element of the Steam using the provided mapper
	// function and returns a new Steam of type any.
	// If result to specific type is needed, use integration function Mapping[T, R](s Steam[T], mapper func(T) R)
	MapToAny(mapper func(T) any) Steam[any]

	// MapToInt transforms each element of the Steam using the provided mapper
	// function and returns a new Steam of type int.
	MapToInt(mapper func(T) int) Steam[int]

	// MapToString transforms each element of the Steam using the provided mapper
	// function and returns a new Steam of type string.
	MapToString(mapper func(T) string) Steam[string]

	// FilterMapToAny combines filtering and mapping. Returns a new Steam containing
	// elements that match the predicate and are transformed by the mapper.
	FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any]

	// FlatMapToAny transforms each element of the Steam into a new Steam using
	// the provided mapper function and flattens the result into a single Steam.
	FlatMapToAny(mapper func(T) Steam[any]) Steam[any]

	// ForEach executes the provided consumer function for each element in the Steam.
	ForEach(consumer func(T))

	// Peek allows for inspecting each element without consuming it. It returns
	// a new Steam that allows for further operations after consuming the elements.
	Peek(consumer func(T)) Steam[T]

	// Limit returns a new Steam containing only the first 'limit' elements.
	Limit(limit int) Steam[T]

	// AllMatch returns true if all elements match the given predicate.
	AllMatch(predicate func(T) bool) bool

	// AnyMatch returns true if any element matches the given predicate.
	AnyMatch(predicate func(T) bool) bool

	// NoneMatch returns true if no elements match the given predicate.
	NoneMatch(predicate func(T) bool) bool

	// TakeWhile returns a new Steam containing elements as long as they match
	// the given predicate.
	TakeWhile(predicate func(T) bool) Steam[T]

	// DropWhile returns a new Steam that skips elements as long as they match
	// the given predicate.
	DropWhile(predicate func(T) bool) Steam[T]

	// Reduce combines elements of the Steam into a single value using the provided
	// accumulator function, starting with the given initial value.
	Reduce(initValue T, acc func(T, T) T) T

	// Reverse returns a new Steam with the elements in reverse order.
	Reverse() Steam[T]

	// Sorted returns a new Steam with the elements sorted according to the
	// provided comparison function.
	Sorted(cmp func(T, T) bool) Steam[T]

	// GetCompared returns the first element that matches the comparison function
	// wrapped in an opt.Optional[T].
	// This can be used as an implementation of Max or Min function.
	GetCompared(cmp func(T, T) bool) opt.Optional[T]

	// FindFirst returns the first element in the Steam as an opt.Optional[T]
	// indicating if an element exists.
	FindFirst() opt.Optional[T]

	// Last returns an opt.Optional of the last element in the Steam
	// indicating if an element exists.
	Last() opt.Optional[T]

	// Position returns the index of the first element that matches the predicate
	// wrapped in an opt.Optional[int].
	Position(predicate func(T) bool) opt.Optional[int]

	// Skip returns a new Steam that skips the first 'n' elements.
	Skip(n int) Steam[T]

	// Count returns the number of elements in the Steam.
	Count() int

	// Collect returns a slice containing all elements in the Steam.
	Collect() []T
}

// Steam2[K, V] is an interface for a map of elements of type K and V,
// providing various methods for functional-style processing.
// Steam2 is a generic interface that provides a fluent API for processing key-value pairs.
type Steam2[K comparable, V any] interface {

	// Filter returns a new Steam2 instance containing only the elements that match the given predicate.
	// The predicate function takes a key and a value and returns true if the element should be included.
	Filter(predicate func(K, V) bool) Steam2[K, V]

	// MapToAny transforms the elements of the Steam2 instance using the provided mapper function,
	// which takes a key and a value and returns a value of any type.
	// The result is a new Steam2 instance with the transformed values.
	MapToAny(mapper func(K, V) any) Steam2[K, any]

	// MapToInt transforms the elements of the Steam2 instance using the provided mapper function,
	// which takes a key and a value and returns an int.
	// The result is a new Steam2 instance with the transformed integer values.
	MapToInt(mapper func(K, V) int) Steam2[K, int]

	// MapToString transforms the elements of the Steam2 instance using the provided mapper function,
	// which takes a key and a value and returns a string.
	// The result is a new Steam2 instance with the transformed string values.
	MapToString(mapper func(K, V) string) Steam2[K, string]

	// FilterMapToAny applies a predicate to filter elements and then maps the remaining elements
	// using the provided mapper function. The result is a new Steam2 instance with the transformed values.
	FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any]

	// ForEach applies the provided consumer function to each element in the Steam2 instance.
	// The consumer function takes a key and a value and performs an action for each element.
	ForEach(consumer func(K, V))

	// Peek allows you to inspect each element in the Steam2 instance using the provided consumer function
	// without modifying the stream. The result is a new Steam2 instance with the same elements.
	Peek(consumer func(K, V)) Steam2[K, V]

	// Limit restricts the number of elements in the Steam2 instance to the specified limit.
	// The result is a new Steam2 instance containing only the first 'limit' elements.
	Limit(limit int) Steam2[K, V]

	// AllMatch checks if all elements in the Steam2 instance match the given predicate.
	// The predicate function takes a key and a value and returns true if the element matches.
	// It returns true if the stream is empty.
	AllMatch(predicate func(K, V) bool) bool

	// AnyMatch checks if any element in the Steam2 instance matches the given predicate.
	// The predicate function takes a key and a value and returns true if the element matches.
	// It returns false if the stream is empty.
	AnyMatch(predicate func(K, V) bool) bool

	// NoneMatch checks if no elements in the Steam2 instance match the given predicate.
	// The predicate function takes a key and a value and returns true if the element matches.
	// It returns true if the stream is empty.
	NoneMatch(predicate func(K, V) bool) bool

	// Sorted returns a new Steam2 instance with elements sorted according to the provided comparison function.
	// The comparison function takes two keys and returns true if the first key is less than the second.
	Sorted(cmp func(K, K) bool) Steam2[K, V]

	// GetCompared returns an optional Pair containing the first two elements of the stream compared
	// using the provided comparison function. If the stream has fewer than two elements, it returns an empty optional.
	GetCompared(cmp func(K, K) bool) opt.Optional[Pair[K, V]]

	// Count returns the number of elements in the Steam2 instance.
	Count() int

	// Collect gathers all elements in the Steam2 instance into a map with keys of type K and values of type V.
	Collect() map[K]V

	// KeysToSteam returns a new Steam instance containing only the keys from the Steam2 instance.
	KeysToSteam() Steam[K]

	// ValuesToSteam returns a new Steam instance containing only the values from the Steam2 instance.
	ValuesToSteam() Steam[V]

	// ToAnySteam transforms the elements of the Steam2 instance using the provided mapper function,
	// which takes a key and a value and returns a value of any type.
	// The result is a new Steam instance with the transformed values.
	ToAnySteam(mapper func(K, V) any) Steam[any]
}
