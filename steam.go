package steams

// Steam[T] is an interface for a collection of elements of type T,
// providing various methods for functional-style processing.
type Steam[T any] interface {
	// Filter returns a new Steam containing only the elements that match the
	// given predicate function.
	Filter(predicate func(T) bool) Steam[T]

	// MapToAny transforms each element of the Steam using the provided mapper
	// function and returns a new Steam of type any.
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
	// along with a boolean indicating if such an element exists.
	// This can be used as an implementation of Max or Min function
	GetCompared(cmp func(T, T) bool) (*T, bool)

	// FindFirst returns the first element in the Steam along with a boolean
	// indicating if an element exists.
	FindFirst() (*T, bool)

	// Last returns the last element in the Steam along with a boolean indicating
	// if an element exists.
	Last() (*T, bool)

	// Position returns the index of the first element that matches the predicate
	// along with a boolean indicating if such an element exists.
	Position(predicate func(T) bool) (*int, bool)

	// Skip returns a new Steam that skips the first 'n' elements.
	Skip(n int) Steam[T]

	// Count returns the number of elements in the Steam.
	Count() int

	// Collect returns a slice containing all elements in the Steam.
	Collect() []T
}

// Steam2[K, V] is an interface for a map of elements of type K and V,
// providing various methods for functional-style processing.
type Steam2[K comparable, V any] interface {
	Filter(predicate func(K, V) bool) Steam2[K, V]
	MapToAny(mapper func(K, V) any) Steam2[K, any]
	MapToInt(mapper func(K, V) int) Steam2[K, int]
	MapToString(mapper func(K, V) string) Steam2[K, string]
	FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any]
	ForEach(consumer func(K, V))
	Peek(consumer func(K, V)) Steam2[K, V]
	Limit(limit int) Steam2[K, V]
	AllMatch(predicate func(K, V) bool) bool
	AnyMatch(predicate func(K, V) bool) bool
	NoneMatch(predicate func(K, V) bool) bool
	Sorted(cmp func(K, K) bool) Steam2[K, V]
	GetCompared(cmp func(K, K) bool) (*Pair[K, V], bool)
	Count() int
	Collect() map[K]V
	KeysToSteam() Steam[K]
	ValuesToSteam() Steam[V]
	ToAnySteam(mapper func(K, V) any) Steam[any]
}
