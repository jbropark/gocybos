package gocybos

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

var (
	modOleAut32, _               = syscall.LoadDLL("oleaut32.dll")
	procSafeArrayCreateVector, _ = modOleAut32.FindProc("SafeArrayCreateVector")
	procSafeArrayPutElement, _   = modOleAut32.FindProc("SafeArrayPutElement")
)

func SafeArrayCreateVector(variantType ole.VT, lowerBound int32, length uint32) (safearray *ole.SafeArray, err error) {
	sa, _, err := procSafeArrayCreateVector.Call(
		uintptr(variantType),
		uintptr(lowerBound),
		uintptr(length))
	safearray = (*ole.SafeArray)(unsafe.Pointer(sa))
	return
}
func SafeArrayPutElement(safeArray *ole.SafeArray, index int64, element uintptr) {
	hr, _, err := procSafeArrayPutElement.Call(
		uintptr(unsafe.Pointer(safeArray)),
		uintptr(unsafe.Pointer(&index)),
		element)
	if hr != 0 {
		panic(err)
	}
}

func SafeArrayFromInt32Slice[T ~int32](slice []T) *ole.SafeArray {
	array, _ := SafeArrayCreateVector(ole.VT_I4, 0, uint32(len(slice)))

	if array == nil {
		panic("Could not convert []int32 to SafeArray")
	}
	// SysAllocStringLen(s)
	for i, v := range slice {
		SafeArrayPutElement(array, int64(i), uintptr(unsafe.Pointer(&v)))
	}
	return array
}

func VariantInt32Slice[T ~int32](s []T) ole.VARIANT {
	return ole.NewVariant(ole.VT_I4|ole.VT_ARRAY, int64(uintptr(unsafe.Pointer(SafeArrayFromInt32Slice(s)))))
}
