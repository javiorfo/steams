package steams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOf(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	assert.Equal(t, List[int]{1, 2, 3, 4, 5}, list)
}

func TestFilter(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	filtered := list.Filter(func(i int) bool { return i%2 == 0 })
	assert.Equal(t, List[int]{2, 4}, filtered)
}

func TestMap(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	mapped := list.Map(func(i int) int { return i * 2 })
	assert.Equal(t, List[int]{2, 4, 6, 8, 10}, mapped)
}

func TestMapToString(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	mapped := list.MapToString(func(i int) string { return fmt.Sprintf("value: %d", i) })
	assert.Equal(t, List[string]{"value: 1", "value: 2", "value: 3", "value: 4", "value: 5"}, mapped)
}

func TestMapToInt(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	mapped := list.MapToInt(func(i int) int { return i * 2 })
	assert.Equal(t, List[int]{2, 4, 6, 8, 10}, mapped)
}

func TestFilterMap(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	filtered := list.FilterMap(func(i int) bool { return i%2 == 0 }, func(i int) int { return i * 2 })
	assert.Equal(t, List[int]{4, 8}, filtered)
}

func TestFilterMapToInt(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	filtered := list.FilterMapToInt(func(i int) bool { return i%2 == 0 }, func(i int) int { return i * 2 })
	assert.Equal(t, List[int]{4, 8}, filtered)
}

func TestFilterMapToString(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	filtered := list.FilterMapToString(func(i int) bool { return i%2 == 0 }, func(i int) string { return fmt.Sprintf("Res: %d", i*2) })
	assert.Equal(t, List[string]{"Res: 4", "Res: 8"}, filtered)
}

func TestFlatMap(t *testing.T) {
	doubleMapper := func(x int) Steam[int] {
		return List[int]{x, x}
	}
	input := List[int]{1, 2, 3}
	expected := List[int]{1, 1, 2, 2, 3, 3}

	flattened := input.FlatMap(doubleMapper)

	assert.Equal(t, expected, flattened)
}

func TestFlatMapToInt(t *testing.T) {
	list := List[List[int]]{{1, 2}, {2, 4}, {3, 6}}
	flattened := list.FlatMapToInt(func(s List[int]) Steam[int] {
		results := make(List[int], s.Count())
		for i, v := range s.Collect() {
			results[i] = v
		}
		return results
	})
	assert.Equal(t, List[int]{1, 2, 2, 4, 3, 6}, flattened)
}

func TestFlatMapToString(t *testing.T) {
	list := List[List[int]]{{1, 2}, {2, 4}, {3, 6}}
	flattened := list.FlatMapToString(func(s List[int]) Steam[string] {
		results := make(List[string], s.Count())
		for i, v := range s.Collect() {
			results[i] = fmt.Sprintf("v%v", v)
		}
		return results
	})
	assert.Equal(t, List[string]{"v1", "v2", "v2", "v4", "v3", "v6"}, flattened)
}

func TestLimit(t *testing.T) {
	limited := Of(1, 2, 3, 4, 5).Limit(3)
	assert.Equal(t, List[int]{1, 2, 3}, limited)
}

func TestCount(t *testing.T) {
	count := Of(1, 2, 3, 4, 5).Count()
	assert.Equal(t, 5, count)
}

func TestForEach(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	var sum int
	list.ForEach(func(x int) {
		sum += x
	})
	assert.Equal(t, 15, sum, "Expected sum to be 15")
}

func TestForEachWithIndex(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	var sum int
	list.ForEachWithIndex(func(i, x int) {
		sum += x + 1
	})
	assert.Equal(t, 20, sum, "Expected sum to be 20")
}

func TestPeek(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	var sum int
	peekedList := list.Peek(func(x int) {
		sum += x
	})
	assert.Equal(t, 15, sum, "Expected sum to be 15")
	assert.Equal(t, 5, peekedList.Count(), "Expected peekedList to have 5 elements")
}

func TestAllMatch(t *testing.T) {
	list := List[int]{2, 4, 6, 8, 10}
	assert.True(t, list.AllMatch(func(x int) bool {
		return x%2 == 0
	}), "Expected all elements to match the predicate")
}

