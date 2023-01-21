package gocybos

import (
	"testing"
)

func TestCpCybos(t *testing.T) {
	InitCOM()
	defer ReleaseCOM()

	cybos := CpCybos{}
	cybos.Create()
	defer cybos.Release()

	t.Logf("CLSID: %v", cybos.clsid)
	t.Logf("Connect: %v", cybos.IsConnect())
	t.Logf("Connect: %v", ToBool(cybos.IsConnect()))
	t.Logf("TradeCount: %v", ToInt32(cybos.GetLimitRemainCount(LimitTypeTradeRequest)))
	t.Logf("NonTradeCount: %v", ToInt32(cybos.GetLimitRemainCount(LimitTypeNonTradeRequest)))
	t.Logf("SubCount: %v", ToInt32(cybos.GetLimitRemainCount(LimitTypeSubscribe)))
}
