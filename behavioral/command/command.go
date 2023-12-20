package command

import "fmt"

// Receiver

type Shop struct {
	item   string
	amount int
}

func (s *Shop) Sell() {
	fmt.Println("Sold item")
}

func (s *Shop) Buy() {
	fmt.Println("Bought item")
}

// Abstract command

type Order interface {
	Place()
}

// Concrete commands

type OrderBuy struct {
	shop *Shop
}

func (o *OrderBuy) Place() {
	o.shop.Buy()
}

type OrderSell struct {
	shop *Shop
}

func (o *OrderSell) Place() {
	o.shop.Sell()
}

// Invoker

type Salesman struct {
	OrderList []Order
}

func (s *Salesman) Proceed() {
	for _, cmd := range s.OrderList {
		cmd.Place()
	}
}
