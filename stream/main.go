package main

import (
	"fmt"
	"go.uber.org/atomic"
	"net/http"
	"time"
)

type Stream interface {
	Await() interface{}
	AChan() chan interface{}
	Cancel()
	Async(fn func(y Yield))
}

type Yield interface {
	Go(data interface{}) bool
}

func (s stream) Go(data interface{}) bool {
	if s.cancel.Load() {
		return true
	}

	s.data <- data
	return false
}

type stream struct {
	data   chan interface{}
	cancel atomic.Bool
}

func (s stream) Await() interface{} {
	return <-s.data
}

func (s stream) AChan() chan interface{} {
	return s.data
}

func (s stream) Cancel() {
	s.cancel.Store(true)
	close(s.data)
}

func (s stream) Async(fn func(y Yield)) {
	go fn(s)
}

func newStream() Stream {
	var s = &stream{data: make(chan interface{})}
	return s
}

func getData() Stream {
	s := newStream()
	for i := 10; i > 0; i-- {
		s.Async(func(y Yield) {
			fmt.Println("ss")
			resp, err := http.Get("https://www.oschina.net/news/124525/tokio-1-0-released")
			if err != nil {
				y.Go(err)
			}
			y.Go(resp)
		})
	}
	return s
}

func main() {
	ss := getData()
	for dt := range ss.AChan() {
		fmt.Println(dt)
		time.Sleep(time.Second)

		if dt == 5 {
			ss.Cancel()
		}
	}
}
