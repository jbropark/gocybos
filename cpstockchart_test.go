package gocybos

import (
	"fmt"
	"testing"
	"time"
)

func TestCpStockChart(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	c := CpStockChart{}
	c.Create([]StockChartField{StockChartFieldDate, StockChartFieldClose})

	c.SetInputValues(
		"005930",
		StockChartCountTypePeriod,
		time.Date(2023, 1, 9, 0, 0, 0, 0, time.Local),
		time.Date(2023, 1, 19, 0, 0, 0, 0, time.Local),
		0,
		StockChartDataTypeMinute,
		1,
		StockChartGapNoCorrection,
		StockChartPriceNoAmend,
		StockChartVolumeContainAll,
	)

	c.BlockRequest()

	fmt.Println(c.GetHeaderValue(0))
	fmt.Println(c.GetHeaderValue(1))
	fmt.Println(ToSS(c.GetHeaderValue(2)))
	fmt.Println(ToInt64(c.GetHeaderValue(3)))

	count := ToInt32(c.GetHeaderValue(StockChartHeaderDataCount))
	fmt.Println(CastSlice(c.GetDataArray(0, count), VarToDate))
	fmt.Println(CastSlice(c.GetDataArray(1, count), ToInt64))

	c.Release()
}
