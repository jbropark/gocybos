package gocybos

import (
	"fmt"
	"testing"
	"time"
)

type MstReceiver struct {
	t *testing.T
}

func (r *MstReceiver) Received(c *CpTrait) {
	fmt.Println("Received: " + time.Now().String())
	fmt.Println(c.GetHeaderValue(1))
}

func TestCpMst(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	mst := CpTrait{}
	err := mst.Create("Dscbo1.StockMst")
	if err != nil {
		panic(err)
	}
	mst.BindEvent(&MstReceiver{t})

	mst.SetInputValue(0, "A005930")
	mst.Request()
	fmt.Println("Request: " + time.Now().String())

	for {
		PumpWaitingMessages()
	}
}

type CurReceiver struct {
	t *testing.T
}

func (r *CurReceiver) Received(c *CpTrait) {
	fmt.Println("Received: " + time.Now().String())
	fmt.Println(c.GetHeaderValue(1))
}

func TestCpStockCur(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	cur := CpStockCur{}
	cur.Create()
	cur.BindEvent(&CurReceiver{t})

	cur.SetInputValue(0, "A005930")
	cur.Subscribe()

	t.Log(PumpWaitingMessages())
}
