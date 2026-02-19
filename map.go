package steams

import (
	"iter"
	"maps"
	"sort"

	"github.com/javiorfo/nilo"
)

// It2 is a wrapper around iter.Seq2[K, V] that provides a fluent API
// for functional-style transformations on key-value sequences.
type It2[K comparable, V any] iter.Seq2[K, V]

// Entry is a generic struct that holds a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// FromMap creates an It2 iterator from a standard Go map.
func FromMap[K comparable, V any](m map[K]V) It2[K, V] {
	return It2[K, V](maps.All(m))
}

// Filter returns an It2 iterator containing only the pairs
// that satisfy the predicate.
func (it It2[K, V]) Filter(predicate func(K, V) bool) It2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range it {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Map returns an It2 iterator that applies a transformation
// function to each key-value pair.
func (it It2[K, V]) Map(mapper func(K, V) (K, V)) It2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range it {
			k2, v2 := mapper(k, v)
			if !yield(k2, v2) {
				return
			}
		}
	}
}

// MapToString applies a mapper that transforms the value into a string
// while keeping or transforming the key.
func (it It2[K, V]) MapToString(mapper func(K, V) (K, string)) It2[K, string] {
	return func(yield func(K, string) bool) {
		for k, v := range it {
			k2, v2 := mapper(k, v)
			if !yield(k2, v2) {
				return
			}
		}
	}
}

// MapToInt applies a mapper that transforms the value into an int
// while keeping or transforming the key.
func (it It2[K, V]) MapToInt(mapper func(K, V) (K, int)) It2[K, int] {
	return func(yield func(K, int) bool) {
		for k, v := range it {
			k2, v2 := mapper(k, v)
			if !yield(k2, v2) {
				return
			}
		}
	}
}

// ForEach executes the consumer function for every key-value pair
// in the sequence. This is a terminal operation.
func (it It2[K, V]) ForEach(consumer func(K, V)) {
	for k, v := range it {
		consumer(k, v)
	}
}

// SortBy returns an iterator that yields pairs sorted based on the keys
// according to the provided comparison function.
// Note: This collects the entire sequence into memory before sorting.
func (it It2[K, V]) SortBy(cmp func(K, K) bool) It2[K, V] {
	return func(yield func(K, V) bool) {
		var entries []Entry[K, V]

		for k, v := range it.Collect() {
			entries = append(entries, Entry[K, V]{k, v})
		}

		sort.Slice(entries, func(i, j int) bool {
			return cmp(entries[i].Key, entries[j].Key)
		})

		for _, p := range entries {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// Inspect applies a function to each pair without modifying the sequence.
// Note: In its current implementation, this consumes the iterator immediately.
func (it It2[K, V]) Inspect(consumer func(K, V)) It2[K, V] {
	for k, v := range it {
		consumer(k, v)
	}
	return it
}

// Limit returns a new iterator that yields at most 'n' elements.
func (it It2[K, V]) Take(n int) It2[K, V] {
	return func(yield func(K, V) bool) {
		if n <= 0 {
			return
		}

		count := 0
		for k, v := range it {
			if !yield(k, v) {
				return
			}
			count++
			if count >= n {
				return
			}
		}
	}
}

// Values returns a single-value iterator for the values.
func (it It2[K, V]) Values() It[V] {
	return func(yield func(V) bool) {
		for _, v := range it {
			if !yield(v) {
				return
			}
		}
	}
}

// Keys returns a single-value iterator for the keys.
func (it It2[K, V]) Keys() It[K] {
	return func(yield func(K) bool) {
		for k := range it {
			if !yield(k) {
				return
			}
		}
	}
}

// All returns true if every element in the iterator satisfies the predicate.
func (it It2[K, V]) All(predicate func(K, V) bool) bool {
	for k, v := range it {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the iterator satisfies the predicate.
func (it It2[K, V]) Any(predicate func(K, V) bool) bool {
	for k, v := range it {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// None returns true if no elements in the iterator satisfy the predicate.
func (it It2[K, V]) None(predicate func(K, V) bool) bool {
	for k, v := range it {
		if predicate(k, v) {
			return false
		}
	}
	return true
}

// Compare returns the "best" Entry based on the comparison function.
// If the iterator is empty, it returns a Nil option.
func (it It2[K, V]) Compare(cmp func(K, K) bool) nilo.Option[Entry[K, V]] {
	var result Entry[K, V]
	found := false

	for k, v := range it {
		if !found {
			result = Entry[K, V]{Key: k, Value: v}
			found = true
			continue
		}

		if cmp(k, result.Key) {
			result.Key = k
			result.Value = v
		}
	}

	if !found {
		return nilo.Nil[Entry[K, V]]()
	}
	return nilo.Value(result)
}

// Collect consumes the iterator and returns a map of all key-value pairs.
func (it It2[K, V]) Collect() map[K]V {
	return maps.Collect(iter.Seq2[K, V](it))
}

// Count consumes the iterator and returns the total number of pairs.
// Note: This implementation collects the iterator into a map to determine length.
func (it It2[K, V]) Count() int {
	return len(it.Collect())
}
