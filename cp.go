package gocybos

import (
	"github.com/go-ole/go-ole"
	"strings"
	"syscall"
	"unsafe"
)

var (
	// IID_CpCybosEvents, _ = ole.CLSIDFromString("{17F70631-56E5-40FC-B94F-44ADD3A850B1}") Not working
	IID_DibEvents, _ = ole.CLSIDFromString("{B8944520-09C3-11D4-8232-00105A7C4F8C}")
)

type CpTrait struct {
	name     string
	clsid    *ole.GUID
	unknown  *ole.IUnknown
	Object   *ole.Dispatch
	event    *ICpEvent
	point    *ole.IConnectionPoint
	callback CpReceiver
	cookie   uint32
}

type CpReceiver interface {
	Received(*CpTrait)
}

type ICpEvent struct {
	vTable *ICpEventVTable
	ref    int32
	host   *CpTrait
}

type ICpEventVTable struct {
	pQueryInterface   uintptr
	pAddRef           uintptr
	pRelease          uintptr
	pGetTypeInfoCount uintptr
	pGetTypeInfo      uintptr
	pGetIDsOfNames    uintptr
	pInvoke           uintptr
}

func (t *CpTrait) Create(name string) (err error) {
	t.clsid, err = ole.CLSIDFromString(name)
	if err != nil {
		return err
	}

	t.unknown, err = ole.CreateInstance(t.clsid, ole.IID_IUnknown)
	if err != nil {
		return err
	}

	iDispatch, err := t.unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}

	t.Object = &ole.Dispatch{Object: iDispatch}
	t.name = name
	return nil
}

func (t *CpTrait) Release() {
	// Release Dispatch
	if t.Object != nil {
		t.Object.Release()
		t.Object = nil
	}

	// Release IUnknown
	if t.unknown != nil {
		t.unknown.Release()
		t.unknown = nil
	}
}

// https://ippoeyeslhw.github.io/post/golang-with-cybosplus/
func createICpEvent(host *CpTrait) *ICpEvent {
	event := new(ICpEvent)
	event.vTable = new(ICpEventVTable)
	event.vTable.pQueryInterface = syscall.NewCallback(dispQueryInterface)
	event.vTable.pAddRef = syscall.NewCallback(dispAddRef)
	event.vTable.pRelease = syscall.NewCallback(dispRelease)
	event.vTable.pGetTypeInfoCount = syscall.NewCallback(dispGetTypeInfoCount)
	event.vTable.pGetTypeInfo = syscall.NewCallback(dispGetTypeInfo)
	event.vTable.pGetIDsOfNames = syscall.NewCallback(dispGetIDsOfNames)
	event.vTable.pInvoke = syscall.NewCallback(dispInvoke)
	event.host = host
	return event
}

