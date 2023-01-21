package gocybos

import "github.com/go-ole/go-ole"

type CpStockCur struct {
	CpTrait
}

type StockCurHeaderType int32

const (
	StockCurHeaderCode StockCurHeaderType = iota
	StockCurHeaderName
	StockCurHeaderDelta
	StockCurHeaderTime
	StockCurHeaderOpen
	StockCurHeaderHigh
	StockCurHeaderLow
	StockCurHeaderSellBidding
	StockCurHeaderBuyBidding
	StockCurHeaderAccVolume
	StockCurHeaderAccValue
	StockCurHeaderClose
	StockCurHeaderVolume = iota + 5
	StockCurHeaderSecond
	StockCurHeaderConclusionType = iota + 12
	StockCurHeaderAccSell
	StockCurHeaderAccBuy
)

type ConclusionType rune

const (
	ConclusionTypeBuy  ConclusionType = '1'
	ConclusionTypeSell ConclusionType = '2'
)

const (
	StockCurInputCode = 0
)

func (c *CpStockCur) Create() {
	err := c.CpTrait.Create("Dscbo1.StockCur")
	if err != nil {
		panic(err)
	}
}

func (c *CpStockCur) SetInputCode(stockCode string) {
	c.SetInputValue(StockCurInputCode, Stock(stockCode))
}

func (c *CpStockCur) GetHeaderValue(hType StockCurHeaderType) *ole.VARIANT {
	return c.CpTrait.GetHeaderValue(int32(hType))
}

func ToConclusionType(value *ole.VARIANT) ConclusionType {
	return ConclusionType(ToRune(value))
}
