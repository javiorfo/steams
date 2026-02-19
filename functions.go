package steams

import "fmt"

// Number is a type constraint that includes all sumable types.
// It allows for summing various numeric types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Ordered is a type constraint that includes all ordered types.
// It allows for comparison of various numeric types and strings.
type Ordered interface {
	Number | ~string
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

// OrderDesc compares two Ordered values in descending order.
// It returns true if the first value is less than the second.
func OrderDesc2[T Ordered](a, b T) int {
	if a < b {
		return -1
	}
	return 0
}

// OrderAsc compares two Ordered values in ascending order.
// It returns true if the first value is greater than the second.
func OrderAsc[T Ordered](a, b T) bool {
	return a > b
}

// Println prints the value of the provided argument v to the standard output.
// It can accept any type of value due to the use of a type parameter T.
func Println[T any](v T) {
	fmt.Println(v)
}

// Println prints the values of the provided arguments k and v to the standard output.
// It can accept a comparable K value and any type to second value
// due to the use of a type parameter T.
func Println2[K comparable, T any](k K, v T) {
	fmt.Println(k, v)
}

// Min returns true if the first argument a is greater than the second argument b.
// It is intended to be used with types that implement the Ordered interface,
// which allows for comparison operations.
func Min[T Ordered](a, b T) bool {
	return a < b
}

// Max returns true if the first argument a is less than the second argument b.
// It is intended to be used with types that implement the Ordered interface,
// which allows for comparison operations.
func Max[T Ordered](a, b T) bool {
	return a > b
}

// Sum returns the sum of two numbers a and b.
// It is intended to be used with types that implement the Number interface,
// which allows for addition operations.
func Sum[T Number](a, b T) T {
	return a + b
}

// FindPosition returns a function that checks if a given value x is equal to the specified position p.
// The returned function takes a single argument of type T and returns true if it matches p.
func FindPosition[T Number](p T) func(T) bool {
	return func(x T) bool {
		return x == p
	}
}
