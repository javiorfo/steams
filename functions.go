package steams

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type OrderedStruct[T any] interface {
	Compare(other T) int
}

func OrderStructDesc[T OrderedStruct[T]](a, b T) bool {
	return a.Compare(b) < 0
}

func OrderStructAsc[T OrderedStruct[T]](a, b T) bool {
	return a.Compare(b) > 0
}

func OrderDesc[T Ordered](a, b T) bool {
	return a < b
}

func OrderAsc[T Ordered](a, b T) bool {
	return a > b
}
