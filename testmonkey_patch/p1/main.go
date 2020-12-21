package main

import (
	"fmt"
	"reflect"
)

func init1() {
	fmt.Println("hello")
}

func init2() {
	fmt.Println("init2")
}

func main() {
	init1()

	ri := reflect.ValueOf(init1)
	fmt.Println(ri.CanSet(), ri.Pointer())

}
