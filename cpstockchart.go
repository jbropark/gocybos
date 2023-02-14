package gocybos

import (
	"github.com/go-ole/go-ole"
	"time"
)

type stockChartInputType int32

const (
	stockChartInputCode stockChartInputType = iota
	stockChartInputCountType
	stockChartInputDateEnd
	stockChartInputDateStart
	stockChartInputCount
	stockChartInputFields
	stockChartInputDataType
	stockChartInputPeriod
	stockChartInputGapType
	stockChartInputPriceType
	stockChartInputVolumeType
)

type StockChartCountType rune

const (
	StockChartCountTypePeriod StockChartCountType = '1'
	StockChartCountTypeNum    StockChartCountType = '2'
)

type StockChartDataType rune

const (
	StockChartDataTypeDay    StockChartDataType = 'D'
	StockChartDataTypeMonth  StockChartDataType = 'M'
	StockChartDataTypeMinute StockChartDataType = 'm'
	StockChartDataTypeWeek   StockChartDataType = 'W'
	StockChartDataTypeTick   StockChartDataType = 'T'
)

type StockChartGapType rune

const (
	StockChartGapNoCorrection StockChartGapType = '0'
	StockChartGapCorrection   StockChartGapType = '1'
)

type StockChartPriceType rune

const (
	StockChartPriceNoAmend StockChartPriceType = '0'
	StockChartPriceAmend   StockChartPriceType = '1'
)

type StockChartVolumeType rune

const (
	StockChartVolumeContainAll    StockChartVolumeType = '1'
	StockChartVolumeContainAfter  StockChartVolumeType = '2'
	StockChartVolumeOnlyMarket    StockChartVolumeType = '3'
	StockChartVolumeContainBefore StockChartVolumeType = '4'
)

type StockChartField int32

const (
	StockChartFieldDate StockChartField = iota
	StockChartFieldTime
	StockChartFieldOpen
	StockChartFieldHigh
	StockChartFieldLow
	StockChartFieldClose
	StockChartFieldPriceDelta
	StockChartFieldVolume StockChartField = 1 + iota
	StockChartFieldValue
	StockChartFieldVolumeBidSellTotal
	StockChartFieldVolumeBidBuyTotal
	StockChartFieldListedShare
	StockChartFieldMarketCap
	StockChartFieldPriceDeltaSign     = 37
	StockChartFieldVolumeConSellTotal = 62
	StockChartFieldVolumeConBuyTotal  = 63
)

type StockChartHeaderType int32

const (
	StockChartHeaderCode          StockChartHeaderType = iota //string
	StockChartHeaderFieldCount                                //short
	StockChartHeaderFieldNames                                //string array
	StockChartHeaderDataCount                                 //long
	StockChartHeaderLastTickCount                             //ushort
	StockChartHeaderLastTradeDay                              //ulong
	StockChartHeaderPrevClose
	StockChartHeaderCurrentPrice
	StockChartHeaderPriceDeltaType
	StockChartHeaderPriceDelta
	StockChartHeaderVolume
	StockChartHeaderAskPrice
	StockChartHeaderBidPrice
	StockChartHeaderOpen
	StockChartHeaderHigh
	StockChartHeaderLow
	StockChartHeaderValue
	StockChartHeaderStatus
	StockChartHeaderListedShare
	StockChartHeaderCapital // 자본금
	StockChartHeaderPrevVolume
	StockChartHeaderLastUpdateTimeHM
	StockChartHeaderMaxPrice
	StockChartHeaderMinPrice
)

type CpStockChart struct {
	CpTrait
	fields *ole.VARIANT
}

func (c *CpStockChart) Create(fields *ole.VARIANT) {
	err := c.CpTrait.Create("CpSysDib.StockChart")
	if err != nil {
		panic(err)
	}
	c.fields = fields
	c.SetInputValue(stockChartInputFields, c.fields)
}

func (c *CpStockChart) SetInputValue(iType stockChartInputType, value any) {
	c.CpTrait.SetInputValue(int32(iType), value)
}

func (c *CpStockChart) GetHeaderValue(hType StockChartHeaderType) *ole.VARIANT {
	return c.CpTrait.GetHeaderValue(int32(hType))
}

func (c *CpStockChart) SetInputCode(stockCode string) {
	c.SetInputValue(stockChartInputCode, Stock(stockCode))
}

func (c *CpStockChart) SetInputCountType(countType StockChartCountType) {
	c.SetInputValue(stockChartInputCountType, rune(countType))
}

func (c *CpStockChart) SetInputDateStart(date time.Time) {
	c.SetInputValue(stockChartInputDateStart, DateToUInt(date))
}

func (c *CpStockChart) SetInputDateEnd(date time.Time) {
	c.SetInputValue(stockChartInputDateEnd, DateToUInt(date))
}

func (c *CpStockChart) SetInputCount(count uint64) {
	c.SetInputValue(stockChartInputCount, count)
}

func (c *CpStockChart) SetInputDataType(dataType StockChartDataType) {
	c.SetInputValue(stockChartInputDataType, rune(dataType))
}

func (c *CpStockChart) SetInputPeriod(period uint16) {
	c.SetInputValue(stockChartInputPeriod, period)
}

func (c *CpStockChart) SetInputGapType(gapType StockChartGapType) {
	c.SetInputValue(stockChartInputGapType, rune(gapType))
}

func (c *CpStockChart) SetInputPriceType(priceType StockChartPriceType) {
	c.SetInputValue(stockChartInputPriceType, rune(priceType))
}

func (c *CpStockChart) SetInputVolumeType(volumeType StockChartVolumeType) {
	c.SetInputValue(stockChartInputVolumeType, rune(volumeType))
}
