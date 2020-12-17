package main

import (
	"crypto/rand"
	"expvar"
	"github.com/pubgo/xerror"
	mr "math/rand"
	"net"
	"net/http"
	"time"
)

type mem struct {
	A int `json:"a"`
}

func init() {
	mr.Seed(time.Now().UnixNano())
}

// RangeBytes ...
func mockBytes(min, max int) []byte {
	var dt = make([]byte, mockInt(min, max))
	rand.Read(dt)
	return dt
}

// Range ...
func mockInt(min, max int) int {
	if min >= max {
		return max
	}
	return min + mr.Intn(max-min)
}

func main() {
	var m = &mem{}
	go func() {
		for {
			m.A = mockInt(10, 100)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	expvar.Publish("mem", expvar.Func(func() interface{} { return m }))
	l, err := net.Listen("tcp", ":8080")
	xerror.Panic(err)
	xerror.Panic(http.Serve(l, nil))
}
