package main

import (
	"context"
	"fmt"
	"github.com/pubgo/gotests/testmonkey_patch/internal/a1"
	_ "google.golang.org/grpc"
	"net"
	"reflect"
	"unsafe"
	_ "unsafe"
)

type hh struct {
	h1 string
}

//go:linkname github.com/pubgo/gotests/testmonkey_patch/internal/a1.(*MyStruct).hello hello
func hello(t *a1.MyStruct) string {
	return "test hello ssss"
}

//go:linkname github.com/pubgo/gotests/testmonkey_patch/internal/a1.TestName TestName1111
func TestName1111() string {
	return "TestName1111"
}

//go:linkname init11 github.com/pubgo/gotests/testmonkey_patch/internal/a1.(*hh).init1
func init11(t *hh) string

//go:linkname init111 github.com/pubgo/gotests/testmonkey_patch/internal/a1.(*hh).init1
func init111(t *hh) string {
	return "nnn init111" + t.h1
}

func main() {
	b := a1.New("hello", "hello")
	fmt.Println(hello(b))
	fmt.Println(b.Hello())

	//fmt.Println(dial(context.Background(), func(_ context.Context, _ string) (net.Conn, error) {
	//	return nil, nil
	//}, ""))

	ss:=a1.Newhh("sss")
	dt := MyStructToBytes(unsafe.Pointer(ss))
	fmt.Println(BytesToMyStruct(dt).h1)
	hh1 := BytesToMyStruct(dt)
	fmt.Println(init11(hh1))
	fmt.Println(ss.Init1())

	fmt.Println(TestName1111())
	fmt.Println(a1.TestName())
}

//go:linkname dial google.golang.org/grpc/internal/transport.dial
func dial(ctx context.Context, fn func(context.Context, string) (net.Conn, error), addr string) (net.Conn, error)

var sizeOfMyStruct = int(unsafe.Sizeof(hh{}))

func MyStructToBytes(s unsafe.Pointer) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(s)
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *hh {
	return (*hh)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
