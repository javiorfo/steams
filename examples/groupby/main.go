package main

import (
	"fmt"

	"github.com/javiorfo/steams/v2"
	"github.com/javiorfo/steams/v2/examples/data"
)

func main() {
	fmt.Println("Get all the animals")
	animals := steams.FlatMap(steams.FromSlice(data.PeopleWithPets), func(p data.Person) steams.It[data.Pet] {
		results := make([]data.Pet, 0)
		for _, v := range p.Pets {
			results = append(results, v)
		}
		return steams.FromSlice(results)
	})
	fmt.Println(animals)

	fmt.Println()
	fmt.Println("GroupBy type of animal")
	steams.GroupBy(animals, classifier).ForEach(func(s string, i steams.It[data.Pet]) {
		fmt.Println(s, i.Collect())
	})

	fmt.Println()
	fmt.Println("GroupByCounting type of animal")
	steams.GroupByCounting(animals, classifier).ForEach(func(s string, i int) {
		fmt.Println(s, i)
	})
}

func classifier(p data.Pet) string {
	return p.Type
}
