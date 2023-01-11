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
}
