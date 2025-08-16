package steams

import (
	"sort"

	"github.com/javiorfo/nilo"
)

// Map is a generic type that represents a collection of key-value pairs,
// where keys are of type K and values are of type V.
type Map[K comparable, V any] map[K]V

// Pair is a generic struct that holds a key-value pair.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Filter returns a new Map containing only the key-value pairs that match the provided predicate function.
func (m Map[K, V]) Filter(predicate func(K, V) bool) Steam2[K, V] {
	results := make(Map[K, V])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return results
}

// Map applies the provided mapper function to each key-value pair in the Map
// and returns a new Map with values of type V.
func (m Map[K, V]) Map(mapper func(K, V) V) Steam2[K, V] {
	results := make(Map[K, V], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

// MapToString applies the provided mapper function to each key-value pair in the Map
// and returns a new Map with values of type string.
func (m Map[K, V]) MapToString(mapper func(K, V) string) Steam2[K, string] {
	results := make(Map[K, string], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

// MapToInt applies the provided mapper function to each key-value pair in the Map
// and returns a new Map with values of type int.
func (m Map[K, V]) MapToInt(mapper func(K, V) int) Steam2[K, int] {
	results := make(Map[K, int], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

// FilterMap filters the key-value pairs based on the provided predicate
// and then maps the remaining pairs using the provided mapper function,
// returning a new Map with values of type V.
func (m Map[K, V]) FilterMap(predicate func(K, V) bool, mapper func(K, V) V) Steam2[K, V] {
	results := make(Map[K, V])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = mapper(k, v)
		}
	}
	return results
}

// FilterMapToInt filters the key-value pairs based on the provided predicate
// and then maps the remaining pairs using the provided mapper function,
// returning a new Map with values of type int.
func (m Map[K, V]) FilterMapToInt(predicate func(K, V) bool, mapper func(K, V) int) Steam2[K, int] {
	results := make(Map[K, int])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = mapper(k, v)
		}
	}
	return results
}

// FilterMapToString filters the key-value pairs based on the provided predicate
// and then maps the remaining pairs using the provided mapper function,
// returning a new Map with values of type string.
func (m Map[K, V]) FilterMapToString(predicate func(K, V) bool, mapper func(K, V) string) Steam2[K, string] {
	results := make(Map[K, string])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = mapper(k, v)
		}
	}
	return results
}

// ForEach applies the provided consumer function to each key-value pair in the Map.
func (m Map[K, V]) ForEach(consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}

// Peek applies the provided consumer function to each key-value pair in the Map
// without modifying it, and returns the original Map.
func (m Map[K, V]) Peek(consumer func(K, V)) Steam2[K, V] {
	for k, v := range m {
		consumer(k, v)
	}
	return m
}

// Limit restricts the number of key-value pairs in the Map to the specified limit
// and returns a new Map containing only the first 'limit' pairs.
func (m Map[K, V]) Limit(limit int) Steam2[K, V] {
	results := make(Map[K, V], 0)
	var counter int
	for k, v := range m {
		if counter >= limit {
			break
		}
		results[k] = v
		counter++
	}
	return results
}

// Count returns the number of key-value pairs in the Map.
func (m Map[K, V]) Count() int {
	return len(m)
}

// ValuesToSteam returns a Steam containing all the values from the Map.
func (m Map[K, V]) ValuesToSteam() Steam[V] {
	res := make(List[V], len(m))
	var index uint
	for _, v := range m {
		res[index] = v
		index++
	}
	return res
}

// KeysToSteam returns a Steam containing all the keys from the Map.
func (m Map[K, V]) KeysToSteam() Steam[K] {
	res := make(List[K], len(m))
	var index uint
	for k := range m {
		res[index] = k
		index++
	}
	return res
}

// ToAnySteam applies the provided mapper function to each key-value pair in the Map
// and returns a Steam containing the mapped values of type any.
func (m Map[K, V]) ToAnySteam(mapper func(K, V) any) Steam[any] {
	res := make(List[any], len(m))
	var index uint
	for k, v := range m {
		res[index] = mapper(k, v)
		index++
	}
	return res
}

// AllMatch checks if all key-value pairs in the Map match the provided predicate function.
// It returns true if all pairs match; otherwise, it returns false.
func (m Map[K, V]) AllMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// AnyMatch checks if any key-value pair in the Map matches the provided predicate function.
// It returns true if at least one pair matches; otherwise, it returns false.
func (m Map[K, V]) AnyMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// NoneMatch checks if no key-value pairs in the Map match the provided predicate function.
// It returns true if no pairs match; otherwise, it returns false.
func (m Map[K, V]) NoneMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return false
		}
	}
	return true
}

// Sorted returns a new Map containing the key-value pairs sorted according to the provided comparison function.
// The comparison function should define the order of the keys.
func (m Map[K, V]) Sorted(cmp func(K, K) bool) Steam2[K, V] {
	pairs := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair[K, V]{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return cmp(pairs[i].Key, pairs[j].Key)
	})

	results := make(Map[K, V], len(pairs))
	for _, pair := range pairs {
		results[pair.Key] = pair.Value
	}
	return results
}

// GetCompared returns an Option containing the key-value pair that is compared according to the provided comparison function.
// If the Map is empty, it returns an empty Option.
func (m Map[K, V]) GetCompared(cmp func(K, K) bool) nilo.Option[Pair[K, V]] {
	if len(m) == 0 {
		return nilo.None[Pair[K, V]]()
	}
	var item *Pair[K, V]
	for k, v := range m {
		if item == nil {
			item = &Pair[K, V]{Key: k, Value: v}
		} else if cmp(k, item.Key) {
			item.Key = k
			item.Value = v
		}
	}
	return nilo.Some(*item)
}

// Collect returns the underlying map of key-value pairs.
func (m Map[K, V]) Collect() map[K]V {
	return m
}
