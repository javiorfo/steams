package steams

import "slices"

type List[T any] []T

func ListOf[T any](args ...T) Steam[T] {
	return List[T](args)
}

func (list List[T]) Filter(predicate func(T) bool) Steam[T] {
	results := make([]T, 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		}
	}
	return List[T](results)
}

func (list List[T]) MapToAny(mapper func(T) any) Steam[any] {
	results := make([]any, len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return List[any](results)
}

func (list List[T]) MapToString(mapper func(T) string) Steam[string] {
	results := make([]string, len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return List[string](results)
}

func (list List[T]) MapToInt(mapper func(T) int) Steam[int] {
	results := make([]int, len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return List[int](results)
}

func (list List[T]) FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any] {
	results := make([]any, 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return List[any](results)
}

func (list List[T]) FlatMap(mapper func(T) Steam[T]) Steam[T] {
	results := make(List[T], len(list))
    for _, v := range list {
       results = slices.Concat(results, mapper(v).(List[T]))
    }
    return List[T](results)
}

func (list List[T]) FlatMapToAny(mapper func(T) Steam[any]) Steam[any] {
	results := make(List[any], len(list))
    for _, v := range list {
       results = slices.Concat(results, mapper(v).(List[any]))
    }
    return List[any](results)
}

func (list List[T]) Limit(limit int) Steam[T] {
    results := make([]T, 0)
    for i := 0; i < len(list) && i < limit; i++ {
        results = append(results, list[i])
    }
    return List[T](results)
}

func (list List[T]) Count() int {
    return len(list)
}

func (list List[T]) ForEach(consumer func(T)) {
	for _, v := range list {
		consumer(v)
	}
}

func (list List[T]) Collect() []T {
	return list
}
