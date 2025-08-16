package steams

import "slices"

// Distinct returns a new Steam containing only the unique elements from the input Steam.
// It uses a map to track seen elements and filters out duplicates.
func Distinct[T comparable](s Steam[T]) Steam[T] {
	m := make(map[T]bool)
	slice := s.Collect()
	results := make(List[T], 0, s.Count())
	for _, v := range slice {
		if !m[v] {
			m[v] = true
			results = append(results, v)
		}
	}
	return results
}

// Mapper applies the provided mapper function to each element in the List and returns a new List of type R.
func Mapper[T, R any](s Steam[T], mapper func(T) R) Steam[R] {
	results := make(List[R], s.Count())
	for i, v := range s.Collect() {
		results[i] = mapper(v)
	}
	return results
}

// FlatMapper applies the provided mapper function to each element in the List and returns a new List of type R.
func FlatMapper[T, R any](s Steam[T], mapper func(T) Steam[R]) Steam[R] {
	results := make(List[R], 0, s.Count())
	for _, v := range s.Collect() {
		results = slices.Concat(results, mapper(v).(List[R]))
	}
	return results
}

// CollectSteamToSteam2 transforms a Steam of type T into a Steam2 of key-value pairs,
// where keys and values are derived from the provided keyFunc and valueFunc.
// It collects elements from the input Steam and maps them to a new Steam2 instance.
func CollectSteamToSteam2[K comparable, V, T any](s Steam[T], keyFunc func(T) K, valueFunc func(T) V) Steam2[K, V] {
	m := make(map[K]V)
	for _, v := range s.Collect() {
		m[keyFunc(v)] = valueFunc(v)
	}
	return Map[K, V](m)
}

// CollectSteam2ToSteam transforms a Steam2 of key-value pairs into a Steam of type R,
// using the provided mapper function to convert each key-value pair into a single value of type R.
func CollectSteam2ToSteam[K comparable, V, R any](s Steam2[K, V], mapper func(K, V) R) Steam[R] {
	m := s.Collect()
	results := make([]R, len(m))
	var index int
	for k, v := range m {
		results[index] = mapper(k, v)
		index++
	}
	return List[R](results)
}

// GroupBy groups elements of a Steam by a classifier function that maps each element to a key.
// It returns a Steam2 where each key corresponds to a Steam of elements that share that key.
func GroupBy[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, Steam[V]] {
	m := make(Map[K, Steam[V]])
	for _, v := range s.Collect() {
		c := classifier(v)
		if _, ok := m[c]; ok {
			m[c] = append(m[c].(List[V]), v)
		} else {
			m[c] = append(List[V]{}, v)
		}
	}
	return m
}

// GroupByCounting groups elements of a Steam by a classifier function and counts the occurrences of each key.
// It returns a Steam2 where each key corresponds to the count of elements that share that key.
func GroupByCounting[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, int] {
	m := make(Map[K, int])
	for _, v := range s.Collect() {
		c := classifier(v)
		if _, ok := m[c]; ok {
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}
	return m
}

// Zip combines two Steams into a single Steam of structs, where each struct contains one element from each input Steam.
// It panics if the two Steams do not have the same length.
func Zip[T, R any](s1 Steam[T], s2 Steam[R]) Steam[struct {
	first  T
	second R
}] {
	slice1 := s1.Collect()
	slice2 := s2.Collect()
	if len(slice1) != len(slice2) {
		panic("Steams must have the same length")
	}

	result := make(List[struct {
		first  T
		second R
	}], len(slice1))
	for i := range slice1 {
		result[i] = struct {
			first  T
			second R
		}{slice1[i], slice2[i]}
	}
	return result
}

// Of creates a Steam from a variadic list of elements of type T.
func Of[T any](args ...T) Steam[T] {
	return List[T](args)
}

// OfSlice creates a Steam from a slice of elements of type T.
func OfSlice[T any](slice []T) Steam[T] {
	return Of(slice...)
}

// OfMap creates a Steam2 from a map of key-value pairs.
// The keys and values are derived from the provided map.
func OfMap[K comparable, V any](m map[K]V) Steam2[K, V] {
	return Map[K, V](m)
}
