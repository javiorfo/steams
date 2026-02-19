package steams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFilter(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	})

	filtered := m.Filter(func(k int, v string) bool {
		return k%2 == 0
	})

	assert.Equal(t, map[int]string{
		2: "two",
		4: "four",
	}, filtered.Collect(), "Expected the filtered map to be {2: 'two', 4: 'four'}")
}

func TestMapMap(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	mapped := m.Map(func(k int, v string) (int, string) {
		switch v {
		case "one":
			return k, "1"
		case "two":
			return k, "2"
		case "three":
			return k, "3"
		default:
			return k, ""
		}
	}).SortBy(func(a, b int) bool { return a < b })

	assert.Equal(t, map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}, mapped.Collect(), "Expected the mapped map to contain Pair values")
}

func TestMapMapToString(t *testing.T) {
	m := FromMap(map[int]int{
		1: 10,
		2: 20,
		3: 30,
	})

	mapped := m.MapToString(func(k int, v int) (int, string) {
		return k, fmt.Sprint(v)
	}).SortBy(func(a, b int) bool { return a < b })

	assert.Equal(t, map[int]string{
		1: "10",
		2: "20",
		3: "30",
	}, mapped.Collect(), "Expected the mapped map to contain string values")
}

func TestMapMapToInt(t *testing.T) {
	m := FromMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	mapped := m.MapToInt(func(k string, v int) (string, int) {
		return k, v * 2
	}).SortBy(func(a, b string) bool { return a < b })

	assert.Equal(t, map[string]int{
		"one":   2,
		"three": 6,
		"two":   4,
	}, mapped.Collect(), "Expected the mapped map to contain doubled integer values")
}

func TestMapForEach(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	var result []Entry[int, string]
	m.ForEach(func(k int, v string) {
		result = append(result, Entry[int, string]{k, v})
	})

	assert.Equal(t, 3, len(result), "Expected the result to contain all key-value pairs")
}

func TestMapInspect(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	var result []Entry[int, string]
	peeked := m.Inspect(func(k int, v string) {
		result = append(result, Entry[int, string]{k, v})
	})

	assert.Equal(t, 3, len(result), "Expected the result to contain all key-value pairs")
	assert.Equal(t, m.Collect(), peeked.Collect(), "Expected the peeked map to be the same as the original map")
}

func TestMapTake(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	})

	limited := m.Take(3)
	assert.Equal(t, 3, limited.Count(), "Expected the limited map to contain the first 3 key-value pairs")
}

func TestMapCount(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})
	assert.Equal(t, 3, m.Count(), "Expected the map to have 3 elements")

	emptyMap := FromMap(map[int]string{})
	assert.Equal(t, 0, emptyMap.Count(), "Expected the empty map to have 0 elements")
}

func TestValues(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})
	values := m.Values().SortBy(OrderDesc)
	assert.Equal(t, []string{"one", "three", "two"}, values.Collect(), "Expected the values stream to contain all the values")
}

func TestKeys(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})
	keys := m.Keys().SortBy(OrderDesc)
	assert.Equal(t, []int{1, 2, 3}, keys.Collect(), "Expected the keys stream to contain all the keys")
}

func TestMapAll(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	assert.True(t, m.All(func(k int, v string) bool {
		return len(v) > 2
	}), "Expected all key-value pairs to match the predicate")

	assert.False(t, m.All(func(k int, v string) bool {
		return len(v) > 3
	}), "Expected not all key-value pairs to match the predicate")
}

func TestMapAny(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	assert.True(t, m.Any(func(k int, v string) bool {
		return len(v) == 3
	}), "Expected at least one key-value pair to match the predicate")

	assert.False(t, m.Any(func(k int, v string) bool {
		return len(v) == 4
	}), "Expected no key-value pairs to match the predicate")
}

func TestMapNone(t *testing.T) {
	m := FromMap(map[int]string{
		1: "one",
		2: "two",
		3: "three",
	})

	assert.True(t, m.None(func(k int, v string) bool {
		return len(v) == 4
	}), "Expected no key-value pairs to match the predicate")

	assert.False(t, m.None(func(k int, v string) bool {
		return len(v) == 3
	}), "Expected at least one key-value pair to match the predicate")
}

func TestMapSorted(t *testing.T) {
	m := FromMap(map[int]string{
		5: "five",
		2: "two",
		8: "eight",
		1: "one",
		9: "nine",
	})

	sorted := m.SortBy(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, map[int]string{
		1: "one",
		2: "two",
		5: "five",
		8: "eight",
		9: "nine",
	}, sorted.Collect(), "Expected the sorted map to have the keys in ascending order")
}

func TestMapCompare(t *testing.T) {
	m := FromMap(map[int]string{
		5: "five",
		2: "two",
		8: "eight",
		1: "one",
		9: "nine",
	})

	max := m.Compare(func(a, b int) bool {
		return a > b
	})
	assert.True(t, max.IsValue(), "Expected to find the maximum key-value pair")
	assert.Equal(t, Entry[int, string]{9, "nine"}, max.AsValue(), "Expected the maximum key-value pair to be {9, 'nine'}")

	min := m.Compare(func(a, b int) bool {
		return a < b
	})
	assert.True(t, min.IsValue(), "Expected to find the minimum key-value pair")
	assert.Equal(t, Entry[int, string]{1, "one"}, min.AsValue(), "Expected the minimum key-value pair to be {1, 'one'}")

	emptyMap := FromMap(map[int]string{})
	min = emptyMap.Compare(func(a, b int) bool {
		return a > b
	})
	assert.False(t, min.IsValue(), "Expected not to find any key-value pair")
}
