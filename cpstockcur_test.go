package gocybos

import (
	"fmt"
	"testing"
	"time"
)

type curLogReceiver struct {
	t *testing.T
	c *CpStockCur
}

func checkValid(vol int64, buy int64, sell int64, pre int64) {
	if vol != pre+buy+sell {
		panic(fmt.Errorf("inconsistent: Vol(%v) != pre(%v) + Buy(%v) + Sell(%v)", vol, pre, buy, sell))
	}
}

func (r *curLogReceiver) Received() {
	r.t.Logf(
		"[%s] Received: %s %s %v %v price {%v, %v, %v, %v[%v], %v[%v]} vol {%v (%v) +%v -%v %v} bid {+%v -%v}",
		time.Now().Format(time.RFC3339),
		ToStr(r.c.GetHeaderValue(StockCurHeaderCode)),
		ToStr(r.c.GetHeaderValue(StockCurHeaderName)),
		ToTimeHM(r.c.GetHeaderValue(StockCurHeaderTimeMinute)),
		ToTimeHMS(r.c.GetHeaderValue(StockCurHeaderTimeSecond)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderOpen)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderHigh)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderLow)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderClose)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderDelta)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderVolume)),
		ToConclusionType(r.c.GetHeaderValue(StockCurHeaderConclusionType)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolume)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderPreMarketVolume)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolumeBuy)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolumeSell)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlValue)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderBidPrice)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderAskPrice)),
	)

	checkValid(
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolume)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderPreMarketVolume)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolumeBuy)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderCmlVolumeSell)),
	)
}

func TestCpStockCur(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	cur := CpStockCur{}
	cur.Create()
	cur.BindEvent(&curLogReceiver{t, &cur})

	cur.SetInputCode("005930")
	cur.Subscribe()

	t.Log("Start Subscribing\n")

	time.Sleep(time.Minute)
}
