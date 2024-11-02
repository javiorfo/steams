package steams

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Filter(predicate func(K, V) bool) Steam2[K, V] {
	results := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return Map[K, V](results)
}

func (m Map[K, V]) MapToAny(mapper func(K, V) any) Steam2[K, any] {
	results := make(map[K]any, len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return Map[K, any](results)
}

func (m Map[K, V]) MapToString(mapper func(K, V) string) Steam2[K, string] {
	results := make(map[K]string, len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return Map[K, string](results)
}

func (m Map[K, V]) MapToInt(mapper func(K, V) int) Steam2[K, int] {
	results := make(map[K]int, len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return Map[K, int](results)
}

func (m Map[K, V]) FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any] {
	results := make(map[K]any)
	for k, v := range m {
		if predicate(k, v) {
			results[k] = mapper(k, v)
		}
	}
	return Map[K, any](results)
}

func (m Map[K, V]) ForEach(consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}

func (m Map[K, V]) Limit(limit int) Steam2[K, V] {
    results := make(map[K]V, 0)
    var counter int
    for k, v := range m {
        if limit > counter {
            break
        }
        results[k] = v
        counter++
    }
    return Map[K, V](results)
}

func (m Map[K, V]) Count() int {
    return len(m)
}

func (m Map[K, V]) ValuesToSteam() Steam[V] {
	res := make([]V, len(m))
	var index uint
	for _, v := range m {
		res[index] = v
		index++
	}
	return List[V](res)
}

func (m Map[K, V]) KeysToSteam() Steam[K] {
	res := make([]K, len(m))
	var index uint
	for k := range m {
		res[index] = k
		index++
	}
	return List[K](res)
}

func (m Map[K, V]) ToAnySteam(mapper func(K, V) any) Steam[any] {
	res := make([]any, len(m))
	var index uint
	for k, v := range m {
		res[index] = mapper(k, v)
		index++
	}
	return List[any](res)
}

func (m Map[K, V]) Collect() map[K]V {
	return m
}
