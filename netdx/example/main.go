package main

import (
	_ "expvar"
	"github.com/pubgo/gotests/netdx"
	"net"
	"net/http"
)

func init() {
	_ = netdx.ShowTTY()
	netdx.L4Dx(true)
	netdx.L7Dx(true)
}

func main() {
	tl, _ := net.Listen("tcp", ":8088")
	http.Serve(tl, nil)
}
