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
    Limit(limit int) Steam[T]
    Count() int
	Collect() []T
}

type Steam2[K comparable, V any] interface {
	Filter(predicate func(K, V) bool) Steam2[K, V]
    MapToAny(mapper func(K, V) any) Steam2[K, any]
    MapToInt(mapper func(K, V) int) Steam2[K, int]
    MapToString(mapper func(K, V) string) Steam2[K, string]
    FilterMapToAny(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, any]
    ForEach(consumer func(K, V))
    Limit(limit int) Steam2[K, V]
    Count() int
	Collect() map[K]V
    KeysToSteam() Steam[K]
    ValuesToSteam() Steam[V]
    ToAnySteam(mapper func(K, V) any) Steam[any]
}

func CollectSteamToSteam2[K comparable, V any](s Steam[V], keyFunc func(V) K, valueFunc func(V) V) Steam2[K, V] {
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

func Of[T any](args ...T) Steam[T] {
	return List[T](args)
}

func OfMap[K comparable, V any](m map[K]V) Steam2[K, V] {
	return Map[K, V](m)
}

// fold, reduce, drop, distinct, sort, min, max, sum, average
// skip, findAny, allMatch, anyMatch, findFirst, dropWhile, takeWhile, noneMatch
// Rust: position, rev, zip, product, nth, last
