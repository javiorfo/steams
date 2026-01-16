package main

import (
	"fmt"

	"github.com/javiorfo/steams"
	"github.com/javiorfo/steams/examples/data"
)

func main() {
	persons := steams.OfSlice(data.PeopleWithPets).
		Filter(func(p data.Person) bool { return p.Age > 21 }).
		Peek(func(p data.Person) { fmt.Println("After Filter => Person:", p.Name) })

	steams.FlatMapper(persons, func(p data.Person) steams.Steam[data.Pet] {
		return steams.OfSlice(p.Pets)
	}).
		Peek(func(p data.Pet) { fmt.Println("After FlatMap = Pet:", p.Name) }).
		Filter(isCat).
		Peek(func(p data.Pet) { fmt.Println("After second Filter => Pet:", p.Name) }).
		GetCompared(comparator).Inspect(print).OrPanic("No results")
}

func isCat(p data.Pet) bool {
	if p.Type == data.CAT {
		return true
	}
	return false
}

func comparator(a data.Pet, b data.Pet) bool {
	return a.Age < b.Age
}

func print(p data.Pet) {
	fmt.Printf("The younger cat of the list is %s, age %d", p.Name, p.Age)
}
