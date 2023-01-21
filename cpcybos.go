package gocybos

import "github.com/go-ole/go-ole"

type CpCybos struct {
	CpTrait
}

type LimitType int16

const (
	LimitTypeTradeRequest LimitType = iota
	LimitTypeNonTradeRequest
	LimitTypeSubscribe
)

func (c *CpCybos) Create() {
	err := c.CpTrait.Create("CpUtil.CpCybos")
	if err != nil {
		panic(err)
	}
}

func (c *CpCybos) IsConnect() *ole.VARIANT {
	return c.Object.MustGet("IsConnect")
}

func (c *CpCybos) ServerType() *ole.VARIANT {
	return c.Object.MustGet("ServerType")
}

func (c *CpCybos) LimitRequestRemainTime() *ole.VARIANT {
	return c.Object.MustGet("LimitRequestRemainTime")
}

func (c *CpCybos) GetLimitRemainCount(limitType LimitType) *ole.VARIANT {
	return c.Object.MustCall("GetLimitRemainCount", int16(limitType))
}

func (c *CpCybos) GetLimitRemainTime(limitType LimitType) *ole.VARIANT {
	return c.Object.MustCall("GetLimitRemainTime", int16(limitType))
}

func (c *CpCybos) PlusDisconnect() {
	c.Object.MustCall("PlusDisconnect")
}
