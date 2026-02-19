package steams

import (
	"fmt"
	"testing"

	"github.com/javiorfo/nilo"
	"github.com/stretchr/testify/assert"
)

func TestListFrom(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Collect())
}

func TestFilter(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	filtered := list.Filter(func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, filtered.Collect())

	var empty []int
	filtered = list.Filter(func(i int) bool { return i == 0 })
	assert.Equal(t, empty, filtered.Collect())
}

func TestMap(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	mapped := list.Map(func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6, 8, 10}, mapped.Collect())
}

func TestMapToString(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	mapped := list.MapToString(func(i int) string { return fmt.Sprintf("value: %d", i) })
	assert.Equal(t, []string{"value: 1", "value: 2", "value: 3", "value: 4", "value: 5"}, mapped.Collect())
}

func TestMapToInt(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	mapped := list.MapToInt(func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6, 8, 10}, mapped.Collect())
}

func TestFilterMap(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	filtered := list.FilterMap(func(i int) nilo.Option[int] {
		if i%2 == 0 {
			return nilo.Value(i * 2)
		}
		return nilo.Nil[int]()
	})
	assert.Equal(t, []int{4, 8}, filtered.Collect())
}

func TestFilterMapToInt(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	filtered := list.FilterMapToInt(func(i int) nilo.Option[int] {
		if i%2 == 0 {
			return nilo.Value(i * 2)
		}
		return nilo.Nil[int]()
	})
	assert.Equal(t, []int{4, 8}, filtered.Collect())
}

func TestFilterMapToString(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	filtered := list.FilterMapToString(func(i int) nilo.Option[string] {
		if i%2 == 0 {
			return nilo.Value(fmt.Sprintf("Res: %d", i*2))
		}
		return nilo.Nil[string]()
	})
	assert.Equal(t, []string{"Res: 4", "Res: 8"}, filtered.Collect())
}

func TestFlatMap(t *testing.T) {
	doubleMapper := func(x int) It[int] {
		return From(x, x)
	}
	input := From(1, 2, 3)
	expected := From(1, 1, 2, 2, 3, 3)

	flattened := input.FlatMap(doubleMapper)

	assert.Equal(t, expected.Collect(), flattened.Collect())
}

func TestFlatMapToInt(t *testing.T) {
	list := FromSlice([][]int{{1, 2}, {2, 4}, {3, 6}})
	flattened := list.FlatMapToInt(func(s []int) It[int] {
		results := make([]int, len(s))
		copy(results, s)
		return FromSlice(results)
	})
	assert.Equal(t, []int{1, 2, 2, 4, 3, 6}, flattened.Collect())
}

func TestFlatMapToString(t *testing.T) {
	list := FromSlice([][]int{{1, 2}, {2, 4}, {3, 6}})
	flattened := list.FlatMapToString(func(s []int) It[string] {
		results := make([]string, len(s))
		for i, v := range s {
			results[i] = fmt.Sprintf("v%v", v)
		}
		return FromSlice(results)
	})
	assert.Equal(t, []string{"v1", "v2", "v2", "v4", "v3", "v6"}, flattened.Collect())
}

func TestTake(t *testing.T) {
	limited := From(1, 2, 3, 4, 5).Take(3)
	assert.Equal(t, []int{1, 2, 3}, limited.Collect())
}

func TestCount(t *testing.T) {
	count := From(1, 2, 3, 4, 5).Count()
	assert.Equal(t, 5, count)
}

func TestForEach(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	var sum int
	list.ForEach(func(x int) {
		sum += x
	})
	assert.Equal(t, 15, sum, "Expected sum to be 15")
}

func TestForEachWithIndex(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	var sum int
	list.ForEachIdx(func(i, x int) {
		sum += x + 1
	})
	assert.Equal(t, 20, sum, "Expected sum to be 20")
}

func TestInspect(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	var sum int
	peekedList := list.Inspect(func(x int) {
		sum += x
	})
	assert.Equal(t, 15, sum, "Expected sum to be 15")
	assert.Equal(t, 5, peekedList.Count(), "Expected peekedList to have 5 elements")
}

func TestAll(t *testing.T) {
	list := From(2, 4, 6, 8, 10)
	assert.True(t, list.All(func(x int) bool {
		return x%2 == 0
	}), "Expected all elements to match the predicate")
}

func TestAny(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	assert.True(t, list.Any(func(x int) bool {
		return x%2 == 0
	}), "Expected at least one element to match the predicate")
}

func TestNone(t *testing.T) {
	list := From(1, 3, 5, 7, 9)
	assert.True(t, list.None(func(x int) bool {
		return x%2 == 0
	}), "Expected no elements to match the predicate")
}

func TestFirst(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	first := list.First()
	assert.True(t, first.IsValue(), "Expected to find the first element")
	assert.Equal(t, 1, first.AsValue(), "Expected the first element to be 1")
}

func TestFind(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	first := list.Find(func(n int) bool { return n > 1 && n < 4 })
	assert.True(t, first.IsValue(), "Expected to find one element")
	assert.Equal(t, 2, first.AsValue(), "Expected the element to be 2")
}

func TestTakeWhile(t *testing.T) {
	list := From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	result := list.TakeWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Count(), "Expected result to have 5 elements")
	for i, v := range result.Collect() {
		assert.Equal(t, i+1, v, "Expected element %d to be %d", i, i+1)
	}
}

