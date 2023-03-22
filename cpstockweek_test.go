package gocybos

import (
	"testing"
)

func TestCpStockWeek(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	cur := CpStockWeek{}
	cur.Create()

	cur.SetInputCode(StockIndex("001"))
	cur.BlockRequest()

	t.Log(cur.GetHeaderValue(StockWeekHeaderCount))
	t.Log(ToDate(cur.GetHeaderValue(StockWeekHeaderDate)))
	t.Log(cur.GetDataValue(StockWeekDataValue, 0))
}
