package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("a")
	f.ImportName("b.c/d", "d")

	f.Func().Id("main").Params().Block(
		jen.Qual("b.c/d", "E").Call(),
	)
	fmt.Printf("%#v", f)
}
