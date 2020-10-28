package main

import (
	"fmt"
	"github.com/pubgo/gotests/protoc-gen-plugin/examples/models"
	"github.com/pubgo/gotests/protoc-gen-plugin/examples/tpl"
	"net/http"
)

func main() {
	user := &models.User{}
	user.Name = "go"
	user.Email = "hello@world.com"
	user.Intro = "I love gorazor!"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, tpl.Home(1, user))
	})

	http.ListenAndServe(":8080", nil)
}
