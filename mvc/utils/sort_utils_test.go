package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// Initialization
	els := []int{9, 8, 7, 6, 5}
	// Execution
	elements := BubbleSort(els)
	// Validation
	assert.EqualValues(t, elements, []int{5, 6, 7, 8, 9})
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	els := []int{5, 6, 7, 8, 9}
	// Execution
	elements := BubbleSort(els)
	// Validation
	assert.EqualValues(t, elements, []int{5, 6, 7, 8, 9})
}

// Run via command line: `go test -bench=.`
func BenchmarkBubbleSort(b *testing.B) {
	els := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
