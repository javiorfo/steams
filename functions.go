package steams

import "fmt"

// Ordered is a type constraint that includes all ordered types.
// It allows for comparison of various numeric types and strings.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// OrderedStruct is an interface for structs that can be compared.
// It requires the implementation of a Compare method that returns an integer.
// The method should return a negative value if the receiver is less than the other,
// zero if they are equal, and a positive value if the receiver is greater.
type OrderedStruct[T any] interface {
	Compare(other T) int
}

// OrderStructDesc compares two OrderedStructs in descending order.
// It returns true if the first struct is less than the second.
func OrderStructDesc[T OrderedStruct[T]](a, b T) bool {
	return a.Compare(b) < 0
}

// OrderStructAsc compares two OrderedStructs in ascending order.
// It returns true if the first struct is greater than the second.
func OrderStructAsc[T OrderedStruct[T]](a, b T) bool {
	return a.Compare(b) > 0
}

// OrderDesc compares two Ordered values in descending order.
// It returns true if the first value is less than the second.
func OrderDesc[T Ordered](a, b T) bool {
	return a < b
}

// OrderAsc compares two Ordered values in ascending order.
// It returns true if the first value is greater than the second.
func OrderAsc[T Ordered](a, b T) bool {
	return a > b
}

func Println[T any](v T) {
    fmt.Println(v)
}
