package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/pubgo/xerror"
)

func TimeoutWith(dur time.Duration, fn func()) error {
	if dur < 0 {
		return xerror.Fmt("1")
	}

	if fn == nil {
		return xerror.Fmt("2")
	}

	var ch = make(chan error)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				switch err := err.(type) {
				case error:
					ch <- err
				default:
					ch <- fmt.Errorf("%s", err)
				}

				return
			}
			ch <- nil
		}()
		fn()
		ch <- nil
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(dur):
		return xerror.Fmt("3")
	}
}

func main() {
	for {
		//fmt.Println(TimeoutWith(time.Second, func() {
		//	time.Sleep(time.Second * 2)
		//}))

		fmt.Println(timeout(func() {
			time.Sleep(time.Second * 2)
		}))

		fmt.Println(runtime.NumGoroutine())

		time.Sleep(time.Second)
	}
}

func timeout(f func()) error {
	done := make(chan bool)
	go func() {
		f()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
