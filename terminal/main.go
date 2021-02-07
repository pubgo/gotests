package main

import (
	"context"
	"fmt"
	glint "github.com/mitchellh/go-glint"
	"github.com/pubgo/xerror"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint32
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			atomic.AddUint32(&counter, 1)
		}
	}()

	d := glint.New()
	var dara string
	d.Append(
		glint.Style(
			glint.TextFunc(func(rows, cols uint) string {
				return fmt.Sprintf("%d tests passed", atomic.LoadUint32(&counter))
			}),
			glint.Color("green"),
		),
	)
	d.Append(
		glint.Style(
			glint.TextFunc(func(rows, cols uint) string {
				return fmt.Sprintf("%d tests passed %s", atomic.LoadUint32(&counter), dara)
			}),
			glint.Color("green"),
		),
	)
	go d.Render(context.Background())

	for {
		xerror.PanicErr(fmt.Scanln(&dara))
		fmt.Println()
		fmt.Println(dara)
		time.Sleep(time.Second)
	}
}