func TestAnyMatch(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	assert.True(t, list.AnyMatch(func(x int) bool {
		return x%2 == 0
	}), "Expected at least one element to match the predicate")
}

func TestNoneMatch(t *testing.T) {
	list := List[int]{1, 3, 5, 7, 9}
	assert.True(t, list.NoneMatch(func(x int) bool {
		return x%2 == 0
	}), "Expected no elements to match the predicate")
}

func TestFindFirst(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	first := list.FindFirst()
	assert.True(t, first.IsValue(), "Expected to find the first element")
	assert.Equal(t, 1, first.AsValue(), "Expected the first element to be 1")
}

func TestFindOne(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	first := list.FindOne(func(n int) bool { return n > 1 && n < 4 })
	assert.True(t, first.IsValue(), "Expected to find one element")
	assert.Equal(t, 2, first.AsValue(), "Expected the element to be 2")
}

func TestTakeWhile(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := list.TakeWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Count(), "Expected result to have 5 elements")
	for i, v := range result.Collect() {
		assert.Equal(t, i+1, v, "Expected element %d to be %d", i, i+1)
	}
}

func TestDropWhile(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := list.DropWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Count(), "Expected result to have 5 elements")
	for i, v := range result.Collect() {
		assert.Equal(t, i+6, v, "Expected element %d to be %d", i, i+6)
	}
}

func TestReduce(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	result := list.Reduce(0, func(acc, x int) int {
		return acc + x
	})
	assert.Equal(t, 15, result, "Expected result to be 15")
}

func TestReverse(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	reversed := list.Reverse()
	assert.Equal(t, List[int]{5, 4, 3, 2, 1}, reversed, "Expected reversed list to be [5, 4, 3, 2, 1]")
}

func TestPosition(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	index := list.Position(func(x int) bool {
		return x == 3
	})
	assert.True(t, index.IsValue(), "Expected to find the element")
	assert.Equal(t, 2, index.AsValue(), "Expected the index to be 2")

	index = list.Position(FindPosition(6))
	assert.False(t, index.IsValue(), "Expected not to find the element")
	assert.Equal(t, -1, index.Or(-1), "Expected the index to be nil")
}

func TestLast(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	last := list.Last()
	assert.True(t, last.IsValue(), "Expected to find the last element")
	assert.Equal(t, 5, last.AsValue(), "Expected the last element to be 5")

	emptyList := List[int]{}
	last = emptyList.Last()
	assert.True(t, last.IsNil(), "Expected not to find the last element")
	assert.Equal(t, 0, last.Or(0), "Expected the last element to be nil")
}

func TestSkip(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	skipped := list.Skip(3)
	assert.Equal(t, List[int]{4, 5, 6, 7, 8, 9, 10}, skipped, "Expected the skipped list to be [4, 5, 6, 7, 8, 9, 10]")
	emptyList := List[int]{}
	skipped = emptyList.Skip(3)
	assert.Equal(t, emptyList, skipped, "Expected the skipped list to be empty")
}

func TestSorted(t *testing.T) {
	list := List[int]{5, 2, 8, 1, 9}
	sorted := list.Sorted(OrderDesc)
	assert.Equal(t, List[int]{1, 2, 5, 8, 9}, sorted, "Expected the sorted list to be [1, 2, 5, 8, 9]")
}

func TestGetCompared(t *testing.T) {
	list := List[int]{5, 2, 8, 1, 9}
	max := list.GetCompared(func(a, b int) bool {
		return a > b
	})
	assert.True(t, max.IsValue(), "Expected to find the maximum element")
	assert.Equal(t, 9, max.AsValue(), "Expected the maximum element to be 9")

	min := list.GetCompared(Max)
	assert.False(t, min.IsNil(), "Expected to find the minimum element")
	assert.Equal(t, 1, min.AsValue(), "Expected the minimum element to be 1")

	emptyList := List[int]{}
	min = emptyList.GetCompared(Min)
	assert.False(t, min.IsValue(), "Expected not to find any element")
}
