package main

import (
	"fmt"
	"github.com/pubgo/gotests/gorazor/examples/models"
	"github.com/pubgo/gotests/gorazor/examples/tpl"
)

func main() {
	user := &models.User{}
	user.Name = "go"
	user.Email = "hello@world.com"
	user.Intro = "I love gorazor!"
	user.IsImport = true
	fmt.Println(tpl.Home(1, user))
}
