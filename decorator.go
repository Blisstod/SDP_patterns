package main

import "fmt"

type Weights interface {
	getWeight() int
}

type Barbell struct {
}

func (b *Barbell) getWeight() int {
	return 20
}

type CurlBar struct {
}

func (c *CurlBar) getWeight() int {
	return 10
}

type Adding5kg struct {
	weights Weights
}

func (a *Adding5kg) getWeight() int {
	return a.weights.getWeight() + 5
}

type Adding10kg struct {
	weights Weights
}

func (a *Adding10kg) getWeight() int {
	return a.weights.getWeight() + 10
}

func main() {
	barbell := &Barbell{}
	barbellWith10kg := &Adding10kg{weights: barbell} // Wrap the original barbell with Adding10kg decorator
	fmt.Printf("Total weight of the barbell: %d kgs\n", barbellWith10kg.getWeight())

	curlbar := &CurlBar{}
	curlbarWith5kg := &Adding5kg{weights: curlbar}
	fmt.Printf("Total weight of the curl bar: %d kgs\n", curlbarWith5kg.getWeight())
}
