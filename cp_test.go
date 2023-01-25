package gocybos

import (
	"testing"
	"time"
)

type MstReceiver struct {
	t        *testing.T
	received bool
}

func (r *MstReceiver) Received() {
	r.t.Log("Received: " + time.Now().String())
	r.received = true
}

func TestCpMst(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	mst := CpTrait{}
	err := mst.Create("Dscbo1.StockMst")
	if err != nil {
		panic(err)
	}

	rec := MstReceiver{t, false}
	mst.BindEvent(&rec)

	mst.SetInputValue(0, "A005930")
	mst.Request()
	t.Log("Request: " + time.Now().String())
	time.Sleep(1 * time.Second)

	t.Logf("Received: %v", rec.received)
	if rec.received == false {
		panic("Cannot receive response")
	}
}
