package main

import "fmt"

// Product interface
type IBarbell interface {
	getWeight() int
}

// Concrete Product
type Barbell struct {
	weight int
}

func (b *Barbell) getWeight() int {
	return b.weight
}

type PowerliftingBarbell struct {
	Barbell
}

type OlympicBarbell struct {
	Barbell
}

// creator
type BarbellFactory interface {
	createBarbell(weight int) IBarbell
}

// Concrete Creators
type PowerliftingBarbellFactory struct {
}

func (f *PowerliftingBarbellFactory) createBarbell(weight int) IBarbell {
	return &PowerliftingBarbell{Barbell{weight: weight}}
}

type OlympicBarbellFactory struct {
}

func (f *OlympicBarbellFactory) createBarbell(weight int) IBarbell {
	return &OlympicBarbell{Barbell{weight: weight}}
}

func main() {
	powerliftingBarbellFactory := PowerliftingBarbellFactory{}
	powerliftingBarbell := powerliftingBarbellFactory.createBarbell(40)

	olympicBarbellFactory := OlympicBarbellFactory{}
	olympicBarbell := olympicBarbellFactory.createBarbell(20)

	fmt.Println("The weight of Powerlifting Barbell:", powerliftingBarbell.getWeight())
	fmt.Println("The weight of Olympic Barbell:", olympicBarbell.getWeight())
}
