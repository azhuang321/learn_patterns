package main

import (
	"fmt"
	"go-patterns/behavior_mode/command"
)

func main() {
	abcStock := command.NewStock()

	buyStock := command.NewBuyStock(abcStock)
	sellStock := command.NewSellStock(abcStock)

	broker := new(command.Broker)

	broker.Add(buyStock)
	broker.Add(sellStock)

	fmt.Printf("%p\n", buyStock)
	fmt.Printf("%p\n", sellStock)

	broker.PlaceOrder()
}
