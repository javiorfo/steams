package opt

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) Get() T {
	return *o.value
}

func (o Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.Get()
	} else {
		return other
	}
}

func (o Optional[T]) Or(supplier func() Optional[T]) Optional[T] {
	if o.IsEmpty() {
		return supplier()
	} else {
		return Of(o.Get())
	}
}

func (o Optional[T]) Filter(filter func(T) bool) Optional[T] {
	if o.IsPresent() && filter(o.Get()) {
		return Of(o.Get())
	} else {
		return Empty[T]()
	}
}

func (o Optional[T]) MapToAny(mapper func(T) any) Optional[any] {
    if o.IsPresent() {
	    return Of(mapper(o.Get()))
    }
    return Empty[any]()
}

func (o Optional[T]) IsEmpty() bool {
	if o.value == nil {
		return true
	}
	return false
}

func (o Optional[T]) IsPresent() bool {
	if o.value != nil {
		return true
	}
	return false
}

func (o Optional[T]) IfPresent(consumer func(T)) {
	if o.IsPresent() {
		consumer(o.Get())
	}
}

func (o Optional[T]) IfPresentOrElse(consumer func(T), or func()) {
	if o.IsPresent() {
		consumer(o.Get())
	} else {
		or()
	}
}

func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return o.Get()
	} else {
		return supplier()
	}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{&value}
}

func OfNullable[T any](value *T) Optional[T] {
	return Optional[T]{value}
}

func Map[T, R any](o Optional[T], mapper func(T) R) Optional[R] {
    if o.IsPresent() {
	    return Of(mapper(o.Get()))
    }
    return Empty[R]()
}
