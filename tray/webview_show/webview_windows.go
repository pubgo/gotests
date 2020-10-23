package webview_show

import (
	"fmt"
	"github.com/lxn/walk"
)

var (
	webView    *walk.WebView
	mainWindow *walk.MainWindow
	err        error
)

func RunWebview(title string, width, height int) {
	mainWindow, err = walk.NewMainWindow()
	if err != nil {
		panic(fmt.Sprintf("Failed to create main window: %v\n", err))
	}
	mainWindow.SetTitle(title)
	mainWindow.SetWidth(width)
	mainWindow.SetHeight(height)
	layout := walk.NewVBoxLayout()
	if err := mainWindow.SetLayout(layout); err != nil {
		panic(fmt.Sprintf("Failed to set layout: %v\n", err))
	}
	webView, err = walk.NewWebView(mainWindow)
	if err != nil {
		panic(fmt.Sprintf("Failed to create webview window: %v\n", err))
	}
	mainWindow.SetVisible(false)
	mainWindow.Run()
}

func ShowWebview(url string) {
	mainWindow.SetVisible(true)
	webView.SetURL(url)
}
