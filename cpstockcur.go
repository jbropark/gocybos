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
	StockCurHeaderTimeHM
	StockCurHeaderOpen
	StockCurHeaderHigh
	StockCurHeaderLow
	StockCurHeaderAskPrice
	StockCurHeaderBidPrice
	StockCurHeaderCmlVolume
	StockCurHeaderCmlValue
	StockCurHeaderClose  = iota + 2
	StockCurHeaderVolume = iota + 5
	StockCurHeaderTimeHMS
	StockCurHeaderPreMarketVolume = iota + 7
	StockCurHeaderConclusionType  = iota + 11
	StockCurHeaderCmlVolumeSell
	StockCurHeaderCmlVolumeBuy
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
