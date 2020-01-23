package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolume(t *testing.T) {
	volumeCollector := make(chan int)
	d := dimensions{Length: 4, Width: 4, Height: 4}
	go volume(d, volumeCollector)
	results := <-volumeCollector
	assert.EqualValues(t, results, 64)
}

func TestCalculateTotalProductVolume(t *testing.T) {
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 4, Width: 4, Height: 4}},
	}
	assert.EqualValues(t, 91, calculateTotalProductVolume(products))
}
