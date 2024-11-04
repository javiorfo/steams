# Steams
*Go functional programming library inspired (mostly) on Java Streams*

## Caveats
- This plugin requires Go 1.23

## Intallation
```bash
go get -u https://github.com/javiorfo/steams
```

## Interfaces
```go
// Steam[T] is an interface for a collection of elements of type T,
// providing various methods for functional-style processing.
type Steam[T any] interface {
	Filter(predicate func(T) bool) Steam[T]
	MapToAny(mapper func(T) any) Steam[any]
	MapToInt(mapper func(T) int) Steam[int]
	MapToString(mapper func(T) string) Steam[string]
	FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any]
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
}

// Steam2[K, V] is an interface for a map of elements of type K and V,
// providing various methods for functional-style processing.
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
	Sorted(cmp func(K, K) bool) Steam2[K, V]
	GetCompared(cmp func(K, K) bool) (*Pair[K, V], bool)
	Count() int
	Collect() map[K]V
	KeysToSteam() Steam[K]
	ValuesToSteam() Steam[V]
	ToAnySteam(mapper func(K, V) any) Steam[any]
}
```

## Examples

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
