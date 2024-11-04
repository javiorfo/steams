package steams

type Steam[T any] interface {
	Filter(predicate func(T) bool) Steam[T]
	MapToAny(mapper func(T) any) Steam[any]
	MapToInt(mapper func(T) int) Steam[int]
	MapToString(mapper func(T) string) Steam[string]
	FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any]
	FlatMap(mapper func(T) Steam[T]) Steam[T]
	FlatMapToAny(mapper func(T) Steam[any]) Steam[any]
	ForEach(consumer func(T))
	Peek(consumer func(T)) Steam[T]
	Limit(limit int) Steam[T]
	AllMatch(predicate func(T) bool) bool
	AnyMatch(predicate func(T) bool) bool
	NoneMatch(predicate func(T) bool) bool
	TakeWhile(predicate func(T) bool) Steam[T]
	DropWhile(predicate func(T) bool) Steam[T]
	Reduce(initValue T, acc func(T, T) T) T
	Reverse() Steam[T]
	Sorted(cmp func(T, T) bool) Steam[T]
    GetCompared(cmp func(T, T) bool) (*T, bool)
	FindFirst() (*T, bool)
	Last() (*T, bool)
	Position(predicate func(T) bool) (*int, bool)
	Skip(n int) Steam[T]
	Count() int
	Collect() []T
    Length() int
}

type Steam2[K comparable, V any] interface {
	Filter(predicate func(K, V) bool) Steam2[K, V]
	MapToAny(mapper func(K, V) any) Steam2[K, any]
	MapToInt(mapper func(K, V) int) Steam2[K, int]
	MapToString(mapper func(K, V) string) Steam2[K, string]
	FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any]
	ForEach(consumer func(K, V))
	Peek(consumer func(K, V)) Steam2[K, V]
	Limit(limit int) Steam2[K, V]
	AllMatch(predicate func(K, V) bool) bool
	AnyMatch(predicate func(K, V) bool) bool
	NoneMatch(predicate func(K, V) bool) bool
	TakeWhile(predicate func(K, V) bool) Steam2[K, V]
	DropWhile(predicate func(K, V) bool) Steam2[K, V]
	FindFirst() (*Pair[K, V], bool)
	Last() (*Pair[K, V], bool)
	Skip(n int) Steam2[K, V]
	Sorted(cmp func(K, K) bool) Steam2[K, V]
    GetCompared(cmp func(K, K) bool) (*Pair[K, V], bool)
	Count() int
	Collect() map[K]V
	KeysToSteam() Steam[K]
	ValuesToSteam() Steam[V]
	ToAnySteam(mapper func(K, V) any) Steam[any]
    Length() int
}

type Comparator[T any] interface {
	Compare(a, b T) bool
}

func Distinct[T comparable](s Steam[T]) Steam[T] {
	m := make(map[T]bool)
	slice := s.Collect()
	results := make(List[T], 0, s.Length())
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

func Zip[T, R any](s1 Steam[T], s2 Steam[R]) Steam[struct{ first T; second R }] {
    slice1 := s1.Collect()
    slice2 := s2.Collect()
    if len(slice1) != len(slice2) {
        panic("slices must have the same length")
    }

    result := make(List[struct{ first T; second R }], len(slice1))
    for i := range slice1 {
        result[i] = struct{ first T; second R }{slice1[i], slice2[i]}
    }
    return result
}

func Of[T any](args ...T) Steam[T] {
	return List[T](args)
}

func OfMap[K comparable, V any](m map[K]V) Steam2[K, V] {
	return Map[K, V](m)
}
