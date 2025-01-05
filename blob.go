package libsrt

// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type Blob interface {
	Pointer() unsafe.Pointer
	Size() C.int
}

type BlobWritable interface {
	Blob
	SizePointer() *C.int
}

type BlobInt int

func (i *BlobInt) Pointer() unsafe.Pointer {
	return unsafe.Pointer(i)
}

func (i *BlobInt) SizePointer() *C.int {
	dummySize := i.Size()
	return &dummySize
}

func (i BlobInt) Size() C.int {
	return C.int(unsafe.Sizeof(i))
}

type BlobStringReadonly string

func (s BlobStringReadonly) Pointer() unsafe.Pointer {
	var cString *C.char = C.CString(string(s))
	runtime.SetFinalizer(cString, func(cString *C.char) {
		C.free(unsafe.Pointer(cString))
	})
	return unsafe.Pointer(cString)
}

func (s BlobStringReadonly) Size() C.int {
	return C.int(len(s))
}
