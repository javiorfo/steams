package main

import (
	"fmt"

	"github.com/javiorfo/steams/v2"
)

func main() {
	s1 := steams.From(1, 2, 3)
	s2 := steams.From("a", "b", "c")
	steams.Zip(s1, s2).ForEach(func(s struct {
		First  int
		Second string
	}) {
		fmt.Println(s.First, s.Second)
	})
}
