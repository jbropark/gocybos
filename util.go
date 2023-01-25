package gocybos

import (
	"fmt"
	ole "github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
	"time"
)

func InitCOM() {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		panic(err)
	}
}

func ReleaseCOM() {
	ole.CoUninitialize()
}

func IsUserAnAdmin() (bool, error) {
	dllShell32 := windows.NewLazySystemDLL("Shell32.dll")
	defer func(handle windows.Handle) {
		err := windows.FreeLibrary(handle)
		if err != nil {
			panic(err)
		}
	}(windows.Handle(dllShell32.Handle()))

	procIsUerAnAdmin := dllShell32.NewProc("IsUserAnAdmin")
	res, _, winErr := procIsUerAnAdmin.Call()

	if winErr != windows.NTE_OP_OK {
		return false, fmt.Errorf("[error %d] Failed call IsUserAnAdmin", winErr)
	}
	return res != 0, nil
}

func DateToUInt(date time.Time) uint64 {
	year, month, day := date.Date()
	return uint64(year*1_00_00 + int(month)*1_00 + day)
}

func UIntToDate(value uint64) time.Time {
	return time.Date(
		int(value/1_00_00),
		time.Month(value/100%100),
		int(value%100),
		0, 0, 0, 0,
		time.Local,
	)
}

func UIntToTimeHM(value uint64) time.Time {
	return time.Date(
		1, time.Month(1), 1,
		int(value/1_00),
		int(value%1_00),
		0, 0, time.Local,
	)
}

func UIntToTimeHMS(value uint64) time.Time {
	return time.Date(
		1, time.Month(1), 1,
		int(value/1_00_00),
		int(value/1_00%1_00),
		int(value%1_00),
		0, time.Local,
	)
}

func ToDate(value *ole.VARIANT) time.Time {
	return UIntToDate(ToUInt64(value))
}

func ToTimeHM(value *ole.VARIANT) time.Time {
	return UIntToTimeHM(ToUInt64(value))
}

func ToTimeHMS(value *ole.VARIANT) time.Time {
	return UIntToTimeHMS(ToUInt64(value))
}

func CastSlice[T any](vArray []*ole.VARIANT, cast func(*ole.VARIANT) T) []T {
	ret := make([]T, len(vArray))
	for idx := 0; idx < len(ret); idx++ {
		ret[idx] = cast(vArray[idx])
	}
	return ret
}

func CombineDateAndTime(d time.Time, t time.Time) time.Time {
	hour, min, sec := t.Clock()
	return d.Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
}
