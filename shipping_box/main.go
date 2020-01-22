package main

import "fmt"

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

func main() {
	fmt.Println("Running...")
	product := Product{Dimensions: dimensions{Length: 3, Width: 3, Height: 3}}
	box := Box{Dimensions: dimensions{Length: 4, Width: 4, Height: 4}}
	canFit := productCanFit(product, box)
	fmt.Println(canFit)
}

func productCanFit(p Product, b Box) bool {
	fit := b.Dimensions.Length > p.Dimensions.Length &&
		b.Dimensions.Height > p.Dimensions.Height &&
		b.Dimensions.Width > p.Dimensions.Width
	return fit
}
