package main

import (
	"fmt"
)

type Product interface {
	buyProduct() int
}

type Bread struct {
	price int
}

func (b *Bread) buyProduct() int {
	return b.price
}

type Water struct {
	price int
}

func (w *Water) buyProduct() int {
	return w.price
}

type Customer struct {
	shoppingCart []Product
	cash         int
}

func (c *Customer) addProductToCart(product Product) {
	c.shoppingCart = append(c.shoppingCart, product)
}

func (c *Customer) checkout() {
	totalCost := 0
	for _, product := range c.shoppingCart {
		totalCost += product.buyProduct()
	}
	c.cash -= totalCost
	c.shoppingCart = nil
}

func main() {
	customer := &Customer{cash: 100}

	bread := &Bread{price: 10}
	water := &Water{price: 5}

	customer.addProductToCart(bread)
	customer.addProductToCart(water)

	customer.checkout()

	fmt.Printf("Remaining money: %d\n", customer.cash)

	customer.addProductToCart(bread)
	customer.addProductToCart(bread)

	customer.checkout()

	fmt.Printf("Remaining money: %d\n", customer.cash)
}