func TestSkipWhile(t *testing.T) {
	list := From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	result := list.SkipWhile(func(x int) bool {
		return x < 6
	})
	assert.Equal(t, 5, result.Count(), "Expected result to have 5 elements")
	for i, v := range result.Collect() {
		assert.Equal(t, i+6, v, "Expected element %d to be %d", i, i+6)
	}
}

func TestFold(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	result := list.Fold(0, Sum)
	assert.Equal(t, 15, result, "Expected result to be 15")
}

func TestRFold(t *testing.T) {
	list := From(2, 3, 3)
	result := list.RFold(1, func(acc, x int) int {
		return acc * x
	})
	assert.Equal(t, 18, result, "Expected result to be 18")
}

func TestReverse(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	reversed := list.Reverse()
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reversed.Collect(), "Expected reversed list to be [5, 4, 3, 2, 1]")
}

func TestPosition(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	index := list.Position(func(x int) bool {
		return x == 3
	})
	assert.True(t, index.IsValue(), "Expected to find the element")
	assert.Equal(t, 2, index.AsValue(), "Expected the index to be 2")

	index = list.Position(FindPosition(6))
	assert.False(t, index.IsValue(), "Expected not to find the element")
	assert.Equal(t, -1, index.Or(-1), "Expected the index to be nil")
}

func TestRPosition(t *testing.T) {
	list := From(3, 1, 2, 3, 5)
	index := list.RPosition(func(x int) bool {
		return x == 3
	})
	assert.True(t, index.IsValue(), "Expected to find the element")
	assert.Equal(t, 3, index.AsValue(), "Expected the index to be 2")

	index = list.RPosition(FindPosition(6))
	assert.False(t, index.IsValue(), "Expected not to find the element")
	assert.Equal(t, -1, index.Or(-1), "Expected the index to be nil")
}

func TestLast(t *testing.T) {
	list := From(1, 2, 3, 4, 5)
	last := list.Last()
	assert.True(t, last.IsValue(), "Expected to find the last element")
	assert.Equal(t, 5, last.AsValue(), "Expected the last element to be 5")

	emptyList := From[int]()
	last = emptyList.Last()
	assert.True(t, last.IsNil(), "Expected not to find the last element")
	assert.Equal(t, 0, last.Or(0), "Expected the last element to be nil")
}

func TestSkip(t *testing.T) {
	list := From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	skipped := list.Skip(3)
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9, 10}, skipped.Collect(), "Expected the skipped list to be [4, 5, 6, 7, 8, 9, 10]")
	emptyList := From[int]()
	skipped = emptyList.Skip(3)
	assert.Equal(t, emptyList.Collect(), skipped.Collect(), "Expected the skipped list to be empty")
}

func TestSorted(t *testing.T) {
	list := From(5, 2, 8, 1, 9)
	sorted := list.SortBy(OrderDesc)
	assert.Equal(t, []int{1, 2, 5, 8, 9}, sorted.Collect(), "Expected the sorted list to be [1, 2, 5, 8, 9]")
}

func TestCompare(t *testing.T) {
	list := From(5, 2, 8, 1, 9)
	max := list.Compare(Max)
	assert.True(t, max.IsValue(), "Expected to find the maximum element")
	assert.Equal(t, 9, max.AsValue(), "Expected the maximum element to be 9")

	min := list.Compare(Min)
	assert.False(t, min.IsNil(), "Expected to find the minimum element")
	assert.Equal(t, 1, min.AsValue(), "Expected the minimum element to be 1")

	emptyList := From[int]()
	min = emptyList.Compare(Min)
	assert.False(t, min.IsValue(), "Expected not to find any element")
}

func TestChain(t *testing.T) {
	s1 := From(1, 2)
	s2 := From(3, 4)

	combined := s1.Chain(s2)

	expected := []int{1, 2, 3, 4}
	assert.Equal(t, expected, combined.Collect(), "The chained iterator should yield elements in order")
}

func TestNth(t *testing.T) {
	data := From("apple", "banana", "cherry")

	t.Run("Find valid index", func(t *testing.T) {
		val := data.Nth(1)
		assert.True(t, val.IsValue())
		assert.Equal(t, "banana", val.AsValue())
	})

	t.Run("Index out of bounds", func(t *testing.T) {
		val := data.Nth(10)
		assert.True(t, val.IsNil())
	})

	t.Run("Negative index", func(t *testing.T) {
		val := data.Nth(-1)
		assert.False(t, val.IsValue())
	})
}

func TestPartition(t *testing.T) {
	t.Run("Split integers by parity", func(t *testing.T) {
		input := From(10, 15, 20, 25)
		isEven := func(n int) bool { return n%2 == 0 }

		pos, neg := input.Partition(isEven)

		assert.Equal(t, []int{10, 20}, pos.Collect(), "Should contain even numbers")
		assert.Equal(t, []int{15, 25}, neg.Collect(), "Should contain odd numbers")
	})

	t.Run("Empty iterator", func(t *testing.T) {
		input := From[int]()
		alwaysTrue := func(n int) bool { return true }

		pos, neg := input.Partition(alwaysTrue)

		assert.Empty(t, pos.Collect())
		assert.Empty(t, neg.Collect())
	})

	t.Run("All elements in one side", func(t *testing.T) {
		input := From("a", "b", "c")
		startsWithA := func(s string) bool { return s == "a" }

		pos, neg := input.Partition(startsWithA)

		assert.Equal(t, pos.Count(), 1)
		assert.Equal(t, neg.Count(), 2)
		assert.Contains(t, pos.Collect(), "a")
		assert.NotContains(t, neg.Collect(), "a")
	})
}
