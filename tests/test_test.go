package tests

import (
	"fmt"
	"github.com/pubgo/xerror"
	"net/url"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	u, err := url.Parse("etcdv3://etcd0.discov:2379,etcd1.discov:2379")
	xerror.Panic(err)
	fmt.Println(u.Host)

	fmt.Println(reflect.TypeOf(TestName).In(0).String())

	fmt.Println(reflect.TypeOf(&s1{}).Implements(reflect.TypeOf(init1).In(0)))
	fmt.Println(reflect.TypeOf(init1).In(0).String())
}

type ss interface {
	Hello()
}

type s1 struct {
}

func (*s1) Hello() {

}

func init1(s ss) {

}

func TestName11(t *testing.T) {
	fmt.Printf("%v",0.003)
}