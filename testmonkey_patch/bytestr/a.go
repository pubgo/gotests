package bytestr

import (
	"github.com/pubgo/gotests/testmonkey_patch/internal/a1"
	"reflect"
	"unsafe"
)

func MyStructToBytes(s *a1.MyStruct) []byte {
	var sizeOfMyStruct = int(unsafe.Sizeof(a1.MyStruct{}))

	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}

type MyStruct struct {
	a string
	b string
}