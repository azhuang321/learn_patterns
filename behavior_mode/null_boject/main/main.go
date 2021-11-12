package main

import (
	"fmt"
	"go-patterns/behavior_mode/null_boject"
)

func main() {
	customer1 := new(null_boject.CustomFactory).GetCustomer("123")
	customer2 := new(null_boject.CustomFactory).GetCustomer("345")
	customer3 := new(null_boject.CustomFactory).GetCustomer("789")
	customer4 := new(null_boject.CustomFactory).GetCustomer("Laura")

	fmt.Println("Customers")
	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
}
