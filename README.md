# Steams
*Go functional programming library using iterators*

## Caveats
- This library requires Go 1.23+
- Contains several streams (called steams) using iterators, so the streams are mostly lazy. 
- Version 1 is a more Javaish approach, version 2 is a more Rustish approach (so to say and considering what Go allows).

## Insstallation
```bash
go get -u github.com/javiorfo/steams/v2@latest
```

## Example
#### More examples [here](https://github.com/javiorfo/steams/tree/master/examples)
```go
package main

import (
  "fmt"

  "github.com/javiorfo/steams/v2"
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
  persons := steams.FromSlice(PeopleWithPets).
	  Filter(func(p Person) bool { return p.Age > 21 }).
	  Inspect(func(p Person) { fmt.Println("After Filter => Person:", p.Name) })

  steams.FlatMap(persons, func(p Person) steams.It[Pet] {
	  return steams.FromSlice(p.Pets)
  }).
	  Inspect(func(p Pet) { fmt.Println("After FlatMap = Pet:", p.Name) }).
	  Filter(isCat).
	  Inspect(func(p Pet) { fmt.Println("After second Filter => Pet:", p.Name) }).
      Compare(comparator).
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

## Api
```go
// It[T] is type based on iter.Seq[T] for a collection of elements,
// providing various methods for functional-style processing.
func (it It[T]) AsSeq() iter.Seq[T]
func (it It[T]) Filter(predicate func(T) bool) It[T]
func (it It[T]) Map(mapper func(T) T) It[T]
func (it It[T]) MapToString(mapper func(T) string) It[string]
func (it It[T]) MapToInt(mapper func(T) int) It[int]
func (it It[T]) FilterMap(mapper func(T) nilo.Option[T]) It[T]
func (it It[T]) FilterMapToString(mapper func(T) nilo.Option[string]) It[string]
func (it It[T]) FilterMapToInt(mapper func(T) nilo.Option[int]) It[int]
func (it It[T]) FlatMap(mapper func(T) It[T]) It[T]
func (it It[T]) FlatMapToString(mapper func(T) It[string]) It[string]
func (it It[T]) FlatMapToInt(mapper func(T) It[int]) It[int]
func (it It[T]) Take(n int) It[T]
func (it It[T]) Count() int
func (it It[T]) ForEach(consumer func(T))
func (it It[T]) ForEachIdx(consumer func(int, T))
func (it It[T]) Inspect(inspector func(T)) It[T]
func (it It[T]) All(predicate func(T) bool) bool
func (it It[T]) Any(predicate func(T) bool) bool
func (it It[T]) None(predicate func(T) bool) bool
func (it It[T]) First() nilo.Option[T]
func (it It[T]) Find(predicate func(T) bool) nilo.Option[T]
func (it It[T]) TakeWhile(predicate func(T) bool) It[T]
func (it It[T]) SkipWhile(predicate func(T) bool) It[T]
func (it It[T]) Fold(initValue T, acc func(T, T) T) T
func (it It[T]) RFold(initValue T, acc func(T, T) T) T
func (it It[T]) Reverse() It[T]
func (it It[T]) Position(predicate func(T) bool) nilo.Option[int]
func (it It[T]) RPosition(predicate func(T) bool) nilo.Option[int]
func (it It[T]) Enumerate() iter.Seq2[int, T]
func (it It[T]) Last() nilo.Option[T]
func (it It[T]) Skip(n int) It[T]
func (it It[T]) SortBy(cmp func(T, T) int) It[T]
func (it It[T]) Compare(cmp func(T, T) bool) nilo.Option[T]
func (it It[T]) Collect() []T
func (it It[T]) Chain(i2 It[T]) It[T]
func (it It[T]) Nth(n int) nilo.Option[T]
func (it It[T]) Partition(politer func(T) bool) (It[T], It[T])

// It2[K, V] is type based on iter.Seq2[K, V] for a map of elements,
// providing various methods for functional-style processing.
func (it It2[K, V]) Filter(predicate func(K, V) bool) It2[K, V]
func (it It2[K, V]) Map(mapper func(K, V) (K, V)) It2[K, V]
func (it It2[K, V]) MapToString(mapper func(K, V) (K, string)) It2[K, string]
func (it It2[K, V]) MapToInt(mapper func(K, V) (K, int)) It2[K, int]
func (it It2[K, V]) ForEach(consumer func(K, V))
func (it It2[K, V]) SortBy(cmp func(K, K) bool) It2[K, V]
func (it It2[K, V]) Inspect(consumer func(K, V)) It2[K, V]
func (it It2[K, V]) Take(n int) It2[K, V]
func (it It2[K, V]) Values() It[V]
func (it It2[K, V]) Keys() It[K]
func (it It2[K, V]) All(predicate func(K, V) bool) bool
func (it It2[K, V]) Any(predicate func(K, V) bool) bool
func (it It2[K, V]) None(predicate func(K, V) bool) bool
func (it It2[K, V]) Compare(cmp func(K, K) bool) nilo.Option[Entry[K, V]]
func (it It2[K, V]) Collect() map[K]V
func (it It2[K, V]) Count() int
```

## Integration functions
```go
func From[T any](args ...T) It[T]
func FromSlice[T any](slice []T) It[T]
func FromMap[K comparable, V any](m map[K]V) It2[K, V]
func Distinct[T comparable](i It[T]) It[T]
func Map[T any, U any](i It[T], transform func(T) U) It[U]
func FlatMap[T any, U any](i It[T], transform func(T) It[U]) It[U]
func Fold[T any, R any](i It[T], initial R, accumulator func(R, T) R) R
func RFold[T any, R any](i It[T], initial R, accumulator func(T, R) R) R
func Flatten[V any](nested It[iter.Seq[V]]) It[V]
func GroupBy[K comparable, V any](i It[V], classifier func(V) K) It2[K, It[V]]
func GroupByCounting[K comparable, V any](i It[V], classifier func(V) K) It2[K, int]
func Zip[T, R any](i1 It[T], i2 It[R]) It[struct {First T; Second R}]
func CollectItToIt2[T, K comparable, V any](i It[T], keyFunc func(T) K, valueFunc func(T) V) It2[K, V]
func CollectIt2ToIt[K comparable, V, R any](i It2[K, V], mapper func(K, V) R) It[R]
func ChainAll[V any](its ...It[V]) It[V]
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
