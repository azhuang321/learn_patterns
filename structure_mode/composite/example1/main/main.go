package main

import "go-patterns/structure_mode/composite"

func main() {
	CEO := composite.New("john", "CEO", 30000)
	headSales := composite.New("Robert", "head sales", 20000)
	headMarketing := composite.New("Mical", "head marketing", 20000)

	clerk1 := composite.New("lar1", "marketing", 10000)
	clerk2 := composite.New("lar2", "marketing", 10000)

	sales1 := composite.New("sales1", "sale", 10000)
	sales2 := composite.New("sales2", "sale", 10000)

	CEO.Add(headSales)
	CEO.Add(headMarketing)

	headSales.Add(sales1)
	headSales.Add(sales2)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	CEO.List()

}
