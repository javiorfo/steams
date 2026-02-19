package main

import (
	"fmt"

	"github.com/javiorfo/steams/v2"
	"github.com/javiorfo/steams/v2/examples/data"
)

func main() {
	persons := steams.FromSlice(data.PeopleWithPets).
		Filter(func(p data.Person) bool { return p.Age > 21 }).
		Inspect(func(p data.Person) { fmt.Println("After Filter => Person:", p.Name) })

	steams.FlatMap(persons, func(p data.Person) steams.It[data.Pet] {
		return steams.FromSlice(p.Pets)
	}).
		Inspect(func(p data.Pet) { fmt.Println("After FlatMap = Pet:", p.Name) }).
		Filter(isCat).
		Inspect(func(p data.Pet) { fmt.Println("After second Filter => Pet:", p.Name) }).
		Compare(comparator).Inspect(print).OrPanic("No results")
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
