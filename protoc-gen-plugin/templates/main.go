package main

import (
	"fmt"
	"github.com/pubgo/xerror"
	"os"
	"text/template"
)

type A struct {
	b *B
}

func (t *A) B(a string) *B {
	fmt.Println(a)
	return t.b
}

type B struct {
	name string
}

func (t *B) Name() string {
	return t.name
}

func main() {
	headerTemplate := template.Must(template.New("header").Parse(`
// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
var _ = utilities.NewDoubleArray
var _ = metadata.Join

{{ .B }}
{{ $b := .B "aa"}}
{{ $b.Name  }}

`))

	xerror.Exit(headerTemplate.Execute(os.Stderr, &A{b: &B{name: "hello"}}))
}
