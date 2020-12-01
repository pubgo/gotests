package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pubgo/xerror"
	"xorm.io/xorm"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	xerror.Panic(err)
	fmt.Println(engine)
}
