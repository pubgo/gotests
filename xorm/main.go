package main

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	sqlite "github.com/mattn/go-sqlite3"
	"github.com/pubgo/gotests/xorm/models"
	"github.com/pubgo/xerror"
	"xorm.io/xorm"
	"xorm.io/xorm/dialects"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func Floor(x float64) float64 {
	return math.Floor(x)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Rand() float32 {
	return rand.Float32()
}

func main() {
	sql.Register("sqlite3_custom", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("floor", Floor, true); err != nil {
				return err
			}

			if err := conn.RegisterFunc("rand", Rand, true); err != nil {
				return err
			}
			return nil
		},
	})
	dialects.RegisterDriver("sqlite3_custom", dialects.QueryDriver("sqlite3"))
	dialects.RegisterDialect("sqlite3_custom", func() dialects.Dialect { return dialects.QueryDialect("sqlite3") })

	//engine, err := xorm.NewEngine("sqlite3", "./test.db")
	engine, err := xorm.NewEngine("sqlite3","file:test.db?cache=shared&mode=memory")
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
