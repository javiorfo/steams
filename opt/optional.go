// Package opt provides a generic Optional type that can be used to represent a value that may or may not be present.
package opt

// Optional is a generic type that encapsulates a value that may or may not be present.
// It provides methods to work with the value safely.
type Optional[T any] struct {
	value *T
}

// Get retrieves the value contained in the Optional.
// It panics if the value is not present (i.e., if IsEmpty() returns true).
func (o Optional[T]) Get() T {
	return *o.value
}

// OrElse returns the value contained in the Optional if present; otherwise, it returns the provided alternative value.
func (o Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.Get()
	} else {
		return other
	}
}

// OrErr returns the value contained in the Optional if present; otherwise, it returns the provided error.
// This returns pointer type or an error
func (o Optional[T]) OrErr(err error) (*T, error) {
	if o.IsPresent() {
		return o.value, nil
	} else {
		return nil, err
	}
}

// Or returns the value contained in the Optional if present; otherwise, it invokes the provided supplier function
// to obtain a new Optional.
func (o Optional[T]) Or(supplier func() Optional[T]) Optional[T] {
	if o.IsEmpty() {
		return supplier()
	} else {
		return Of(o.Get())
	}
}

// Filter returns an Optional containing the value if it is present and satisfies the provided filter function;
// otherwise, it returns an empty Optional.
func (o Optional[T]) Filter(filter func(T) bool) Optional[T] {
	if o.IsPresent() && filter(o.Get()) {
		return Of(o.Get())
	} else {
		return Empty[T]()
	}
}

// MapToAny applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped value; otherwise, it returns an empty Optional.
func (o Optional[T]) MapToAny(mapper func(T) any) Optional[any] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[any]()
}

// IsEmpty returns true if the Optional does not contain a value; otherwise, it returns false.
func (o Optional[T]) IsEmpty() bool {
	return o.value == nil
}

// IsPresent returns true if the Optional contains a value; otherwise, it returns false.
func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

// IfPresent executes the provided consumer function with the value contained in the Optional if present.
func (o Optional[T]) IfPresent(consumer func(T)) {
	if o.IsPresent() {
		consumer(o.Get())
	}
}

// IfPresentOrElse executes the provided consumer function with the value if present; otherwise, it executes the provided alternative function.
func (o Optional[T]) IfPresentOrElse(consumer func(T), or func()) {
	if o.IsPresent() {
		consumer(o.Get())
	} else {
		or()
	}
}

// OrElseGet returns the value contained in the Optional if present; otherwise, it invokes the provided supplier function to obtain the value.
func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return o.Get()
	} else {
		return supplier()
	}
}

// Empty returns an empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// Of creates an Optional containing the provided value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{&value}
}

// OfNullable creates an Optional from a pointer to a value.
// If the pointer is nil, it returns an empty Optional.
func OfNullable[T any](value *T) Optional[T] {
	return Optional[T]{value}
}

// Map applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped value; otherwise, it returns an empty Optional.
func Map[T, R any](o Optional[T], mapper func(T) R) Optional[R] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[R]()
}
