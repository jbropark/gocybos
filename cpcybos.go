package gocybos

import "github.com/go-ole/go-ole"

type CpCybos struct {
	CpTrait
}

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

func (c *CpCybos) GetLimitRemainCount(limitType int16) *ole.VARIANT {
	return c.Object.MustCall("GetLimitRemainCount", limitType)
}

func (c *CpCybos) GetLimitRemainTime(limitType int16) *ole.VARIANT {
	return c.Object.MustCall("GetLimitRemainTime", limitType)
}

func (c *CpCybos) PlusDisconnect() {
	c.Object.MustCall("PlusDisconnect")
}
