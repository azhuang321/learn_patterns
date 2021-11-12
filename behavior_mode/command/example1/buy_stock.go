package example1

type BuyStock struct {
	abcStock *Stock
}

func NewBuyStock(stock *Stock) *BuyStock {
	bs := new(BuyStock)
	bs.abcStock = stock
	return bs
}

func (b *BuyStock) Execute() {
	b.abcStock.Buy()
}