// 이하 콜백 이벤트 바인딩하기 위한 함수 선언들
func dispQueryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(iid, ole.IID_IDispatch) ||
		ole.IsEqualGUID(iid, IID_DibEvents) {
		dispAddRef(this)
		*punk = this
		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func dispAddRef(this *ole.IUnknown) int32 {
	pthis := (*ICpEvent)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func dispRelease(this *ole.IUnknown) int32 {
	pthis := (*ICpEvent)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

func dispGetIDsOfNames(args *uintptr) uint32 {
	p := (*[6]int32)(unsafe.Pointer(args))
	//this := (*ole.IDispatch)(unsafe.Pointer(uintptr(p[0])))
	//iid := (*ole.GUID)(unsafe.Pointer(uintptr(p[1])))

	// wnames := *(*[]*uint16)(unsafe.Pointer(uintptr(p[2])))
	namelen := int(uintptr(p[3]))
	//lcid := int(uintptr(p[4]))
	pdisp := *(*[]int32)(unsafe.Pointer(uintptr(p[5])))
	for n := 0; n < namelen; n++ {
		pdisp[n] = int32(n)
	}
	return ole.S_OK
}

func dispGetTypeInfoCount(_ *ole.IUnknown, pcount *int) uint32 {
	if pcount != nil {
		*pcount = 0
	}
	return ole.S_OK
}

func dispGetTypeInfo(_ *ole.IUnknown, _ int, _ int) uint32 {
	return ole.E_NOTIMPL
}

func dispInvoke(this *ole.IDispatch, dispid int, _ *ole.GUID, _ int, _ int16, _ *ole.DISPPARAMS, _ *ole.VARIANT, _ *ole.EXCEPINFO, _ *uint) uintptr {
	pthis := (*ICpEvent)(unsafe.Pointer(this))
	if dispid == 1 {
		// 이벤트는 Received 한개뿐이다.
		if pthis.host.callback != nil {
			// instance callback
			pthis.host.callback.Received(pthis.host)
			return ole.S_OK
		}
	}
	return ole.E_NOTIMPL
}

func GetEventIID(name string) *ole.GUID {
	head := strings.SplitN(name, ".", 2)[0]
	if head != "Dscbo1" {
		panic("Cannot support event " + name)
	}

	return IID_DibEvents
}

func (t *CpTrait) BindEvent(callback CpReceiver) {
	eventIID := GetEventIID(t.name)

	if t.event == nil {
		t.event = createICpEvent(t)
	}
	t.callback = callback

	if t.point != nil {
		t.UnbindEvent()
	}

	unknown, err := t.Object.Object.QueryInterface(ole.IID_IConnectionPointContainer)
	if err != nil {
		panic(err)
	}
	container := (*ole.IConnectionPointContainer)(unknown)
	point := new(ole.IConnectionPoint)

	err = container.FindConnectionPoint(eventIID, &point)
	if err != nil {
		panic(err)
	}

	t.cookie, err = point.Advise((*ole.IUnknown)(unsafe.Pointer(t.event)))
	container.Release()
	if err != nil {
		point.Release()
		panic(err)
	}

	t.point = point
}

func (t *CpTrait) UnbindEvent() {
	if t.point == nil {
		return
	}

	err := t.point.Unadvise(t.cookie)
	if err != nil {
		panic(err)
	}

	t.point.Release()
	t.point = nil
}

// Implement Method and Property
// Refer: https://money2.daishin.com/e5/mboard/ptype_basic/HTS_Plus_Helper/DW_Basic_Read_Page.aspx?boardseq=284&seq=222&page=1&searchString=DibStatus&p=8839&v=8642&m=9508

func (t *CpTrait) SetInputValue(iType int32, value any) {
	t.Object.MustCall("SetInputValue", iType, value)
}

func (t *CpTrait) Request() {
	t.Object.MustCall("Request")
}

func (t *CpTrait) BlockRequest() *ole.VARIANT {
	return t.Object.MustCall("BlockRequest")
}

func (t *CpTrait) BlockRequest2(option int16) *ole.VARIANT {
	return t.Object.MustCall("BlockRequest2", option)
}

func (t *CpTrait) Subscribe() {
	t.Object.MustCall("Subscribe")
}

func (t *CpTrait) SubscribeLatest() {
	t.Object.MustCall("SubscribeLatest")
}

func (t *CpTrait) Unsubscribe() {
	t.Object.MustCall("Unsubscribe")
}

func (t *CpTrait) GetHeaderValue(hType int32) *ole.VARIANT {
	return t.Object.MustCall("GetHeaderValue", hType)
}

func (t *CpTrait) GetDataValue(dType int32, index int32) *ole.VARIANT {
	return t.Object.MustCall("GetDataValue", dType, index)
}

func (t *CpTrait) GetDibStatus() *ole.VARIANT {
	return t.Object.MustCall("GetDibStatus")
}

func (t *CpTrait) GetDibMsg1() *ole.VARIANT {
	return t.Object.MustCall("GetDibMsg1")
}

func (t *CpTrait) Continue() *ole.VARIANT {
	return t.Object.MustGet("Continue")
}

func (t *CpTrait) Header() *ole.VARIANT {
	return t.Object.MustGet("Header")
}

func (t *CpTrait) Data() *ole.VARIANT {
	return t.Object.MustGet("Data")
}