package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {

	els := []int{5, 6, 7, 8, 9} // best case
	els = Sort(els)

	assert.NotNil(t, els)              // check if slice is nil
	assert.EqualValues(t, 5, len(els)) // check if length of input and output match
	assert.EqualValues(t, 5, els[0])   // check each index
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {

	els := []int{9, 8, 7, 6, 5}
	els = Sort(els)

	assert.NotNil(t, els)              // check if slice is nil
	assert.EqualValues(t, 5, len(els)) // check if length of input and output match
	assert.EqualValues(t, 5, els[0])   // check each index
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortNilSlice(t *testing.T) {

	els := Sort(nil)

	assert.EqualValues(t, []int(nil), els) // check output
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {

	els := getElements(5)
	assert.NotNil(t, els)

	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

// worst case scenario with 10 items
func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

// worst case scenario with 1000 items
func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

// worst case scenario with 10000 items
func BenchmarkBubbleSort20000(b *testing.B) {
	els := getElements(20000)
	for i := 0; i < b.N; i++ {
		// Sort(els)
		sort.Ints(els)
	}
}

// worst case scenario with 25000 items
func BenchmarkBubbleSort25000(b *testing.B) {
	els := getElements(23500)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

// worst case scenario with 50000 items
func BenchmarkBubbleSort50000(b *testing.B) {
	els := getElements(50000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

// worst case scenario with 100000 items
func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
