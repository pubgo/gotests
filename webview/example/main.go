package main

import "github.com/webview/webview"

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Minimal webview example")
	w.SetSize(1000, 800, webview.HintFixed)
	w.Navigate("https://www.baidu.com")
	w.Run()
}
