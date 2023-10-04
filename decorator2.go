package main

import "fmt"

type Coffee interface {
	getPrice() int
}

type Latte struct {
}

func (l *Latte) getPrice() int {
	return 7
}

type Cappuccino struct {
}

func (c *Cappuccino) getPrice() int {
	return 6
}

type AddHoney struct {
	coffee Coffee
}

func (a *AddHoney) getPrice() int {
	coffeePrice := a.coffee.getPrice()
	return coffeePrice + 2
}

type AddCaramel struct {
	coffee Coffee
}

func (a *AddCaramel) getPrice() int {
	coffeePrice := a.coffee.getPrice()
	return coffeePrice + 3
}

func main() {
	latte := &Latte{}
	latteWHoney := AddHoney{latte}
	fmt.Printf("Total price of the coffee: %d dollars\n", latteWHoney.getPrice())

	cappuccino := &Cappuccino{}
	cappuccinoWCaramel := AddCaramel{cappuccino}
	fmt.Printf("Total price of the coffee: %d dollars\n", cappuccinoWCaramel.getPrice())
}
