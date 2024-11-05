package steams

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

func CollectSteamToSteam2[K comparable, V, T any](s Steam[T], keyFunc func(T) K, valueFunc func(T) V) Steam2[K, V] {
	m := make(map[K]V)
	for _, v := range s.Collect() {
		m[keyFunc(v)] = valueFunc(v)
	}
	return Map[K, V](m)
}

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

func GroupingBy[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, Steam[V]] {
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

func GroupingByCounting[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, int] {
	m := make(Map[K, int])
	for _, v := range s.Collect() {
		c := classifier(v)
        if _, ok := m[c]; ok {
			m[c] = m[c]+1
		} else {
			m[c] = 1
		}
	}
	return m
}

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

func Of[T any](args ...T) Steam[T] {
	return List[T](args)
}

func OfSlice[T any](slice []T) Steam[T] {
	return List[T](slice)
}

func OfMap[K comparable, V any](m map[K]V) Steam2[K, V] {
	return Map[K, V](m)
}
