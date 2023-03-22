package gocybos

import (
	"fmt"
	"testing"
)

func TestCpStockChart(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	c := CpStockChart{}
	defer c.Release()

	fields := VariantInt32Slice([]StockChartField{
		StockChartFieldDate,
		StockChartFieldTime,
		StockChartFieldOpen,
		StockChartFieldHigh,
		StockChartFieldLow,
		StockChartFieldClose,
		StockChartFieldPriceDelta,
		StockChartFieldVolume,
		StockChartFieldValue,
		StockChartFieldVolumeBidSellTotal,
		StockChartFieldVolumeBidBuyTotal,
		StockChartFieldListedShare,
		StockChartFieldMarketCap,
		StockChartFieldPriceDeltaSign,
		StockChartFieldVolumeConSellTotal,
		StockChartFieldVolumeConBuyTotal,
	})
	c.Create(&fields)

	c.SetInputCode(Stock("005930"))
	c.SetInputCountType(StockChartCountTypeNum)
	// c.SetInputDateStart(time.Date(2023, 1, 9, 0, 0, 0, 0, time.Local))
	// c.SetInputDateEnd(time.Date(2023, 1, 19, 0, 0, 0, 0, time.Local))
	c.SetInputCount(2)
	c.SetInputDataType(StockChartDataTypeMinute)
	c.SetInputGapType(StockChartGapNoCorrection)
	c.SetInputPriceType(StockChartPriceNoAmend)
	c.SetInputVolumeType(StockChartVolumeContainAll)

	c.BlockRequest()

	fmt.Println(ToStr(c.GetHeaderValue(StockChartHeaderCode)))
	fmt.Println(ToInt64(c.GetHeaderValue(StockChartHeaderFieldCount)))
	fmt.Println(ToSS(c.GetHeaderValue(StockChartHeaderFieldNames)))
	fmt.Println(ToDate(c.GetHeaderValue(StockChartHeaderLastTradeDay)))

	count := ToInt32(c.GetHeaderValue(StockChartHeaderDataCount))
	fmt.Println(CastSlice(c.GetDataArray(0, count), ToDate))
	fmt.Println(CastSlice(c.GetDataArray(1, count), ToTimeHM))
	fmt.Println(CastSlice(c.GetDataArray(2, count), ToInt64))

}

func TestCpStockChartIndex(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	c := CpStockChart{}
	defer c.Release()

	fields := VariantInt32Slice([]StockChartField{
		StockChartFieldDate,
		StockChartFieldOpen,
		StockChartFieldHigh,
		StockChartFieldLow,
		StockChartFieldClose,
		StockChartFieldPriceDelta,
		StockChartFieldVolume,
		StockChartFieldValue,
		StockChartFieldMarketCap,
		StockChartFieldPriceDeltaSign,
	})
	c.Create(&fields)

	c.SetInputCode(StockIndex("201"))
	c.SetInputCountType(StockChartCountTypeNum)
	// c.SetInputDateStart(time.Date(2023, 1, 9, 0, 0, 0, 0, time.Local))
	// c.SetInputDateEnd(time.Date(2023, 1, 19, 0, 0, 0, 0, time.Local))
	c.SetInputCount(2)
	c.SetInputDataType(StockChartDataTypeDay)
	c.SetInputGapType(StockChartGapNoCorrection)
	c.SetInputPriceType(StockChartPriceNoAmend)
	c.SetInputVolumeType(StockChartVolumeContainAll)

	c.BlockRequest()

	fmt.Println(ToStr(c.GetHeaderValue(StockChartHeaderCode)))
	fmt.Println(ToInt64(c.GetHeaderValue(StockChartHeaderFieldCount)))
	fmt.Println(ToSS(c.GetHeaderValue(StockChartHeaderFieldNames)))
	fmt.Println(ToDate(c.GetHeaderValue(StockChartHeaderLastTradeDay)))

	count := ToInt32(c.GetHeaderValue(StockChartHeaderDataCount))
	fmt.Println(CastSlice(c.GetDataArray(0, count), ToDate))
	fmt.Println(CastSlice(c.GetDataArray(1, count), ToFloat32))
	fmt.Println(CastSlice(c.GetDataArray(2, count), ToFloat32))
	fmt.Println(CastSlice(c.GetDataArray(3, count), ToFloat32))
	fmt.Println(CastSlice(c.GetDataArray(4, count), ToFloat32))
	fmt.Println(CastSlice(c.GetDataArray(5, count), ToFloat32))
	fmt.Println(CastSlice(c.GetDataArray(6, count), ToUInt64))
	fmt.Println(CastSlice(c.GetDataArray(7, count), ToUInt64))
	fmt.Println(CastSlice(c.GetDataArray(8, count), ToUInt64))
}
