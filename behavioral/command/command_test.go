package command

import "testing"

// Client

func TestPattern(t *testing.T) {
	// New receiver and salesman
	shop := new(Shop)
	salesman := new(Salesman)
	// New orders to salesman
	salesman.OrderList = append(salesman.OrderList, &OrderBuy{shop})
	salesman.OrderList = append(salesman.OrderList, &OrderBuy{shop})
	salesman.OrderList = append(salesman.OrderList, &OrderSell{shop})
	// Proceed orders
	salesman.Proceed()
}
