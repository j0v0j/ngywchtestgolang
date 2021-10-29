package libdll

import (
	"fmt"
	"sync"
	"syscall"
	"unsafe"
)

var dll *syscall.DLL

// var err error

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func load_dll(libname string) {
	dll = syscall.MustLoadDLL(libname)
	fmt.Println("+++++++syscall.LoadLibrary:", dll, "+++++++")
	// defer syscall.FreeLibrary(dll)
	VCI_OpenDevice := dll.MustFindProc("VCI_OpenDevice")

	// for i := 0; i < 10; i++ {

	ret, _, err := VCI_OpenDevice.Call(IntPtr(4), IntPtr(0))
	if err != nil {
		fmt.Printf("%v ：%v \r\n", libname, ret)
	}

	// }

}

func Usedll(ws *sync.WaitGroup) {
	load_dll("ControlCAN.dll")
	ws.Done()
}
