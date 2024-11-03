package steams

type Map[K comparable, V any] map[K]V

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func (m Map[K, V]) Filter(predicate func(K, V) bool) Steam2[K, V] {
	results := make(Map[K, V])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return results
}

func (m Map[K, V]) MapToAny(mapper func(K, V) any) Steam2[K, any] {
	results := make(Map[K, any], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

func (m Map[K, V]) MapToString(mapper func(K, V) string) Steam2[K, string] {
	results := make(Map[K, string], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

func (m Map[K, V]) MapToInt(mapper func(K, V) int) Steam2[K, int] {
	results := make(Map[K, int], len(m))
	for k, v := range m {
		results[k] = mapper(k, v)
	}
	return results
}

func (m Map[K, V]) FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any] {
	results := make(Map[K, any])
	for k, v := range m {
		if predicate(k, v) {
			results[k] = mapper(k, v)
		}
	}
	return results
}

func (m Map[K, V]) ForEach(consumer func(K, V)) {
	for k, v := range m {
		consumer(k, v)
	}
}

func (m Map[K, V]) Peek(consumer func(K, V)) Steam2[K, V] {
	for k, v := range m {
		consumer(k, v)
	}
	return m
}

func (m Map[K, V]) Limit(limit int) Steam2[K, V] {
	results := make(Map[K, V], 0)
	var counter int
	for k, v := range m {
		if limit > counter {
			break
		}
		results[k] = v
		counter++
	}
	return results
}

func (m Map[K, V]) Count() int {
	return len(m)
}

func (m Map[K, V]) ValuesToSteam() Steam[V] {
	res := make(List[V], len(m))
	var index uint
	for _, v := range m {
		res[index] = v
		index++
	}
	return res
}

func (m Map[K, V]) KeysToSteam() Steam[K] {
	res := make(List[K], len(m))
	var index uint
	for k := range m {
		res[index] = k
		index++
	}
	return res
}

func (m Map[K, V]) ToAnySteam(mapper func(K, V) any) Steam[any] {
	res := make(List[any], len(m))
	var index uint
	for k, v := range m {
		res[index] = mapper(k, v)
		index++
	}
	return res
}

func (m Map[K, V]) AllMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func (m Map[K, V]) AnyMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func (m Map[K, V]) NoneMatch(predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return false
		}
	}
	return true
}

func (m Map[K, V]) FindFirst() (*Pair[K, V], bool) {
	for k, v := range m {
		return &Pair[K, V]{Key: k, Value: v}, true
	}
	return nil, false
}

func (m Map[K, V]) TakeWhile(predicate func(K, V) bool) Steam2[K, V] {
	results := make(Map[K, V], 0)
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		} else {
			break
		}
	}
	return results
}

func (m Map[K, V]) DropWhile(predicate func(K, V) bool) Steam2[K, V] {
	results := make(Map[K, V], 0)
	for k, v := range m {
		if !predicate(k, v) {
			results[k] = v
		}
	}
	return results
}

func (m Map[K, V]) Skip(n int) Steam2[K, V] {
	length := len(m)
	if length > n {
		length = length - n
	} else {
		return *new(Map[K, V])
	}

	results := make(Map[K, V], length)
	var count int
	for k := range m {
		if count == n {
			break
		}
		results[k] = m[k]
		count++
	}
	return results
}

func (m Map[K, V]) Last() (*Pair[K, V], bool) {
	pair := Pair[K, V]{}
	exists := false
	for k, v := range m {
		pair.Key = k
		pair.Value = v
		exists = true
	}
	return &pair, exists
}

func (m Map[K, V]) Collect() map[K]V {
	return m
}
