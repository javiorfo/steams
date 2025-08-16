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
		ForEach(steams.Println)

	values.
		FilterMap(filter, mapperPlusTen).
		ForEach(steams.Println)
}

func filter(value int) bool {
	return value < 5
}

func mapper(value int) string {
	return fmt.Sprintf("Value %d", value)
}

func mapperPlusTen(value int) int {
	return value + 10
}
