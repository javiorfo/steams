package main

import (
	"fmt"

	"github.com/javiorfo/steams"
	"github.com/javiorfo/steams/examples/data"
)

func main() {
	steams.OfSlice(data.PersonsWithPets).
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
