package main

import "github.com/javiorfo/steams"

func main() {
	s1 := steams.Of(1, 2, 3)
	s2 := steams.Of("a", "b", "c")
    steams.Zip(s1, s2).ForEach(steams.Println)
}
