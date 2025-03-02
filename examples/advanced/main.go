package main

import (
	"fmt"

	"github.com/javiorfo/steams"
	"github.com/javiorfo/steams/examples/data"
)

func main() {
	steams.OfSlice(data.PeopleWithPets).
		Filter(func(p data.Person) bool { return p.Age > 21 }).
		Peek(func(p data.Person) { fmt.Println("After Filter => Person:", p.Name) }).
		FlatMapToAny(func(p data.Person) steams.Steam[any] {
			return steams.OfSlice(p.Pets).MapToAny(func(p data.Pet) any { return any(p) })
		}).
		Peek(func(p any) { fmt.Println("After FlatMap = Pet:", p.(data.Pet).Name) }).
		Filter(isCat).
		Peek(func(p any) { fmt.Println("After second Filter => Pet:", p.(data.Pet).Name) }).
		GetCompared(comparator).IfPresentOrElse(print, func() { fmt.Println("No results") })

}

func isCat(p any) bool {
	animal, ok := p.(data.Pet)
	if ok {
		if animal.Type == data.CAT {
			return true
		}
	}
	return false
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
