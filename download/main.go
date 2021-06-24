package main

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/x/pkg/shutil"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pubgo/xerror"
)

type AutoGenerated struct {
	Data []struct {
		Dir           string    `json:"dir"`
		FileExtension string    `json:"file_extension"`
		FileID        string    `json:"file_id"`
		Name          string    `json:"name"`
		Type          string    `json:"type"`
		UpdatedAt     time.Time `json:"updated_at"`
		Category      string    `json:"category"`
		ContentType   string    `json:"content_type"`
		Size          int       `json:"size"`
		Password      string    `json:"password"`
		URL           string    `json:"url"`
	} `json:"data"`
	Meta struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"meta"`
}

func main() {
	// https://alist.now.sh/my/%E5%9B%BE%E7%89%87/a15b4afegy1fmvj99kec5j21hc0u0k7k.jpg
	resp, err := http.Post("https://api.xhofe.top/alist/api/path", "application/json", strings.NewReader(`{"path":"my/图片","password":""}`))
	xerror.Panic(err)

	dt, err := ioutil.ReadAll(resp.Body)
	xerror.Panic(err)

	fmt.Println(string(dt))
	var dd AutoGenerated
	xerror.Panic(json.Unmarshal(dt, &dd))

	for _, d := range dd.Data {
		shutil.Execute(fmt.Sprintf("open %s", d.Name))

	}
}