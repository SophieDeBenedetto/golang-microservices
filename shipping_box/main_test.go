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

func TestGetBestBoxSmallerBox(t *testing.T) {
	boxes := []Box{
		Box{Dimensions: dimensions{Length: 10, Width: 10, Height: 10}},
		Box{Dimensions: dimensions{Length: 5, Width: 5, Height: 5}},
	}
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
	}
	bestBox := getBestBox(boxes, products)
	assert.EqualValues(t, 5, bestBox.Dimensions.Length)
	assert.EqualValues(t, 5, bestBox.Dimensions.Width)
	assert.EqualValues(t, 5, bestBox.Dimensions.Height)
}

func TestGetBestBoxTooSmallBox(t *testing.T) {
	boxes := []Box{
		Box{Dimensions: dimensions{Length: 10, Width: 10, Height: 10}},
		Box{Dimensions: dimensions{Length: 5, Width: 5, Height: 5}},
		Box{Dimensions: dimensions{Length: 1, Width: 1, Height: 1}},
	}
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
	}
	bestBox := getBestBox(boxes, products)
	assert.EqualValues(t, 5, bestBox.Dimensions.Length)
	assert.EqualValues(t, 5, bestBox.Dimensions.Width)
	assert.EqualValues(t, 5, bestBox.Dimensions.Height)
}

func TestGetBestBoxPerfectFit(t *testing.T) {
	boxes := []Box{
		Box{Dimensions: dimensions{Length: 10, Width: 10, Height: 10}},
		Box{Dimensions: dimensions{Length: 5, Width: 5, Height: 5}},
		Box{Dimensions: dimensions{Length: 1, Width: 1, Height: 1}},
		Box{Dimensions: dimensions{Length: 9, Width: 3, Height: 3}},
	}
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
	}
	bestBox := getBestBox(boxes, products)
	assert.EqualValues(t, 9, bestBox.Dimensions.Length)
	assert.EqualValues(t, 3, bestBox.Dimensions.Width)
	assert.EqualValues(t, 3, bestBox.Dimensions.Height)
}

func TestGetBestBoxNoFits(t *testing.T) {
	boxes := []Box{
		Box{Dimensions: dimensions{Length: 1, Width: 1, Height: 1}},
	}
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
	}
	bestBox := getBestBox(boxes, products)
	assert.EqualValues(t, 0, bestBox.Dimensions.Length)
	assert.EqualValues(t, 0, bestBox.Dimensions.Width)
	assert.EqualValues(t, 0, bestBox.Dimensions.Height)
}
