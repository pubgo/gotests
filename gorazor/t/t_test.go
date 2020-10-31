package t

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	var srv="sss"
	init1(srv)
}

func init1(name string) {
	fmt.Println(reflect.ValueOf(name))
}
