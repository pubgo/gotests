package main

import (
	"github.com/dave/jennifer/jen"
	proto_gen_lib "github.com/pubgo/gotests/proto-gen-lib"
	"github.com/pubgo/xerror"
	"log"
)

const (
	rpcPkgPath = "net/rpc"
)

func main() {
	defer xerror.Resp(func(err xerror.XErr) {
		log.Println(err.Println())
	})

	hello := proto_gen_lib.New("hello")
	hello.Parameter(func(key, value string) {
		log.Println(key, "sbhbhbsh", value)
	})

	xerror.Panic(hello.FileDescriptor(func(ss *proto_gen_lib.Service) {
		j := ss.J
		srv := ss.Name

		j.ImportAlias(rpcPkgPath, "rpc")

		j.Comment("// test")
		j.Comment("/* ssss */")
		j.Type().Id(srv).InterfaceFunc(func(group *jen.Group) {
			for _, m := range ss.GetMethod() {
				mthName := proto_gen_lib.CamelCase(m.GetName())
				group.Id(mthName).Params(
					jen.Id("in *"+ss.TypeName(m.GetInputType())),
					jen.Id("out *"+ss.TypeName(m.GetOutputType())),
				).Error()
			}
		})

		// method
		//srv *rpc.Server, x TestApiInterface
		j.Func().Id("Register"+srv).Params(
			jen.Id("srv *").Qual(rpcPkgPath, "Server"),
			jen.Id("x "+srv),
		).Error().BlockFunc(func(group *jen.Group) {
			group.Return().Nil()
		})
	}))
	xerror.Panic(hello.Save())
}
