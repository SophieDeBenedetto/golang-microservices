package main

import (
	"fmt"
	"sync"
)

// Product describes a product
type Product struct {
	Name       string
	Dimensions dimensions
}

// Box describes a box
type Box struct {
	Dimensions dimensions
}

type dimensions struct {
	Length int
	Width  int
	Height int
}

// BoxResult describes the box and the fit, i.e. volume difference between box and products
type BoxResult struct {
	Box Box
	Fit int
}

func main() {
	fmt.Println("Running...")
	// test cases:
	// 1. when one box is a better fit
	// 2. when boxes includes a box that is too small
	// 3. when boxes includes a box that is a perfect fit
	// 4. when no boxes match
	boxes := []Box{
		Box{Dimensions: dimensions{Length: 10, Width: 10, Height: 10}},
		Box{Dimensions: dimensions{Length: 5, Width: 5, Height: 5}},
		Box{Dimensions: dimensions{Length: 1, Width: 1, Height: 1}},
		Box{Dimensions: dimensions{Length: 1, Width: 1, Height: 1}},
		Box{Dimensions: dimensions{Length: 9, Width: 3, Height: 3}},
	}
	products := []Product{
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
		Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}},
	}
	result := getBestBox(boxes, products)
	fmt.Println("Best Box:")
	fmt.Println(result)
}

func getBestBox(availableBoxes []Box, products []Product) Box {
	totalProductVol := calculateTotalProductVolume(products)
	return concurrentGetBestBox(totalProductVol, availableBoxes)
}

func concurrentGetBestBox(totalProductVol int, availableBoxes []Box) Box {
	input := make(chan BoxResult)
	output := make(chan Box)
	defer close(output)
	var wg sync.WaitGroup

	go collectBoxResult(input, output, &wg)

	for _, box := range availableBoxes {
		wg.Add(1)
		go calculateBoxResult(box, totalProductVol, input)
	}
	wg.Wait()
	close(input)
	return <-output
}

func collectBoxResult(input chan BoxResult, output chan Box, wg *sync.WaitGroup) {
	var bestFit int
	var started bool
	var bestBox Box
	for result := range input {
		if !started {
			if result.Fit >= 0 {
				started = true
				bestFit = result.Fit
			}
		}

		if result.Fit >= 0 && result.Fit <= bestFit {
			bestFit = result.Fit
			bestBox = result.Box
		}
		wg.Done()
	}
	output <- bestBox
}

func calculateBoxResult(box Box, totalProductVol int, input chan BoxResult) {
	boxVol := box.Dimensions.Length * box.Dimensions.Width * box.Dimensions.Height
	fit := boxVol - totalProductVol
	input <- BoxResult{
		Box: box,
		Fit: fit,
	}
}

func calculateTotalProductVolume(products []Product) int {
	volumeCollector := make(chan int)
	totalVolCalculator := make(chan int)
	defer close(totalVolCalculator)
	var wg sync.WaitGroup

	go collectVolumes(volumeCollector, totalVolCalculator, &wg)

	for _, product := range products {
		wg.Add(1)
		go volume(product.Dimensions, volumeCollector)
	}
	wg.Wait()
	close(volumeCollector)
	totalProductVol := <-totalVolCalculator
	return totalProductVol
}

func volume(d dimensions, volumeCollector chan int) {
	vol := d.Length * d.Width * d.Height
	volumeCollector <- vol
}

func collectVolumes(volumeCollector chan int, totalVolCalculator chan int, wg *sync.WaitGroup) {
	total := 0
	for vol := range volumeCollector {
		total = total + vol
		wg.Done()
	}
	totalVolCalculator <- total
}
