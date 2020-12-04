package main

import (
	"context"
	"fmt"
	"github.com/pubgo/gotests/testmonkey_patch/internal/a1"
	_ "google.golang.org/grpc"
	"net"
	_ "unsafe"
)

//go:linkname hello github.com/pubgo/gotests/testmonkey/internal/a1.(*MyStruct).hello
func hello(t *a1.MyStruct) string

func main() {
	b := a1.New("hello", "hello")
	fmt.Println(hello(b))

	fmt.Println(dial(context.Background(), func(_ context.Context, _ string) (net.Conn, error) {
		return nil, nil
	}, ""))
}

//go:linkname dial google.golang.org/grpc/internal/transport.dial
func dial(ctx context.Context, fn func(context.Context, string) (net.Conn, error), addr string) (net.Conn, error)
