package gocybos

import "github.com/go-ole/go-ole"

const (
	KOSPI  = "001"
	KOSDAQ = "201"
)

type CpStockWeek struct {
	CpTrait
}

const (
	StockWeekInputCode = 0
)

type StockWeekHeaderType int32

const (
	StockWeekHeaderCode StockWeekHeaderType = iota
	StockWeekHeaderCount
	StockWeekHeaderDate
)

type StockWeekDataType int32

const (
	StockWeekDataDate StockWeekDataType = iota
	StockWeekDataOpen
	StockWeekDataHigh
	StockWeekDataLow
	StockWeekDataClose
	StockWeekDataDelta
	StockWeekDataVolume
	StockWeekDataDeltaRatio = 10
	StockWeekDataValue      = 20
)

func (c *CpStockWeek) Create() {
	err := c.CpTrait.Create("Dscbo1.StockWeek")
	if err != nil {
		panic(err)
	}
}

func (c *CpStockWeek) SetInputCode(code string) {
	c.SetInputValue(StockWeekInputCode, code)
}

func (c *CpStockWeek) GetHeaderValue(hType StockWeekHeaderType) *ole.VARIANT {
	return c.CpTrait.GetHeaderValue(int32(hType))
}

func (c *CpStockWeek) GetDataValue(dType StockWeekDataType, index int32) *ole.VARIANT {
	return c.CpTrait.GetDataValue(int32(dType), index)
}

func (c *CpStockWeek) GetValue(index int32) uint64 {
	return ToUInt64(c.CpTrait.GetDataValue(int32(StockWeekDataValue), index)) * 1_0000
}
