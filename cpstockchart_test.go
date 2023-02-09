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
		StockChartFieldClose,
	})
	c.Create(&fields)

	c.SetInputCode("005930")
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
	fmt.Println(ToSS(c.GetHeaderValue(StockChartHeaderFieldNames)))
	fmt.Println(ToDate(c.GetHeaderValue(StockChartHeaderLastTradeDay)))

	count := ToInt32(c.GetHeaderValue(StockChartHeaderDataCount))
	fmt.Println(CastSlice(c.GetDataArray(0, count), ToDate))
	fmt.Println(CastSlice(c.GetDataArray(1, count), ToTimeHM))
	fmt.Println(CastSlice(c.GetDataArray(2, count), ToInt64))

}
