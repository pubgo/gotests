package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"net/url"
)

func main() {
	systray.Register(onReady, nil)
	configureWebview("Webview example", 200, 100)
}
var indexHTML = `
<!DOCTYPE html>
<html>
<head>
<title>测试 - 幕布</title>
<meta charset="utf-8"/>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="renderer" content="webkit"/>
<meta name="author" content="mubu.com"/>
</head>
<body style="margin: 50px 20px;color: #333;font-family: SourceSansPro,-apple-system,BlinkMacSystemFont,'PingFang SC',Helvetica,Arial,'Microsoft YaHei',微软雅黑,黑体,Heiti,sans-serif,SimSun,宋体,serif">
<div class="export-wrapper"><div style="font-size: 22px; padding: 0 15px 0;"><div style="padding-bottom: 24px">测试</div><div style="background: #e5e6e8; height: 1px; margin-bottom: 20px;"></div></div><ul style="list-style: disc outside;"><li style="line-height: 34px;"><span class="content mubu-node" heading="1" style="line-height: 34px; min-height: 34px; font-size: 24px; padding: 2px 0px; display: inline-block; vertical-align: top;"># sssnj</span></li><li style="line-height: 34px;"><span class="content mubu-node" heading="1" style="line-height: 34px; min-height: 34px; font-size: 24px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold" style="font-weight: bold;">heade1</span></span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 30px;"><span class="content mubu-node" heading="2" style="line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;">heade2</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 27px;"><span class="content mubu-node" heading="3" style="line-height: 27px; min-height: 27px; font-size: 19px; padding: 2px 0px; display: inline-block; vertical-align: top;">heade3</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" color="#dc2d1e" style="color: rgb(220, 45, 30); line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">三生三世</span></li></ul></li></ul></li></ul></li><li style="line-height: 30px;"><span class="content mubu-node" heading="2" style="line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold" style="font-weight: bold;">heade2</span></span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 27px;"><span class="content mubu-node" heading="3" style="line-height: 27px; min-height: 27px; font-size: 19px; padding: 2px 0px; display: inline-block; vertical-align: top;">heade3</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold" style="font-weight: bold;">三生三世</span></span></li></ul></li></ul></li><li style="line-height: 27px;"><span class="content mubu-node" heading="3" style="line-height: 27px; min-height: 27px; font-size: 19px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold" style="font-weight: bold;">heade3</span></span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold underline" style="font-weight: bold; text-decoration: underline;">三生三世</span></span></li></ul></li><li style="line-height: 24px;"><span class="content mubu-node" color="#dc2d1e" style="color: rgb(220, 45, 30); line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">三生三世</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" color="#dc2d1e" style="color: rgb(220, 45, 30); line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">hello</span></li></ul></li><li style="line-height: 30px;"><span class="content mubu-node" color="#333333" heading="2" style="color: rgb(51, 51, 51); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;">测试</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 30px;"><span class="content mubu-node" color="#3da8f5" heading="2" images="%5B%7B%22id%22%3A%221d916f3267c0b118a-40263%22%2C%22oh%22%3A1004%2C%22ow%22%3A742%2C%22uri%22%3A%22document_image%2F7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg%22%2C%22w%22%3A87%7D%5D" style="color: rgb(61, 168, 245); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="bold italic underline" style="font-weight: bold; text-decoration: underline; font-style: italic;">测试图片</span></span><div style="padding: 3px 0"><img src="https://img.mubu.com/document_image/7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg" style="max-width: 720px; width: 87px;" class="attach-img"></div></li><li style="line-height: 30px;"><span class="content mubu-node" color="#3da8f5" heading="2" style="color: rgb(61, 168, 245); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;">是是是</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 30px;"><span class="content mubu-node" color="#3da8f5" heading="2" style="color: rgb(61, 168, 245); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;">ssss</span></li><li style="line-height: 30px;"><span class="content mubu-node" color="#dc2d1e" heading="2" style="color: rgb(220, 45, 30); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"><a class="content-link" target="_blank" href="https://mubu.com/doclcoXBPA2D" style="text-decoration: underline; opacity: 0.6; color: inherit;"><span class="bold italic" style="font-weight: bold; font-style: italic;">https://mubu.com/doclcoXBPA2D</span></a></span></li><li style="line-height: 24px;"><span class="content mubu-node" color="#dc2d1e" images="%5B%7B%22id%22%3A%2237d16f3289cb16101%22%2C%22oh%22%3A1004%2C%22ow%22%3A742%2C%22uri%22%3A%22document_image%2F7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg%22%2C%22w%22%3A87%7D%5D" style="color: rgb(220, 45, 30); line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;"><a class="content-link" target="_blank" href="https://mubu.com/doclcoXBPA2D" style="text-decoration: underline; opacity: 0.6; color: inherit;"><span class="bold italic" style="font-weight: bold; font-style: italic;">https://mubu.com/doclcoXBPA2D</span></a></span><div style="padding: 3px 0"><img src="https://img.mubu.com/document_image/7fabd28a-8c59-4ffe-b9f3-ab2ef4c91549-40263.jpg" style="max-width: 720px; width: 87px;" class="attach-img"></div></li><li style="line-height: 30px;"><span class="content mubu-node" color="#dc2d1e" heading="2" style="color: rgb(220, 45, 30); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"><span class="italic" style="font-style: italic;">sss</span></span></li></ul></li></ul></li><li style="line-height: 30px;"><span class="content mubu-node" color="#dc2d1e" heading="2" style="color: rgb(220, 45, 30); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"></span></li><li style="line-height: 30px;"><span class="content mubu-node" color="#dc2d1e" heading="2" style="color: rgb(220, 45, 30); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;">ok</span></li><li style="line-height: 30px;"><span class="content mubu-node" color="#dc2d1e" heading="2" style="color: rgb(220, 45, 30); line-height: 30px; min-height: 30px; font-size: 21px; padding: 2px 0px; display: inline-block; vertical-align: top;"><a class="content-link" target="_blank" href="https://github.com/alash3al/redix" style="text-decoration: underline; opacity: 0.6; color: inherit;">https://github.com/alash3al/redix</a></span></li><li style="line-height: 27px;"><span class="content mubu-node" color="#dc2d1e" heading="3" style="color: rgb(220, 45, 30); line-height: 27px; min-height: 27px; font-size: 19px; padding: 2px 0px; display: inline-block; vertical-align: top;">标签</span></li></ul></div>

<audio controls="controls">
  <source src="./horse.mp3" type="audio/mp3" />
Your browser does not support this audio format.
</audio>
</body>
</html>
`
func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Webview example")
	mShowLantern := systray.AddMenuItem("Show Lantern", "")
	mShowWikipedia := systray.AddMenuItem("Show Wikipedia", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShowLantern.ClickedCh:
				//showWebview("https://www.baidu.com")
				showWebview("data:text/html," + url.PathEscape(indexHTML))
			case <-mShowWikipedia.ClickedCh:
				showWebview("https://www.wikipedia.org")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}
