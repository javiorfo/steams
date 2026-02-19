package steams

import (
	"iter"
	"maps"
)

type it2[K, V any] iter.Seq2[K, V]

func OfMap2[K comparable, V any](m map[K]V) it2[K, V] {
	return it2[K, V](maps.All(m))
}

func (i it2[K, V]) Filter(predicate func(K, V) bool) it2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range i {
			if predicate(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
