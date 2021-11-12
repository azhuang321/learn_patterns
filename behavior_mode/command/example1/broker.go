package example1

import "fmt"

type Broker struct {
	orderList []Order
}

func (b *Broker) Add(order Order) {
	b.orderList = append(b.orderList, order)
}

func (b *Broker) PlaceOrder() {
	fmt.Println(b.orderList)
	for _, v := range b.orderList {
		v.Execute()
	}
	b.orderList = make([]Order, 0)
}
