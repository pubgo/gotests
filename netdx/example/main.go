package main

import (
	"github.com/pubgo/gotests/netdx"
	"net"
	"net/http"
)

func init() {
	_ = netdx.ShowTTY()
	netdx.L4Dx(true)
}

func main() {
	tl, _ := net.Listen("tcp", ":8088")
	http.Serve(tl, nil)
}
