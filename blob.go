package libsrt

// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type Blob interface {
	Pointer() unsafe.Pointer
	Size() uintptr
}

type BlobInt int

func (i *BlobInt) Pointer() unsafe.Pointer {
	return unsafe.Pointer(i)
}

func (i BlobInt) Size() uintptr {
	return unsafe.Sizeof(i)
}

type BlobStringReadonly string

func (s BlobStringReadonly) Pointer() unsafe.Pointer {
	var cString *C.char = C.CString(string(s))
	runtime.SetFinalizer(cString, func(cString *C.char) {
		C.free(unsafe.Pointer(cString))
	})
	return unsafe.Pointer(cString)
}

func (s BlobStringReadonly) Size() uintptr {
	return uintptr(len(s))
}
