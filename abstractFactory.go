package main

import "fmt"

type AbstractFactory interface {
	createBarbell() IBarbell
	createPlate(weight int) IPlate
}

// iron factory
type IronFactory struct {
}

func (i *IronFactory) createBarbell() IBarbell {
	return &IronFactoryBarbell{}
}

func (i *IronFactory) createPlate(weight int) IPlate {
	return &IronFactoryPlate{Plate: Plate{weight: weight}}
}

// bumper factory
type BumperFactory struct{}

func (b *BumperFactory) createBarbell() IBarbell {
	return &BumperFactoryBarbell{}
}

func (b *BumperFactory) createPlate(weight int) IPlate {
	return &BumperFactoryPlate{Plate: Plate{weight: weight}}
}

// barbell
type IBarbell interface {
	getWeight() int
	addPlate(plate IPlate)
}

type Barbell struct {
	weight int
}

func (b *Barbell) getWeight() int {
	return b.weight + 20
}

func (b *Barbell) addPlate(plate IPlate) {
	b.weight += plate.getWeight()
}

// plate
type IPlate interface {
	setWeight(weight int)
	getWeight() int
}

type Plate struct {
	weight int
}

func (p *Plate) setWeight(weight int) {
	p.weight = weight
}

func (p *Plate) getWeight() int {
	return p.weight
}

// concrete products
type IronFactoryBarbell struct {
	Barbell
}
type IronFactoryPlate struct {
	Plate
}
type BumperFactoryBarbell struct {
	Barbell
}
type BumperFactoryPlate struct {
	Plate
}

func main() {
	ironFactory := &IronFactory{}
	barbell := ironFactory.createBarbell()
	barbell.addPlate(ironFactory.createPlate(20))
	fmt.Println("The weight of Iron Factory barbell:", barbell.getWeight())
	plate := ironFactory.createPlate(20)
	fmt.Println("The weight of Iron Factory plate:", plate.getWeight())
	barbell.addPlate(plate)
	fmt.Println("The weight of Iron Factory barbell:", barbell.getWeight())

	fmt.Println()

	bumperFactory := &BumperFactory{}
	bumperFactoryBarbell := bumperFactory.createBarbell()
	bumperFactoryBarbell.addPlate(bumperFactory.createPlate(10))
	fmt.Println("The weight of Bumper Factory barbell:", bumperFactoryBarbell.getWeight())
	bumperFactoryPlate := bumperFactory.createPlate(10)
	fmt.Println("The weight of Bumper Factory plate:", bumperFactoryPlate.getWeight())
	bumperFactoryBarbell.addPlate(bumperFactoryPlate)
	fmt.Println("The weight of Bumper Factory barbell:", bumperFactoryBarbell.getWeight())
}
