package gocybos

type CpStockCur struct {
	CpTrait
}

func (c *CpStockCur) Create() {
	err := c.CpTrait.Create("Dscbo1.StockCur")
	if err != nil {
		panic(err)
	}
}
