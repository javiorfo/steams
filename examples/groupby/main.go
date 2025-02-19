package main

import (
	"fmt"

	"github.com/javiorfo/steams"
	"github.com/javiorfo/steams/examples/data"
)

func main() {
    fmt.Println("Get all the animals")
	animals := steams.OfSlice(data.PeopleWithPets).FlatMapToAny(func(p data.Person) steams.Steam[any] {
		results := make(steams.List[any], 0)
		for _, v := range p.Pets {
			results = append(results, v)
		}
		return results
	})
    fmt.Println(animals)

    fmt.Println()
    fmt.Println("GroupBy type of animal")
	steams.GroupBy(animals, classifier).ForEach(steams.Println2)
    
    fmt.Println()
    fmt.Println("GroupByCounting type of animal")
	steams.GroupByCounting(animals, classifier).ForEach(steams.Println2)
}

func classifier(v any) string {
	animal, ok := v.(data.Pet)
	if ok {
		return animal.Type
	}
	return "None"
}
