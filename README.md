# Steams
*Go functional programming library inspired (mostly) on Java Streams*

## Caveats
- This library requires Go 1.23+
- Contains several Java style streams (called steams).

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

type Person struct {
  Name string
  Age  int
  Pets []Pet
}

type Pet struct {
  Name string
  Type string
  Age  int
}

const DOG = "DOG"
const CAT = "CAT"

var PeopleWithPets = []Person{
  {Name: "Carl", Age: 34, Pets: []Pet{}},
  {Name: "John", Age: 20, Pets: []Pet{{Name: "Bobby", Type: DOG, Age: 2}, {Name: "Mike", Type: DOG, Age: 12}}},
  {Name: "Grace", Age: 40, Pets: []Pet{{Name: "Pepe", Type: DOG, Age: 4}, {Name: "Snowball", Type: CAT, Age: 8}}},
  {Name: "Robert", Age: 40, Pets: []Pet{{Name: "Ronny", Type: CAT, Age: 3}}},
}


func main() {
  persons := steams.OfSlice(PeopleWithPets).
	  Filter(func(p Person) bool { return p.Age > 21 }).
	  Peek(func(p Person) { fmt.Println("After Filter => Person:", p.Name) })

  steams.FlatMapper(persons, func(p Person) steams.Steam[Pet] {
	  return steams.OfSlice(p.Pets)
  }).
	  Peek(func(p Pet) { fmt.Println("After FlatMap = Pet:", p.Name) }).
	  Filter(isCat).
	  Peek(func(p Pet) { fmt.Println("After second Filter => Pet:", p.Name) }).
      GetCompared(comparator).
      Inspect(print).
      OrPanic("No results")
}

func isCat(p Pet) bool {
  if p.Type == data.CAT {
	  return true
  }
  return false
}

func comparator(a Pet, b Pet) bool {
  return a.Age < b.Age
}

func print(p Pet) {
  fmt.Printf("The younger cat of the list is %s, age %d", p.Name, p.Age)
}

```

## Interfaces
```go
// Steam[T] is an interface for a collection of elements of type T,
// providing various methods for functional-style processing.
type Steam[T any] interface {
  Filter(predicate func(T) bool) Steam[T]
  Map(mapper func(T) T) Steam[T]
  MapToInt(mapper func(T) int) Steam[int]
  MapToString(mapper func(T) string) Steam[string]
  FilterMap(predicate func(T) bool, mapper func(T) T) Steam[T]
  FilterMapToInt(predicate func(T) bool, mapper func(T) int) Steam[int]
  FilterMapToString(predicate func(T) bool, mapper func(T) string) Steam[string]
  FlatMap(mapper func(T) Steam[T]) Steam[T]
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
  GetCompared(cmp func(T, T) bool) nilo.Option[T]
  FindFirst() nilo.Option[T]
  FindOne(predicate func(T) bool) nilo.Option[T]
  Last() nilo.Option[T]
  Position(predicate func(T) bool) nilo.Option[int]
  Skip(n int) Steam[T]
  Count() int
  Collect() []T
}

// Steam2[K, V] is an interface for a map of elements of type K and V,
// providing various methods for functional-style processing.
type Steam2[K comparable, V any] interface {
  Filter(predicate func(K, V) bool) Steam2[K, V]
  Map(mapper func(K, V) any) Steam2[K, V]
  MapToInt(mapper func(K, V) int) Steam2[K, int]
  MapToString(mapper func(K, V) string) Steam2[K, string]
  FilterMap(predicate func(K, V) bool, mapper func(K, V) any) Steam2[K, V]
  FilterMapToInt(predicate func(K, V) bool, mapper func(K, V) int) Steam2[K, int]
  FilterMapToString(predicate func(K, V) bool, mapper func(K, V) string) Steam2[K, string]
  ForEach(consumer func(K, V))
  Peek(consumer func(K, V)) Steam2[K, V]
  Limit(limit int) Steam2[K, V]
  AllMatch(predicate func(K, V) bool) bool
  AnyMatch(predicate func(K, V) bool) bool
  NoneMatch(predicate func(K, V) bool) bool
  Sorted(cmp func(K, K) bool) Steam2[K, V]
  GetCompared(cmp func(K, K) bool) nilo.Option[Pair[K, V]]
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
func Mapper[T, R any](s Steam[T], mapper func(T) R) Steam[R]
func FlatMapper[T, R any](s Steam[T], mapper func(T) Steam[R]) Steam[R]
func Distinct[T comparable](s Steam[T]) Steam[T]
func CollectSteamToSteam2[K comparable, V, T any](s Steam[T], keyFunc func(T) K, valueFunc func(T) V) Steam2[K, V] 
func CollectSteam2ToSteam[K comparable, V, R any](s Steam2[K, V], mapper func(K, V) R) Steam[R]
func GroupBy[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, Steam[V]] 
func GroupByCounting[K comparable, V any](s Steam[V], classifier func(V) K) Steam2[K, int]
func Zip[T, R any](s1 Steam[T], s2 Steam[R]) Steam[struct { first  T; second R }]
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
