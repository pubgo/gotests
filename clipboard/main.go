package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/pubgo/xerror"
	"log"
	"time"
)

func main() {
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("だんしゃり", 1.0)
	robotgo.Sleep(1)

	xerror.Panic(robotgo.WriteAll("日本語"))
	for range time.Tick(time.Second) {
		text := xerror.PanicStr(clipboard.ReadAll())
		if text != "" {
			log.Println("text is: ", text)
		}
		xerror.Panic(robotgo.WriteAll(text + " ok"))
		fmt.Println(robotgo.KeyTap("c", "command"))
	}
}
