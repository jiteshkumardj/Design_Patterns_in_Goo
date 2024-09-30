package main

import "fmt"

//OCP
// type should be open for extention, closed for modification
// we're also going to talk about an enterprise pattern called specification
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}
type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", red, small}
	tree := Product{"Tree", green, medium}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}
	fmt.Println("Green products")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Println("Green product is :", v.name)
	}
	for _, v := range f.FilterBySize(products, large) {
		fmt.Println("Green product is :", v.name)
	}
}

/*And let's suppose that you want the end user of your website to be able to filter those items like filter

by price or filter by size and that sort of thing.

So you have a specification, you have a sort of product description that states, well, our website

has to allow the user to be able to filter by certain criteria.*/
