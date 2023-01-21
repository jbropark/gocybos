package gocybos

import (
	"fmt"
	"testing"
	"time"
)

type MstReceiver struct {
	t *testing.T
}

func (r *MstReceiver) Received() {
	fmt.Println("Received: " + time.Now().String())
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

	// time.Sleep(5 * time.Second)

	/*
		for {
			PumpWaitingMessages()
		}

	*/
}
