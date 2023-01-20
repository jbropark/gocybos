package gocybos

import (
	"fmt"
	ole "github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32, _       = syscall.LoadLibrary("user32.dll")
	pPeekMessage, _ = syscall.GetProcAddress(user32, "PeekMessageW")
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

func PeekMessage(msg *ole.Msg, hwnd uint32, MsgFilterMin uint32, MsgFilterMax uint32, RemoveMsg uint32) (int32, error) {
	r0, _, err := syscall.SyscallN(pPeekMessage,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(MsgFilterMin),
		uintptr(MsgFilterMax),
		uintptr(RemoveMsg))

	return int32(r0), err
}

func PumpWaitingMessages() int32 {
	ret := int32(0)

	var msg ole.Msg

	runtime.LockOSThread()
	for {
		r, _ := PeekMessage(&msg, 0, 0, 0, 1)
		if r == 0 {
			break
		}
		if msg.Message == 0x0012 { // WM_QUIT
			ret = int32(1)
			break
		}
		ole.DispatchMessage(&msg)
	}
	runtime.UnlockOSThread()
	return ret
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

func VarToDate(value *ole.VARIANT) time.Time {
	return UIntToDate(ToUInt64(value))
}

func CastSlice[T any](vArray []*ole.VARIANT, cast func(*ole.VARIANT) T) []T {
	ret := make([]T, len(vArray))
	for idx := 0; idx < len(ret); idx++ {
		ret[idx] = cast(vArray[idx])
	}
	return ret
}
