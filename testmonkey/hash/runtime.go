package main

import (
	"fmt"
	"unsafe"
)

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(unsafe.Pointer, uintptr, uintptr) uintptr

type stringStruct struct {
	str unsafe.Pointer
	len int
}

func MemHash(data []byte) uint64 {
	ss := (*stringStruct)(unsafe.Pointer(&data))
	return uint64(memhash(ss.str, 0, uintptr(ss.len)))
}

func main() {
	fmt.Println(MemHash([]byte("hello")))
}
