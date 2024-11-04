package main

import (
	"fmt"

	"github.com/javiorfo/steams"
)

func main() {
	values := steams.Of(1, 2, 3, 4, 5, 6, 7)

	values.
		Filter(filter).
		Reverse().
		MapToString(mapper).
		ForEach(func(v string) { fmt.Println(v) })

	values.
		FilterMapToAny(filter, mapperToAny).
		ForEach(func(v any) { fmt.Println(v) })
}

func filter(value int) bool {
	return value < 5
}

func mapper(value int) string {
	return fmt.Sprintf("Value %d", value)
}

func mapperToAny(value int) any {
	return struct{ v int }{value}
}
