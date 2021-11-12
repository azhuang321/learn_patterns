package example1

import "fmt"

type Stock struct {
	name     string
	quantity int
}

func NewStock() *Stock {
	stock := &Stock{}
	stock.name = "abc"
	stock.quantity = 10
	return stock
}

func (s *Stock) Buy() {
	fmt.Printf("buy name:%s,quantity:%d stock", s.name, s.quantity)
}

func (s *Stock) Sell() {
	fmt.Printf("sell name:%s,quantity:%d stock", s.name, s.quantity)
}
