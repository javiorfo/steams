package steams

type Steam[T any] interface {
	Filter(predicate func(T) bool) Steam[T]
    MapToAny(mapper func(T) any) Steam[any]
    MapToInt(mapper func(T) int) Steam[int]
    MapToString(mapper func(T) string) Steam[string]
    FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any]
    ForEach(consumer func(T))
	Collect() []T
}

type Steam2[K comparable, V any] interface {
	Filter(predicate func(K, V) bool) Steam2[K, V]
    ForEach(consumer func(K, V))
	Collect() map[K]V
    CollectKeysToSteam() Steam[K]
    CollectValuesToSteam() Steam[V]
//     CollectToAnySteam(mapper func(K, V) any) Steam[any]
}

func (list List[T]) CollectToMap() map[any]T {
	return nil
}

func CollectSteamToSteam2[K comparable, V any](s Steam[V], keyFunc func(V) K, valueFunc func(V) V) Steam2[K, V] {
    m := make(map[K]V)
    for _, v := range s.Collect() {
        m[keyFunc(v)] = valueFunc(v)
    }
	return Map[K, V](m)
}

func SteamOf[T any](args ...T) Steam[T] {
	return List[T](args)
}

func Steam2Of[K comparable, V any](m map[K]V) Steam2[K, V] {
	return Map[K, V](m)
}

// fold, reduce, flatMap, drop, distinct, sort, peek, limit, min, max, sum, average
// count, skip, findAny, allMatch, anyMatch, findFirst, dropWhile, takeWhile, noneMatch
