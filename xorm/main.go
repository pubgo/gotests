package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pubgo/gotests/xorm/models"
	"github.com/pubgo/xerror"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	xerror.Panic(err)
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)

	engine.SetMapper(names.GonicMapper{})
	engine.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, "t_"))
	xerror.Panic(engine.Sync2(new(models.Article), new(models.Comment)))

	fmt.Printf("%#v\n", xerror.PanicErr(engine.DBMetas()))
	xerror.Panic(engine.DumpAll(os.Stdout))

	//fmt.Println(engine.Insert(&models.Article{Name: "hello"}))
	//fmt.Println(engine.Insert(&models.Comment{Name: "hello111"}))
	var a models.Article
	fmt.Println(engine.SQL("select * from t_article").Get(&a))
	fmt.Println(a)
}
