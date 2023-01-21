package gocybos

import (
	"fmt"
	"testing"
	"time"
)

type CurReceiver struct {
	t *testing.T
	c *CpStockCur
}

func (r *CurReceiver) Received() {
	fmt.Printf(
		"[%s] Received: %s %s %v %v price {%v, %v, %v, %v[%v], %v[%v]} total {%v +%v -%v %v} bidding {+%v -%v}",
		time.Now().Format(time.RFC3339),
		ToStr(r.c.GetHeaderValue(StockCurHeaderCode)),
		ToStr(r.c.GetHeaderValue(StockCurHeaderName)),
		ToTimeHM(r.c.GetHeaderValue(StockCurHeaderTime)),
		ToTimeHMS(r.c.GetHeaderValue(StockCurHeaderSecond)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderOpen)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderHigh)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderLow)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderClose)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderDelta)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderVolume)),
		ToConclusionType(r.c.GetHeaderValue(StockCurHeaderConclusionType)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderAccVolume)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderAccBuy)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderAccSell)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderAccValue)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderBuyBidding)),
		ToInt64(r.c.GetHeaderValue(StockCurHeaderSellBidding)),
	)
}

func TestCpStockCur(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	cur := CpStockCur{}
	cur.Create()
	cur.BindEvent(&CurReceiver{t, &cur})

	cur.SetInputValue(0, "A005930")
	cur.Subscribe()

	t.Log("Start Subscribing\n")

	for {
		PumpWaitingMessages()
	}
}
