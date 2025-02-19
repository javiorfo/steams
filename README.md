# Steams
*Go functional programming library inspired (mostly) on Java Streams*

## Caveats
- This plugin requires Go 1.23
- Contains several Java style streams (called steams) and Optionals too.

## Intallation
```bash
go get -u github.com/javiorfo/steams@latest
```

## Example
#### More examples [here](https://github.com/javiorfo/steams/tree/master/examples)
```go
package main

import (
	"fmt"

	"github.com/javiorfo/steams"
)

var PeopleWithPets = []Person{
	{Name: "Carl", Age: 34, Pets: []Pet{}},
    {Name: "John", Age: 20, Pets: []Pet{{Name: "Bobby", Type: "DOG", Age: 2}, {Name: "Mike", Type: "DOG", Age: 12}}},
	{Name: "Grace", Age: 40, Pets: []Pet{{Name: "Pepe", Type: "DOG", Age: 4}, {Name: "Snowball", Type: "CAT", Age: 8}}},
	{Name: "Robert", Age: 40, Pets: []Pet{{Name: "Ronny", Type: "CAT", Age: 3}}},
}

func main() {
	steams.OfSlice(data.PeopleWithPets).
		Filter(func(p data.Person) bool {
			return p.Age > 21
		}).
		Peek(func(p data.Person) { fmt.Println("After Filter => Person:", p.Name) }).
		FlatMapToAny(func(p data.Person) steams.Steam[any] {
			results := make(steams.List[any], 0)
			for _, v := range p.Pets {
				results = append(results, v)
			}
			return results
		}).
		Peek(func(p any) { fmt.Println("After FlatMap = Pet:", p.(data.Pet).Name) }).
		Filter(func(p any) bool {
			animal, ok := p.(data.Pet)
			if ok {
				if animal.Type == data.CAT {
					return true
				}
			}
			return false
		}).
		Peek(func(p any) { fmt.Println("After second Filter => Pet:", p.(data.Pet).Name) }).
		GetCompared(comparator).IfPresentOrElse(print, func() { fmt.Println("No results") })

}

func comparator(a any, b any) bool {
	ageA := a.(data.Pet).Age
	ageB := b.(data.Pet).Age
	return ageA < ageB
}

func print(cat any) {
	younger := cat.(data.Pet)
	fmt.Printf("The younger cat of the list is %s, age %d", younger.Name, younger.Age)
}
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
	FilterMapToInt(predicate func(T) bool, mapper func(T) int) Steam[int]
	FilterMapToString(predicate func(T) bool, mapper func(T) string) Steam[string]
	FlatMapToAny(mapper func(T) Steam[any]) Steam[any]
	FlatMapToInt(mapper func(T) Steam[int]) Steam[int]
	FlatMapToString(mapper func(T) Steam[string]) Steam[string]
	ForEach(consumer func(T))
	ForEachWithIndex(consumer func(int, T))
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
	GetCompared(cmp func(T, T) bool) opt.Optional[T]
	FindFirst() opt.Optional[T]
	Last() opt.Optional[T]
	Position(predicate func(T) bool) opt.Optional[int]
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
	FilterMapToInt(predicate func(K, V) bool, mapper func(K, V) int) Steam2[K, int]
	FilterMapToString(predicate func(K, V) bool, mapper func(K, V) string) Steam2[K, string]
	ForEach(consumer func(K, V))
	Peek(consumer func(K, V)) Steam2[K, V]
	Limit(limit int) Steam2[K, V]
	AllMatch(predicate func(K, V) bool) bool
	AnyMatch(predicate func(K, V) bool) bool
	NoneMatch(predicate func(K, V) bool) bool
	Sorted(cmp func(K, K) bool) Steam2[K, V]
	GetCompared(cmp func(K, K) bool) opt.Optional[Pair[K, V]]
	Count() int
	Collect() map[K]V
	KeysToSteam() Steam[K]
	ValuesToSteam() Steam[V]
	ToAnySteam(mapper func(K, V) any) Steam[any]
}
```

## Integration functions
```go
func Of[T any](args ...T) Steam[T]
func OfSlice[T any](slice []T) Steam[T]
func OfMap[K comparable, V any](m map[K]V) Steam2[K, V]
func Mapping[T, R any](s Steam[T], mapper func(T) R) Steam[R]
func Distinct[T comparable](s Steam[T]) Steam[T]
func CollectSteamToSteam2[K comparable, V, T any](s Steam[T], keyFunc func(T) K, valueFunc func(T) V) Steam2[K, V] 
func CollectSteam2ToSteam[K comparable, V, R any](s Steam2[K, V], mapper func(K, V) R) Steam[R]
func GroupBy[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, Steam[V]] 
func GroupByCounting[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, int]
func Zip[T, R any](s1 Steam[T], s2 Steam[R]) Steam[struct { first  T; second R }]
```

## Optionals
- [Test examples](https://github.com/javiorfo/steams/blob/master/opt/optional_test.go)

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
