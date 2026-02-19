package steams

func Fold[T any, R any](i it[T], initial R, accumulator func(R, T) R) R {
	result := initial
	for v := range i {
		result = accumulator(result, v)
	}
	return result
}

func RFold[T any, R any](i it[T], initial R, accumulator func(T, R) R) R {
	var elements []T = i.Collect()
	result := initial
	for i := len(elements) - 1; i >= 0; i-- {
		result = accumulator(elements[i], result)
	}

	return result
}
