package goleak

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestName(t *testing.T) {
	defer goleak.VerifyNone(t)

	fmt.Println(goWithTimeout(time.Second*1, func() {
		time.Sleep(time.Second * 2)
	}))
	fmt.Println("ok")
	time.Sleep(time.Second * 2)
}

func goWithTimeout(dur time.Duration, fn func()) string {
	var ch = make(chan struct{})
	go func() {
		defer func() {
			ch <- struct{}{}
		}()

		fn()
		fmt.Println("over")
	}()

	select {
	case <-ch:
		return "ok1"
	case <-time.After(dur):
		return "timeout"
	}
}
