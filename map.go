package steams

type Map[K comparable, V any] map[K]V

/* func (m Map[K, V]) Filter(predicate func(V) bool) Steam[V] {
	results := make(map[K]V)
	for k, v := range m {
		if predicate(v) {
			results[k] = v
		}
	}
	return Map[K, V](results)
} */

func (m Map[K, V]) CollectValuesToSteam() Steam[V] {
    res := make([]V, len(m))
    var index uint
    for _, v := range m {
        res[index] = v
        index++
    }
	return List[V](res)
}

func (m Map[K, V]) CollectKeysToSteam() Steam[K] {
    res := make([]K, len(m))
    var index uint
    for k := range m {
        res[index] = k
        index++
    }
	return List[K](res)
}

func (m Map[K, V]) Filter(predicate func(K, V) bool) Steam2[K, V] {
    results := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return Map[K, V](results)
}

func (m Map[K, V]) ForEach(consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}


func (m Map[K, V]) Collect() map[K]V {
	return m
}
