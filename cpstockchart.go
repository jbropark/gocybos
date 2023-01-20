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
	StockChartFieldDelta
	StockChartFieldVolume StockChartField = 1 + iota
	StockChartFieldValue
	StockChartFieldAccSell // 호가 방식
	StockChartFieldAccBuy  // 호가 방식
	StockChartFieldListedNum
	StockChartFieldMarketCap
	StockChartFieldDeltaSign = 37
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
	StockChartHeaderDeltaCode
	StockChartHeaderDelta
	StockChartHeaderVolume
	StockChartHeaderSellBidding
	StockChartHeaderBuyBidding
	StockChartHeaderOpen
	StockChartHeaderHigh
	StockChartHeaderLow
	StockChartHeaderValue
	StockChartHeaderStatus
	StockChartHeaderListedNum
	StockChartHeaderCapital
	StockChartHeaderPrevTransaction
	StockChartHeaderLastUpdateTime
	StockChartHeaderUpperBoundPrice
	StockChartHeaderLowerBoundPrice
)

type CpStockChart struct {
	CpTrait
	fields ole.VARIANT
}

func (c *CpStockChart) Create(fields []StockChartField) {
	err := c.CpTrait.Create("CpSysDib.StockChart")
	if err != nil {
		panic(err)
	}
	c.fields = VariantInt32Slice(fields)
	c.SetInputValue(stockChartInputFields, &c.fields)
}

func (c *CpStockChart) SetInputValue(iType stockChartInputType, value any) {
	c.CpTrait.SetInputValue(int32(iType), value)
}

func (c *CpStockChart) GetHeaderValue(hType StockChartHeaderType) *ole.VARIANT {
	return c.CpTrait.GetHeaderValue(int32(hType))
}

func (c *CpStockChart) SetInputValues(
	stockCode string,
	countType StockChartCountType,
	dateStart time.Time,
	dateEnd time.Time,
	count uint64,
	dataType StockChartDataType,
	period uint16,
	gapType StockChartGapType,
	priceType StockChartPriceType,
	volumeType StockChartVolumeType,
) {
	c.SetInputValue(stockChartInputCode, Stock(stockCode))
	c.SetInputValue(stockChartInputCountType, rune(countType))
	c.SetInputValue(stockChartInputDateEnd, DateToUInt(dateEnd))
	c.SetInputValue(stockChartInputDateStart, DateToUInt(dateStart))
	c.SetInputValue(stockChartInputCount, count)
	c.SetInputValue(stockChartInputDataType, rune(dataType))
	c.SetInputValue(stockChartInputPeriod, period)
	c.SetInputValue(stockChartInputGapType, rune(gapType))
	c.SetInputValue(stockChartInputPriceType, rune(priceType))
	c.SetInputValue(stockChartInputVolumeType, rune(volumeType))
}
