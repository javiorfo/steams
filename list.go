package steams

import (
	"slices"
	"sort"

	"github.com/javiorfo/steams/opt"
)

type List[T any] []T

func ListOf[T any](args ...T) Steam[T] {
	return List[T](args)
}

func (list List[T]) Filter(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

func (list List[T]) MapToAny(mapper func(T) any) Steam[any] {
	results := make(List[any], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) MapToString(mapper func(T) string) Steam[string] {
	results := make(List[string], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) MapToInt(mapper func(T) int) Steam[int] {
	results := make(List[int], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any] {
	results := make(List[any], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return results
}

func (list List[T]) FlatMapToAny(mapper func(T) Steam[any]) Steam[any] {
	results := make(List[any], 0, list.Count())
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[any]))
	}
	return results
}

func (list List[T]) Limit(limit int) Steam[T] {
	results := make(List[T], 0)
	for i := 0; i < len(list) && i < limit; i++ {
		results = append(results, list[i])
	}
	return results
}

func (list List[T]) Count() int {
	return len(list)
}

func (list List[T]) ForEach(consumer func(T)) {
	for _, v := range list {
		consumer(v)
	}
}

func (list List[T]) Peek(consumer func(T)) Steam[T] {
	for _, v := range list {
		consumer(v)
	}
	return list
}

func (list List[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (list List[T]) AnyMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (list List[T]) NoneMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return false
		}
	}
	return true
}

func (list List[T]) FindFirst() opt.Optional[T] {
	if len(list) > 0 {
		return opt.Of(list[0])
	}
	return opt.Empty[T]()
}

func (list List[T]) TakeWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		} else {
			break
		}
	}
	return results
}

func (list List[T]) DropWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if !predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

func (list List[T]) Reduce(initValue T, acc func(T, T) T) T {
	result := initValue
	for _, v := range list {
		result = acc(result, v)
	}
	return result
}

func (list List[T]) Reverse() Steam[T] {
	length := len(list)
	results := make(List[T], length)
	index := length - 1
	for i := range list {
		results[i] = list[index]
		index--
	}
	return results
}

func (list List[T]) Position(predicate func(T) bool) opt.Optional[int] {
	for i, v := range list {
		if predicate(v) {
			return opt.Of(i)
		}
	}
	return opt.Empty[int]()
}

func (list List[T]) Last() opt.Optional[T] {
	length := list.Count()
	if length > 0 {
		return opt.Of(list[length-1])
	}
	return opt.Empty[T]()
}

func (list List[T]) Skip(n int) Steam[T] {
	length := len(list)
	if length > n {
		length = length - n
	} else {
		return List[T]{}
	}

	results := make(List[T], length)
	for i := 0; i < length; i++ {
		results[i] = list[i+n]
	}
	return results
}

func (list List[T]) Sorted(cmp func(T, T) bool) Steam[T] {
	slice := list.Collect()
	results := make(List[T], len(slice))
	copy(results, slice)
	sort.Slice(results, func(i, j int) bool {
		return cmp(results[i], results[j])
	})
	return results
}

func (list List[T]) GetCompared(cmp func(T, T) bool) opt.Optional[T] {
	if len(list) == 0 {
		return opt.Empty[T]()
	}
	item := list[0]
	for i := 1; i < len(list); i++ {
		if cmp(list[i], item) {
			item = list[i]
		}
	}
	return opt.Of(item)
}

func (list List[T]) Collect() []T {
	return list
}
