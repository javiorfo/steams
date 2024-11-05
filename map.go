package steams

import (
	"sort"

	"github.com/javiorfo/steams/opt"
)

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
		if counter >= limit {
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

func (m Map[K, V]) GetCompared(cmp func(K, K) bool) opt.Optional[Pair[K, V]] {
	if len(m) == 0 {
		return opt.Empty[Pair[K, V]]()
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
	return opt.Of(*item)
}

func (m Map[K, V]) Collect() map[K]V {
	return m
}
