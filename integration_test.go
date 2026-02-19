package steams

import (
	"slices"
	"testing"

	"github.com/javiorfo/nilo"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDistinct(t *testing.T) {
	intSlice := From(1, 2, 3, 2, 4, 1, 5)
	distinctInts := Distinct(intSlice)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, distinctInts.Collect(), "Expected the distinct integers to be [1, 2, 3, 4, 5]")

	stringSlice := []string{"apple", "banana", "cherry", "banana", "date"}
	distinctStrings := Distinct(FromSlice(stringSlice))
	assert.Equal(t, []string{"apple", "banana", "cherry", "date"}, distinctStrings.Collect(), "Expected the distinct strings to be [apple, banana, cherry, date]")

	var emptySlice []int
	distinctEmpty := Distinct(FromSlice(emptySlice))
	assert.Equal(t, emptySlice, distinctEmpty.Collect(), "Expected the distinct elements of an empty slice to be an empty slice")
}

func TestIntegrationMap(t *testing.T) {
	t.Run("Standard Transform", func(t *testing.T) {
		nums := From(1, 2, 3)

		squared := Map(nums, func(n int) int {
			return n * n
		})

		result := squared.Collect()
		expected := []int{1, 4, 9}

		if !slices.Equal(result, expected) {
			t.Errorf("Map failed: expected %v, got %v", expected, result)
		}
	})

	t.Run("Type Transformation", func(t *testing.T) {
		nums := From(1, 2)

		asString := Map(nums, func(n int) string {
			if n == 1 {
				return "one"
			}
			return "two"
		})

		result := asString.Collect()
		expected := []string{"one", "two"}

		if !slices.Equal(result, expected) {
			t.Errorf("Map type conversion failed: expected %v, got %v", expected, result)
		}
	})
}

func TestIntegrationFlatMap(t *testing.T) {
	t.Run("Expand Elements", func(t *testing.T) {
		words := From("Go", "it")

		chars := FlatMap(words, func(s string) It[rune] {
			return func(yield func(rune) bool) {
				for _, r := range s {
					if !yield(r) {
						return
					}
				}
			}
		})

		result := chars.Collect()
		expected := []rune{'G', 'o', 'i', 't'}

		if !slices.Equal(result, expected) {
			t.Errorf("FlatMap failed: expected %v, got %v", expected, result)
		}
	})

	t.Run("Filtering via FlatMap", func(t *testing.T) {
		nums := From(1, 2, 3, 4)

		evensOnly := FlatMap(nums, func(n int) It[int] {
			return func(yield func(int) bool) {
				if n%2 == 0 {
					yield(n)
				}
			}
		})

		result := evensOnly.Collect()
		expected := []int{2, 4}

		if !slices.Equal(result, expected) {
			t.Errorf("FlatMap filtering failed: expected %v, got %v", expected, result)
		}
	})
}

func TestIntegrationFold(t *testing.T) {
	nums := From(1, 2, 3, 4)
	sum := Fold(nums, 0, func(acc, v int) int {
		return acc + v
	})

	if sum != 10 {
		t.Errorf("Fold sum failed: expected 10, got %d", sum)
	}
}

func TestIntegrationRFold(t *testing.T) {
	chars := From("a", "b", "c")
	result := RFold(chars, "!", func(v, acc string) string {
		return v + acc
	})

	expected := "abc!"
	if result != expected {
		t.Errorf("RFold failed: expected %s, got %s", expected, result)
	}
}

func TestIntegrationFlatten(t *testing.T) {
	t.Run("Standard Flatten", func(t *testing.T) {
		nested := From(nilo.Value(1).Iter(), nilo.Value(2).Iter(), nilo.Nil[int]().Iter())

		flat := Flatten(nested)
		result := flat.Collect()

		expected := []int{1, 2}
		if !slices.Equal(result, expected) {
			t.Errorf("Flatten failed: expected %v, got %v", expected, result)
		}
	})

	t.Run("Early Termination", func(t *testing.T) {
		inner := From(1, 2, 3, 4, 5).AsSeq()
		nested := From(inner)

		count := 0
		for v := range Flatten(nested) {
			count++
			if v == 2 {
				break
			}
		}

		if count != 2 {
			t.Errorf("Flatten did not respect break: processed %d elements", count)
		}
	})
}

func TestIntegrationGroupBy(t *testing.T) {
	people := From("Alice", "Bob", "Brian", "Charlie", "David", "Eve", "Eddie")

	classifier := func(name string) string {
		return string(name[0])
	}

	expected := map[string][]string{
		"A": {"Alice"},
		"B": {"Bob", "Brian"},
		"C": {"Charlie"},
		"D": {"David"},
		"E": {"Eve", "Eddie"},
	}

	grouped := GroupBy(people, classifier).Collect()

	assert.Equal(t, expected["A"], grouped["A"].Collect())
	assert.Equal(t, len(expected["B"]), grouped["B"].Count())
	assert.Equal(t, 5, len(grouped))

	groupedByCounting := GroupByCounting(people, classifier).Collect()

	assert.Equal(t, 1, groupedByCounting["A"])
	assert.Equal(t, 2, groupedByCounting["B"])
	assert.Equal(t, 1, groupedByCounting["C"])
	assert.Equal(t, 1, groupedByCounting["D"])
	assert.Equal(t, 2, groupedByCounting["E"])

}

func TestIntegrationZip(t *testing.T) {
	s1 := From(1, 2, 3)
	s2 := From("a", "b", "c")
	expected := []struct {
		First  int
		Second string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	assert.Equal(t, expected, Zip(s1, s2).Collect())

	s1 = From(1, 2, 3, 4)
	s2 = From("a", "b", "c")
	assert.Equal(t, expected, Zip(s1, s2).Collect())
}

func TestIntegrationCollectItToIt2(t *testing.T) {
	s := From(1, 2, 3, 4, 5)
	result := CollectItToIt2(s, func(i int) int { return i }, func(i int) int { return i * 2 }).
		SortBy(func(a, b int) bool { return a < b })
	assert.Equal(t, map[int]int{1: 2, 2: 4, 3: 6, 4: 8, 5: 10}, result.Collect())

	s2 := From("a", "b", "c")
	result2 := CollectItToIt2(s2, func(i string) string { return i }, func(i string) int { return len(i) }).
		SortBy(func(a, b string) bool { return a < b })
	assert.Equal(t, map[string]int{"a": 1, "b": 1, "c": 1}, result2.Collect())
}

func TestIntegrationCollectIt2ToIt(t *testing.T) {
	s := FromMap(map[int]string{1: "one", 2: "two", 3: "three"})
	result := CollectIt2ToIt(s, func(k int, v string) string { return v })
	assert.Equal(t, []string{"one", "three", "two"}, result.SortBy(OrderDesc).Collect())

	s2 := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	result2 := CollectIt2ToIt(s2, func(k string, v int) int { return v * 2 })
	assert.Equal(t, []int{2, 4, 6}, result2.SortBy(OrderDesc).Collect())
}

func TestIntegrationChainAll(t *testing.T) {
	tests := []struct {
		name     string
		input    []It[int]
		expected []int
	}{
		{
			name: "Three sequences",
			input: []It[int]{
				From(1),
				From(2),
				From(3),
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty input",
			input:    []It[int]{},
			expected: nil,
		},
		{
			name: "Mixed empty and full",
			input: []It[int]{
				From[int](),
				From(10),
				From[int](),
			},
			expected: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combined := ChainAll(tt.input...)
			assert.Equal(t, tt.expected, combined.Collect())
		})
	}
}
