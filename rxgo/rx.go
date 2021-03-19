package main

import (
	"context"
	"fmt"
	rxgo "github.com/reactivex/rxgo/v2"
	"reflect"
	"unsafe"
)

func init1() {

}

func promise(fn func(ctx context.Context, next chan<- rxgo.Item)) rxgo.Observable {
	return rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		fn(ctx, next)
	}})
}

func main() {
	p := promise(func(ctx context.Context, next chan<- rxgo.Item) {

	})
	p.Map(func(ctx context.Context, i interface{}) (interface{}, error) {

	})

	p.First(func(ctx context.Context, i interface{}) (interface{}, error) {

	})

	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}

	fmt.Println(unpackEface(init1))
	fmt.Println(unpackEface(init1))
	fmt.Println(reflect.ValueOf(init1).Pointer())
	fmt.Println(reflect.ValueOf(init1).Pointer())
	fmt.Println(reflect.ValueOf(init1).Pointer())

	go func() {
		fmt.Println(1, " ", reflect.ValueOf(init1).Pointer())
		fmt.Println(1, " ", reflect.ValueOf(init1).Pointer())
	}()

	fmt.Printf("%p", init1)
	fmt.Printf("%p", init1)
	fmt.Printf("%p", init1)

	var data = make(chan []reflect.Value, 1)
	data = nil
	select {
	case <-data:
		fmt.Println("ok")
	}
}

// unpackEface converts the empty interface i to a Value.
func unpackEface(i interface{}) uintptr {
	return uintptr(unsafe.Pointer(&i))
}
