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

func TestMapToAny(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	mapped := list.MapToAny(func(i int) any { return struct{ v int }{i * 2} })
	assert.Equal(t, List[any]{struct{ v int }{2}, struct{ v int }{4}, struct{ v int }{6}, struct{ v int }{8}, struct{ v int }{10}}, mapped)
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

func TestFilterMapToAny(t *testing.T) {
	list := Of(1, 2, 3, 4, 5)
	filtered := list.FilterMapToAny(func(i int) bool { return i%2 == 0 }, func(i int) any { return i * 2 })
	assert.Equal(t, List[any]{4, 8}, filtered)
}

func TestFlatMapToAny(t *testing.T) {
	list := List[List[int]]{{1, 2}, {2, 4}, {3, 6}}
	flattened := list.FlatMapToAny(func(s List[int]) Steam[any] {
		results := make(List[any], s.Length())
		for i, v := range s.Collect() {
			results[i] = fmt.Sprintf("v%v", v)
		}
		return results
	})
	assert.Equal(t, List[any]{"v1", "v2", "v2", "v4", "v3", "v6"}, flattened)
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

func TestPeek(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	var sum int
	peekedList := list.Peek(func(x int) {
		sum += x
	})
	assert.Equal(t, 15, sum, "Expected sum to be 15")
	assert.Equal(t, 5, peekedList.Length(), "Expected peekedList to have 5 elements")
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
	first, ok := list.FindFirst()
	assert.True(t, ok, "Expected to find the first element")
	assert.Equal(t, 1, *first, "Expected the first element to be 1")
}

func TestTakeWhile(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := list.TakeWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Length(), "Expected result to have 5 elements")
	for i, v := range result.Collect() {
		assert.Equal(t, i+1, v, "Expected element %d to be %d", i, i+1)
	}
}

func TestDropWhile(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := list.DropWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Length(), "Expected result to have 5 elements")
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
	index, ok := list.Position(func(x int) bool {
		return x == 3
	})
	assert.True(t, ok, "Expected to find the element")
	assert.Equal(t, 2, *index, "Expected the index to be 2")

	index, ok = list.Position(func(x int) bool {
		return x == 6
	})
	assert.False(t, ok, "Expected not to find the element")
	assert.Nil(t, index, "Expected the index to be nil")
}

func TestLast(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	last, ok := list.Last()
	assert.True(t, ok, "Expected to find the last element")
	assert.Equal(t, 5, *last, "Expected the last element to be 5")

	emptyList := List[int]{}
	last, ok = emptyList.Last()
	assert.False(t, ok, "Expected not to find the last element")
	assert.Nil(t, last, "Expected the last element to be nil")
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
	sorted := list.Sorted(func(a, b int) bool {
		return a < b
	})
	assert.Equal(t, List[int]{1, 2, 5, 8, 9}, sorted, "Expected the sorted list to be [1, 2, 5, 8, 9]")
}

func TestGetCompared(t *testing.T) {
	list := List[int]{5, 2, 8, 1, 9}
	max, ok := list.GetCompared(func(a, b int) bool {
		return a > b
	})
	assert.True(t, ok, "Expected to find the maximum element")
	assert.Equal(t, 9, *max, "Expected the maximum element to be 9")

	min, ok := list.GetCompared(func(a, b int) bool {
		return a < b
	})
	assert.True(t, ok, "Expected to find the minimum element")
	assert.Equal(t, 1, *min, "Expected the minimum element to be 1")

	emptyList := List[int]{}
	_, ok = emptyList.GetCompared(func(a, b int) bool {
		return a > b
	})
	assert.False(t, ok, "Expected not to find any element")
}