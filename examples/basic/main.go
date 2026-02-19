package main

import (
	"fmt"

	"github.com/javiorfo/nilo"
	"github.com/javiorfo/steams/v2"
)

func main() {
	values := steams.From(1, 2, 3, 4, 5, 6, 7)

	values.
		Filter(filter).
		Reverse().
		MapToString(mapper).
		ForEach(func(s string) { fmt.Println(s) })

	values.
		FilterMap(mapperPlusTen).
		ForEach(func(i int) { fmt.Println(i) })
}

func filter(value int) bool {
	return value < 5
}

func mapper(value int) string {
	return fmt.Sprintf("Value %d", value)
}

func mapperPlusTen(value int) nilo.Option[int] {
	if value < 5 {
		return nilo.Value(value + 10)
	}
	return nilo.Nil[int]()
}
