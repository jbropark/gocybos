package gocybos

import (
	"github.com/go-ole/go-ole"
)

func Stock(stockCode string) string {
	if stockCode[0] == 'A' {
		return stockCode
	}

	return "A" + stockCode
}

func StockIndex(indexCode string) string {
	if indexCode[0] == 'U' {
		return indexCode
	}

	return "U" + indexCode
}

func ToSS(r *ole.VARIANT) []string {
	sa := r.ToArray().ToValueArray()
	ret := make([]string, len(sa))
	for i, v := range sa {
		ret[i] = v.(string)
	}
	return ret
}

func ToInt32(r *ole.VARIANT) int32 {
	return int32(r.Val)
}

func ToUInt32(r *ole.VARIANT) uint32 {
	return uint32(r.Val)
}

func ToStr(r *ole.VARIANT) string {
	return r.ToString()
}

func ToBool(r *ole.VARIANT) bool {
	return r.Val != 0
}

func ToInt64(r *ole.VARIANT) int64 {
	return r.Val
}

func ToUInt64(r *ole.VARIANT) uint64 {
	return uint64(r.Val)
}

func ToRune(r *ole.VARIANT) rune {
	return r.Value().(rune)
}

func ToFloat32(r *ole.VARIANT) float32 {
	return r.Value().(float32)
}

func ToFloat64(r *ole.VARIANT) float64 {
	return r.Value().(float64)
}
