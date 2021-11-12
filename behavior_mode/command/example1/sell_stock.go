package example1

type SellStock struct {
	abcStock *Stock
}

func NewSellStock(stock *Stock) *SellStock {
	bs := new(SellStock)
	bs.abcStock = stock
	return bs
}

func (b *SellStock) Execute() {
	b.abcStock.Sell()
}
