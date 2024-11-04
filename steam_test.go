package steams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	intSlice := List[int]{1, 2, 3, 2, 4, 1, 5}
	distinctInts := Distinct(intSlice)
	assert.Equal(t, List[int]{1, 2, 3, 4, 5}, distinctInts, "Expected the distinct integers to be [1, 2, 3, 4, 5]")

	stringSlice := List[string]{"apple", "banana", "cherry", "banana", "date"}
	distinctStrings := Distinct(stringSlice)
	assert.Equal(t, List[string]{"apple", "banana", "cherry", "date"}, distinctStrings, "Expected the distinct strings to be [apple, banana, cherry, date]")

	emptySlice := List[int]{}
	distinctEmpty := Distinct(emptySlice)
	assert.Equal(t, List[int]{}, distinctEmpty, "Expected the distinct elements of an empty slice to be an empty slice")
}

func TestCollectSteamToSteam2(t *testing.T) {
	s := List[int]{1, 2, 3, 4, 5}
	result := CollectSteamToSteam2(s, func(i int) int { return i }, func(i int) int { return i * 2 }).
		Sorted(func(a, b int) bool { return a < b })
	assert.Equal(t, Map[int, int]{1: 2, 2: 4, 3: 6, 4: 8, 5: 10}, result)

	s2 := List[string]{"a", "b", "c"}
	result2 := CollectSteamToSteam2(s2, func(i string) string { return i }, func(i string) int { return len(i) }).
		Sorted(func(a, b string) bool { return a < b })
	assert.Equal(t, Map[string, int]{"a": 1, "b": 1, "c": 1}, result2)
}

func TestCollectSteam2ToSteam(t *testing.T) {
	s := Map[int, string]{1: "one", 2: "two", 3: "three"}
	result := CollectSteam2ToSteam(Map[int, string](s), func(k int, v string) string { return v })
	assert.Equal(t, List[string]{"one", "three", "two"}, result.Sorted(OrderDesc))

	s2 := Map[string, int]{"a": 1, "b": 2, "c": 3}
	result2 := CollectSteam2ToSteam(s2, func(k string, v int) int { return v * 2 })
	assert.Equal(t, List[int]{2, 4, 6}, result2.Sorted(OrderDesc))
}

func TestZip(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := Of("a", "b", "c")
	result := Zip(s1, s2)
	expected := List[struct {
		first  int
		second string
	}]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	assert.Equal(t, expected, result)

	s1 = List[int]{1, 2, 3, 4}
	s2 = List[string]{"a", "b", "c"}
	assert.Panics(t, func() {
		Zip(s1, s2)
	})
}
