package main

import (
	"fmt"
	"github.com/pubgo/gotests/testmonkey/internal/a1"
	_ "unsafe"
)

//go:linkname hello github.com/pubgo/gotests/testmonkey/internal/a1.(*MyStruct).hello
func hello(t *a1.MyStruct) string

func main() {
	b := a1.New("hello", "hello")
	fmt.Println(hello(b))
}
