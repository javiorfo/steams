package steams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFilter(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	filtered := m.Filter(func(k int, v string) bool {
		return k%2 == 0
	})

	assert.Equal(t, Map[int, string]{
		2: "two",
		4: "four",
	}, filtered, "Expected the filtered map to be {2: 'two', 4: 'four'}")
}

func TestMapMap(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	mapped := m.Map(func(k int, v string) string {
		switch v {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		default:
			return ""
		}
	}).Sorted(func(a, b int) bool { return a < b })

	assert.Equal(t, Map[int, string]{
		1: "1",
		2: "2",
		3: "3",
	}, mapped, "Expected the mapped map to contain Pair values")
}

func TestMapMapToString(t *testing.T) {
	m := Map[int, int]{
		1: 10,
		2: 20,
		3: 30,
	}

	mapped := m.MapToString(func(k int, v int) string {
		return fmt.Sprint(v)
	}).Sorted(func(a, b int) bool { return a < b })

	assert.Equal(t, Map[int, string]{
		1: "10",
		2: "20",
		3: "30",
	}, mapped, "Expected the mapped map to contain string values")
}

func TestMapMapToInt(t *testing.T) {
	m := Map[string, int]{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	mapped := m.MapToInt(func(k string, v int) int {
		return v * 2
	}).Sorted(func(a, b string) bool { return a < b })

	assert.Equal(t, Map[string, int]{
		"one":   2,
		"three": 6,
		"two":   4,
	}, mapped, "Expected the mapped map to contain doubled integer values")
}

func TestMapFilterMap(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	filtered := m.FilterMap(func(k int, v string) bool {
		return k%2 == 0
	}, func(k int, v string) string {
		return v
	})

	assert.Equal(t, Map[int, string]{
		2: "two",
		4: "four",
	}, filtered, "Expected the filtered and mapped map to contain Pair values for even keys")
}

func TestMapFilterMapToInt(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	filtered := m.FilterMapToInt(func(k int, v string) bool {
		return k%2 == 0
	}, func(k int, v string) int {
		return k
	})

	assert.Equal(t, Map[int, int]{
		2: 2,
		4: 4,
	}, filtered, "Expected the filtered and mapped map to contain Pair values for even keys")
}

func TestMapForEach(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	var result []Pair[int, string]
	m.ForEach(func(k int, v string) {
		result = append(result, Pair[int, string]{k, v})
	})

	assert.Equal(t, 3, len(result), "Expected the result to contain all key-value pairs")
}

func TestMapPeek(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	var result []Pair[int, string]
	peeked := m.Peek(func(k int, v string) {
		result = append(result, Pair[int, string]{k, v})
	})

	assert.Equal(t, 3, len(result), "Expected the result to contain all key-value pairs")
	assert.Equal(t, m, peeked, "Expected the peeked map to be the same as the original map")
}

func TestMapLimit(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	limited := m.Limit(3)
	assert.Equal(t, 3, limited.Count(), "Expected the limited map to contain the first 3 key-value pairs")
}

func TestMapCount(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}
	assert.Equal(t, 3, m.Count(), "Expected the map to have 3 elements")

	emptyMap := Map[int, string]{}
	assert.Equal(t, 0, emptyMap.Count(), "Expected the empty map to have 0 elements")
}

func TestValuesToSteam(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}
	values := m.ValuesToSteam().Sorted(func(a, b string) bool { return a < b })
	assert.Equal(t, List[string]{"one", "three", "two"}, values, "Expected the values stream to contain all the values")
}

func TestKeysToSteam(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}
	keys := m.KeysToSteam().Sorted(func(a, b int) bool { return a < b })
	assert.Equal(t, List[int]{1, 2, 3}, keys, "Expected the keys stream to contain all the keys")
}

func TestToAnySteam(t *testing.T) {
	m := OfMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})
	stream := m.ToAnySteam(func(k int, v string) any {
		return Pair[int, string]{k, v}
	})
	assert.Equal(t, 3, stream.Count(), "Expected the stream to contain all the key-value pairs as Pair values")
}

func TestMapAllMatch(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.True(t, m.AllMatch(func(k int, v string) bool {
		return len(v) > 2
	}), "Expected all key-value pairs to match the predicate")

	assert.False(t, m.AllMatch(func(k int, v string) bool {
		return len(v) > 3
	}), "Expected not all key-value pairs to match the predicate")
}

func TestMapAnyMatch(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.True(t, m.AnyMatch(func(k int, v string) bool {
		return len(v) == 3
	}), "Expected at least one key-value pair to match the predicate")

	assert.False(t, m.AnyMatch(func(k int, v string) bool {
		return len(v) == 4
	}), "Expected no key-value pairs to match the predicate")
}

func TestMapNoneMatch(t *testing.T) {
	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.True(t, m.NoneMatch(func(k int, v string) bool {
		return len(v) == 4
	}), "Expected no key-value pairs to match the predicate")

	assert.False(t, m.NoneMatch(func(k int, v string) bool {
		return len(v) == 3
	}), "Expected at least one key-value pair to match the predicate")
}

func TestMapSorted(t *testing.T) {
	m := Map[int, string]{
		5: "five",
		2: "two",
		8: "eight",
		1: "one",
		9: "nine",
	}

	sorted := m.Sorted(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, Map[int, string]{
		1: "one",
		2: "two",
		5: "five",
		8: "eight",
		9: "nine",
	}, sorted, "Expected the sorted map to have the keys in ascending order")
}

func TestMapGetCompared(t *testing.T) {
	m := Map[int, string]{
		5: "five",
		2: "two",
		8: "eight",
		1: "one",
		9: "nine",
	}

	max := m.GetCompared(func(a, b int) bool {
		return a > b
	})
	assert.True(t, max.IsSome(), "Expected to find the maximum key-value pair")
	assert.Equal(t, Pair[int, string]{9, "nine"}, max.Unwrap(), "Expected the maximum key-value pair to be {9, 'nine'}")

	min := m.GetCompared(func(a, b int) bool {
		return a < b
	})
	assert.True(t, min.IsSome(), "Expected to find the minimum key-value pair")
	assert.Equal(t, Pair[int, string]{1, "one"}, min.Unwrap(), "Expected the minimum key-value pair to be {1, 'one'}")

	emptyMap := Map[int, string]{}
	min = emptyMap.GetCompared(func(a, b int) bool {
		return a > b
	})
	assert.False(t, min.IsSome(), "Expected not to find any key-value pair")
}
