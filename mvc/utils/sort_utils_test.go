package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	// Initialization:
	els := []int{9,8,7,6,5}

	// Execution:
	Sort(els)

	// Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func getElements(n int)[]int{
	result := make([]int, n)
	i := 0
	for j := n-1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

// BenchmarkBubbleSort10-12 => 151920348 => 7.82 ns/op
func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i:= 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

// BenchmarkSort10-12 => 14.673.352 => 80.7 ns/op
func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i:= 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

// BenchmarkBubbleSort1000-12 => 2.376.387 => 490 ns/op
func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i:= 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

// BenchmarkSort1000-12 => 36.918 => 31.818 ns/op
func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i:= 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

// BenchmarkBubbleSort100000-12 => 1 => 7.612.221.100 ns/op
func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i:= 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

// BenchmarkSort100000-12 => 232 => 5.055.562 ns/op
func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i:= 0; i < b.N; i++ {
		sort.Ints(els)
	}
}


