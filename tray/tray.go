package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"time"
)
import "github.com/getlantern/systray"

func main() {
	onExit := func() {
		fmt.Println("Starting onExit")
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
		fmt.Println("Finished onExit")
	}
	// Should be called at the very beginning of main().
	systray.RunWithAppWindow("Lantern", 1024, 768, onReady, onExit)
}

func onReady() {
	systray.SetIcon(Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")

		mEnabled.SetTooltip("测试tooltip")

		systray.AddMenuItem("Ignored", "Ignored")
		mUrl := systray.AddMenuItem("Open UI", "my home")
		mQuit := systray.AddMenuItem("退出", "Quit the whole app")

		// Sets the icon of a menu item. Only available on Mac.
		mQuit.SetIcon(Data)

		//mQuit.Hide()

		systray.AddSeparator()
		mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
		shown := true
		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-mUrl.ClickedCh:
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

</body>
</html>
`
				//systray.ShowAppWindow("https://www.github.com/getlantern/lantern")
				systray.ShowAppWindow("data:text/html," + url.PathEscape(indexHTML),)
			case <-mToggle.ClickedCh:
				if shown {
					mQuitOrig.Hide()
					mEnabled.Hide()
					shown = false
				} else {
					mQuitOrig.Show()
					mEnabled.Show()
					shown = true
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
var Data []byte = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x20,
	0x08, 0x06, 0x00, 0x00, 0x00, 0x73, 0x7a, 0x7a, 0xf4, 0x00, 0x00, 0x00,
	0x19, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x00, 0x41, 0x64, 0x6f, 0x62, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x79, 0x71, 0xc9, 0x65, 0x3c, 0x00, 0x00,
	0x03, 0x66, 0x69, 0x54, 0x58, 0x74, 0x58, 0x4d, 0x4c, 0x3a, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x2e, 0x78, 0x6d, 0x70, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x3c, 0x3f, 0x78, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x20, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x3d, 0x22, 0xef, 0xbb, 0xbf,
	0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x57, 0x35, 0x4d, 0x30, 0x4d, 0x70,
	0x43, 0x65, 0x68, 0x69, 0x48, 0x7a, 0x72, 0x65, 0x53, 0x7a, 0x4e, 0x54,
	0x63, 0x7a, 0x6b, 0x63, 0x39, 0x64, 0x22, 0x3f, 0x3e, 0x20, 0x3c, 0x78,
	0x3a, 0x78, 0x6d, 0x70, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x78, 0x6d, 0x6c,
	0x6e, 0x73, 0x3a, 0x78, 0x3d, 0x22, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x3a,
	0x6e, 0x73, 0x3a, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x22, 0x20, 0x78, 0x3a,
	0x78, 0x6d, 0x70, 0x74, 0x6b, 0x3d, 0x22, 0x41, 0x64, 0x6f, 0x62, 0x65,
	0x20, 0x58, 0x4d, 0x50, 0x20, 0x43, 0x6f, 0x72, 0x65, 0x20, 0x35, 0x2e,
	0x30, 0x2d, 0x63, 0x30, 0x36, 0x30, 0x20, 0x36, 0x31, 0x2e, 0x31, 0x33,
	0x34, 0x37, 0x37, 0x37, 0x2c, 0x20, 0x32, 0x30, 0x31, 0x30, 0x2f, 0x30,
	0x32, 0x2f, 0x31, 0x32, 0x2d, 0x31, 0x37, 0x3a, 0x33, 0x32, 0x3a, 0x30,
	0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x3e, 0x20,
	0x3c, 0x72, 0x64, 0x66, 0x3a, 0x52, 0x44, 0x46, 0x20, 0x78, 0x6d, 0x6c,
	0x6e, 0x73, 0x3a, 0x72, 0x64, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77, 0x33, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x31, 0x39, 0x39, 0x39, 0x2f, 0x30, 0x32, 0x2f, 0x32, 0x32,
	0x2d, 0x72, 0x64, 0x66, 0x2d, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2d,
	0x6e, 0x73, 0x23, 0x22, 0x3e, 0x20, 0x3c, 0x72, 0x64, 0x66, 0x3a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72,
	0x64, 0x66, 0x3a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x3d, 0x22, 0x22, 0x20,
	0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3d,
	0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61,
	0x64, 0x6f, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70,
	0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6d, 0x6d, 0x2f, 0x22, 0x20, 0x78, 0x6d,
	0x6c, 0x6e, 0x73, 0x3a, 0x73, 0x74, 0x52, 0x65, 0x66, 0x3d, 0x22, 0x68,
	0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f,
	0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31,
	0x2e, 0x30, 0x2f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x23, 0x22, 0x20, 0x78,
	0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d, 0x70, 0x3d, 0x22, 0x68, 0x74,
	0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f, 0x62,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31, 0x2e,
	0x30, 0x2f, 0x22, 0x20, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x64, 0x69,
	0x64, 0x3a, 0x36, 0x37, 0x32, 0x34, 0x42, 0x45, 0x31, 0x35, 0x45, 0x44,
	0x32, 0x30, 0x36, 0x38, 0x31, 0x31, 0x38, 0x38, 0x43, 0x36, 0x46, 0x32,
	0x38, 0x31, 0x35, 0x44, 0x41, 0x33, 0x43, 0x35, 0x35, 0x35, 0x22, 0x20,
	0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x64, 0x69,
	0x64, 0x3a, 0x41, 0x33, 0x42, 0x34, 0x46, 0x42, 0x36, 0x36, 0x33, 0x41,
	0x41, 0x38, 0x31, 0x31, 0x45, 0x32, 0x42, 0x32, 0x43, 0x41, 0x39, 0x37,
	0x42, 0x44, 0x33, 0x34, 0x34, 0x31, 0x45, 0x46, 0x33, 0x32, 0x22, 0x20,
	0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69,
	0x64, 0x3a, 0x41, 0x33, 0x42, 0x34, 0x46, 0x42, 0x36, 0x35, 0x33, 0x41,
	0x41, 0x38, 0x31, 0x31, 0x45, 0x32, 0x42, 0x32, 0x43, 0x41, 0x39, 0x37,
	0x42, 0x44, 0x33, 0x34, 0x34, 0x31, 0x45, 0x46, 0x33, 0x32, 0x22, 0x20,
	0x78, 0x6d, 0x70, 0x3a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x54,
	0x6f, 0x6f, 0x6c, 0x3d, 0x22, 0x41, 0x64, 0x6f, 0x62, 0x65, 0x20, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x20, 0x43, 0x53, 0x35,
	0x20, 0x4d, 0x61, 0x63, 0x69, 0x6e, 0x74, 0x6f, 0x73, 0x68, 0x22, 0x3e,
	0x20, 0x3c, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x44, 0x65, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x3a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69, 0x64, 0x3a, 0x45, 0x36,
	0x38, 0x31, 0x34, 0x43, 0x36, 0x41, 0x45, 0x45, 0x32, 0x30, 0x36, 0x38,
	0x31, 0x31, 0x38, 0x38, 0x43, 0x36, 0x46, 0x32, 0x38, 0x31, 0x35, 0x44,
	0x41, 0x33, 0x43, 0x35, 0x35, 0x35, 0x22, 0x20, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x3a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x64, 0x3a, 0x36, 0x37,
	0x32, 0x34, 0x42, 0x45, 0x31, 0x35, 0x45, 0x44, 0x32, 0x30, 0x36, 0x38,
	0x31, 0x31, 0x38, 0x38, 0x43, 0x36, 0x46, 0x32, 0x38, 0x31, 0x35, 0x44,
	0x41, 0x33, 0x43, 0x35, 0x35, 0x35, 0x22, 0x2f, 0x3e, 0x20, 0x3c, 0x2f,
	0x72, 0x64, 0x66, 0x3a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x3e, 0x20, 0x3c, 0x2f, 0x72, 0x64, 0x66, 0x3a, 0x52,
	0x44, 0x46, 0x3e, 0x20, 0x3c, 0x2f, 0x78, 0x3a, 0x78, 0x6d, 0x70, 0x6d,
	0x65, 0x74, 0x61, 0x3e, 0x20, 0x3c, 0x3f, 0x78, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x3d, 0x22, 0x72, 0x22, 0x3f, 0x3e,
	0x5d, 0xed, 0x35, 0xe2, 0x00, 0x00, 0x04, 0xee, 0x49, 0x44, 0x41, 0x54,
	0x78, 0xda, 0xc4, 0x57, 0xcf, 0x6f, 0x55, 0x45, 0x18, 0x3d, 0xf3, 0xe3,
	0xfe, 0xea, 0x7b, 0xaf, 0xa5, 0x6d, 0x0a, 0xd8, 0x34, 0xbe, 0x16, 0x83,
	0x69, 0x8c, 0x2e, 0x04, 0xe2, 0x86, 0xb8, 0x70, 0xe1, 0x06, 0x35, 0x18,
	0x13, 0x5d, 0x60, 0x8c, 0xd1, 0x68, 0xe2, 0xca, 0xb8, 0x33, 0x31, 0xf1,
	0x6f, 0x70, 0x67, 0x5c, 0xb1, 0x62, 0xe1, 0x46, 0x42, 0x8c, 0x0b, 0xe3,
	0x46, 0x34, 0x25, 0x11, 0x41, 0x14, 0xa4, 0x24, 0xa4, 0x08, 0x58, 0x0a,
	0x29, 0x14, 0x0a, 0x6d, 0xe9, 0xeb, 0xbb, 0xef, 0xce, 0x9d, 0xf1, 0xcc,
	0xbd, 0xaf, 0xa5, 0x44, 0x63, 0x49, 0xee, 0x4b, 0x78, 0xc9, 0xf7, 0xee,
	0x9d, 0x3b, 0x33, 0x77, 0xce, 0x77, 0xbe, 0xf3, 0x7d, 0x33, 0x57, 0x38,
	0xe7, 0xf0, 0x38, 0x7f, 0x7a, 0xab, 0x01, 0xe2, 0xd9, 0x37, 0xff, 0xeb,
	0xb1, 0xa2, 0x7d, 0x46, 0xdb, 0xeb, 0x87, 0xd0, 0x8e, 0xd3, 0xbe, 0xf8,
	0xd7, 0x28, 0x6b, 0xe1, 0x2e, 0x7c, 0xf3, 0xbf, 0xef, 0x97, 0x5b, 0x42,
	0x34, 0x06, 0xf0, 0x2c, 0x6d, 0x36, 0xe0, 0x43, 0xda, 0x88, 0x90, 0xf2,
	0x90, 0xea, 0x6f, 0xbc, 0x0b, 0x21, 0x9e, 0x67, 0xfb, 0x8d, 0xa2, 0xcf,
	0x76, 0xc7, 0x70, 0x5e, 0xff, 0x40, 0x7f, 0x75, 0x06, 0xc6, 0x27, 0x9a,
	0xb8, 0x76, 0x63, 0xbe, 0x70, 0x53, 0xf0, 0x2f, 0xe7, 0x02, 0xd6, 0xda,
	0x71, 0x36, 0xaf, 0xd1, 0x0e, 0xcb, 0x38, 0x16, 0x5c, 0xf4, 0x7a, 0xbe,
	0xbc, 0xd2, 0x14, 0x71, 0x04, 0x19, 0x68, 0xf8, 0xb0, 0x4a, 0xda, 0x2e,
	0xce, 0xad, 0x0c, 0x60, 0xf7, 0x53, 0xbb, 0x90, 0x87, 0x11, 0x76, 0x8f,
	0xee, 0x40, 0xa8, 0x35, 0xfe, 0x9a, 0xbf, 0x35, 0x36, 0x73, 0xf9, 0xea,
	0x24, 0xd2, 0xf4, 0x20, 0x5d, 0x2d, 0xbc, 0x15, 0x49, 0x0c, 0x1d, 0x86,
	0xdf, 0x09, 0x25, 0x7f, 0x80, 0x14, 0xd3, 0x8e, 0x20, 0x93, 0x30, 0x44,
	0xa3, 0xd1, 0xd8, 0x12, 0x80, 0xdc, 0x3a, 0x02, 0x06, 0xa4, 0xba, 0xb0,
	0x5a, 0x12, 0x2b, 0x2e, 0xf4, 0x35, 0xb4, 0x3e, 0x58, 0x84, 0xde, 0xb3,
	0x91, 0xa6, 0x64, 0x86, 0xf7, 0x4a, 0xbe, 0x4a, 0xcf, 0x8f, 0xe4, 0x26,
	0x0f, 0x42, 0xa5, 0xf0, 0xcc, 0xe8, 0x13, 0x50, 0xfe, 0x79, 0x65, 0x11,
	0xf2, 0x1d, 0x86, 0x62, 0x9a, 0x9a, 0xbe, 0x88, 0xa7, 0x77, 0x8e, 0x04,
	0x9d, 0x34, 0x1b, 0x45, 0xda, 0x81, 0xf7, 0xde, 0x53, 0x9d, 0x77, 0x32,
	0x04, 0x49, 0x82, 0x88, 0x1e, 0x67, 0xb9, 0xa9, 0x37, 0xe2, 0x44, 0x3c,
	0x39, 0x3c, 0x84, 0xa8, 0x1b, 0x8a, 0xca, 0x00, 0xba, 0xbf, 0x28, 0x35,
	0xe6, 0xc0, 0x9f, 0x17, 0x2f, 0x7d, 0x20, 0xad, 0x6d, 0xc2, 0xe4, 0xd8,
	0x10, 0x45, 0x10, 0x20, 0xd0, 0x01, 0xfa, 0x09, 0xa2, 0x16, 0x85, 0x13,
	0xf5, 0x28, 0x3a, 0x1a, 0x28, 0x75, 0x98, 0x4b, 0x7f, 0xcf, 0xde, 0x56,
	0xe5, 0x10, 0xf0, 0xf7, 0x36, 0x6d, 0x4a, 0x0a, 0x71, 0x14, 0x4a, 0x1d,
	0xb0, 0x7e, 0xce, 0x3a, 0xb5, 0xbc, 0x2a, 0xea, 0x22, 0x50, 0x92, 0x11,
	0x90, 0xd0, 0x8a, 0xdc, 0x0b, 0xbc, 0xc2, 0x1e, 0x9f, 0x7b, 0xbf, 0xd0,
	0x3e, 0xea, 0x05, 0x80, 0x8f, 0x69, 0xfb, 0x7c, 0x76, 0x81, 0x8e, 0x23,
	0x75, 0xa5, 0x75, 0x31, 0x68, 0x0f, 0x80, 0x66, 0xac, 0x44, 0x9a, 0x09,
	0x38, 0x5b, 0xbe, 0x92, 0xdd, 0xcf, 0xd1, 0xde, 0xab, 0x1c, 0x82, 0x95,
	0x54, 0xdf, 0x58, 0x5a, 0xe3, 0xab, 0x3a, 0x0e, 0xf5, 0xc8, 0x61, 0x72,
	0x4c, 0x53, 0x5c, 0x0e, 0x27, 0x67, 0xb2, 0x62, 0x76, 0x28, 0x55, 0x51,
	0x97, 0xf6, 0x8c, 0x4b, 0xf4, 0x45, 0xc0, 0xad, 0x7b, 0xe4, 0xbd, 0x23,
	0xb0, 0xda, 0x21, 0x51, 0x10, 0x57, 0x2a, 0x03, 0x38, 0xb4, 0xf7, 0xef,
	0x99, 0xe0, 0x85, 0x35, 0x34, 0x87, 0x15, 0x26, 0x77, 0x0e, 0xa3, 0x39,
	0x5e, 0x73, 0xc7, 0x7e, 0x5a, 0xc5, 0x5b, 0x9f, 0xcf, 0x09, 0xd4, 0x15,
	0x19, 0x50, 0x94, 0x84, 0xc4, 0xfb, 0x2f, 0x25, 0x78, 0x6d, 0x9f, 0xc4,
	0xed, 0x85, 0x0c, 0x73, 0xf7, 0x52, 0xcc, 0x2e, 0x38, 0x5c, 0x5f, 0x74,
	0xe7, 0x2a, 0x03, 0xf8, 0xe4, 0xe5, 0x9b, 0xe7, 0xa0, 0xb7, 0xd1, 0xc9,
	0x41, 0x0a, 0xbf, 0x8f, 0x2e, 0xf7, 0x09, 0xa9, 0x38, 0x4d, 0xcc, 0x16,
	0x9e, 0xfb, 0xb0, 0x07, 0x5a, 0x50, 0x8b, 0x75, 0x48, 0x1d, 0x61, 0x64,
	0xb0, 0x8d, 0xed, 0x43, 0x6d, 0xec, 0x99, 0xa0, 0xfe, 0x4c, 0x6b, 0xba,
	0x32, 0x80, 0xd4, 0xec, 0xb8, 0x2c, 0x5c, 0x9d, 0x45, 0xbd, 0x21, 0x21,
	0x6b, 0x2c, 0x46, 0x35, 0xa6, 0x9d, 0x7d, 0xa0, 0x01, 0x0a, 0x30, 0x24,
	0x0b, 0x51, 0x5c, 0x67, 0x23, 0x41, 0x26, 0x42, 0x6a, 0xc5, 0x9b, 0x36,
	0x70, 0xe1, 0xb5, 0xb0, 0x72, 0x1d, 0x08, 0x86, 0x66, 0xa1, 0xe2, 0x25,
	0xe8, 0x81, 0x41, 0xa1, 0x6a, 0x64, 0xa0, 0x1f, 0x41, 0xdc, 0x59, 0xef,
	0xa5, 0xfa, 0x15, 0x42, 0x0f, 0xa2, 0x00, 0xe0, 0x59, 0x08, 0xe0, 0xf8,
	0xcc, 0x09, 0x71, 0x9b, 0x20, 0x66, 0xab, 0xd7, 0x01, 0x99, 0xdc, 0xa4,
	0xe7, 0x57, 0x84, 0x6e, 0x0c, 0xd2, 0x98, 0xf7, 0x83, 0xf4, 0x76, 0x6d,
	0x23, 0x7f, 0x7c, 0x0a, 0xfa, 0x10, 0xc4, 0x31, 0xfb, 0x14, 0x37, 0x1f,
	0x4d, 0xf1, 0x51, 0x13, 0xac, 0x42, 0x97, 0x9d, 0x50, 0x8b, 0xd5, 0x2b,
	0x61, 0x30, 0x90, 0x41, 0x86, 0x97, 0xe8, 0xfd, 0x9e, 0x02, 0x80, 0xda,
	0x46, 0x6f, 0x57, 0x8b, 0x52, 0xe0, 0x33, 0xd3, 0xe7, 0x3f, 0x0b, 0x0f,
	0xa2, 0x64, 0x80, 0x8d, 0x3a, 0x84, 0x66, 0x85, 0x2c, 0xc2, 0x93, 0xcf,
	0x08, 0x17, 0xd9, 0xea, 0x75, 0x40, 0x51, 0x78, 0x32, 0xfa, 0x83, 0x61,
	0x80, 0xf0, 0xf7, 0x5c, 0x24, 0x8c, 0x06, 0x20, 0x65, 0x39, 0xd5, 0xd7,
	0xfb, 0x52, 0x03, 0x04, 0xa7, 0xfd, 0xd8, 0x84, 0xe0, 0x22, 0xee, 0x1d,
	0xd1, 0x19, 0x7f, 0xad, 0xce, 0x80, 0xf2, 0x2f, 0x91, 0x67, 0x79, 0xe3,
	0xe9, 0xf0, 0x60, 0x10, 0x84, 0x0d, 0x36, 0x65, 0xb1, 0x1f, 0x48, 0xd1,
	0x65, 0x20, 0x68, 0xa0, 0x18, 0x23, 0x03, 0x7f, 0x65, 0x47, 0x78, 0x1e,
	0xc2, 0x55, 0xdf, 0x0d, 0xfd, 0x82, 0x7c, 0xe9, 0x79, 0xde, 0x2d, 0x95,
	0xdb, 0xaf, 0x45, 0x18, 0xf8, 0xf4, 0x2b, 0xa7, 0x7a, 0x22, 0x22, 0xb6,
	0x83, 0x90, 0x0b, 0x93, 0xfb, 0x32, 0x39, 0xe4, 0x02, 0x41, 0x9c, 0x2f,
	0xc0, 0xf4, 0xa0, 0x14, 0x7b, 0x4f, 0xe7, 0xe0, 0xb2, 0x0b, 0xb0, 0x54,
	0x7f, 0xbe, 0x46, 0x00, 0xe9, 0x03, 0x00, 0x3e, 0x04, 0xdc, 0xf9, 0x22,
	0xc5, 0xca, 0xc8, 0x7e, 0xe7, 0x7c, 0xbd, 0xce, 0x09, 0x58, 0x2c, 0x6c,
	0xe4, 0x6a, 0x25, 0x00, 0x8e, 0x2f, 0x76, 0xc6, 0x72, 0xe3, 0x3f, 0xed,
	0xf2, 0x55, 0x96, 0xe4, 0x65, 0xc4, 0xaa, 0xc5, 0xb8, 0xcb, 0x82, 0x10,
	0x81, 0x32, 0x0b, 0x22, 0xc1, 0xcc, 0x30, 0x34, 0xdb, 0x26, 0x88, 0xec,
	0xd7, 0x62, 0x2f, 0x76, 0xb6, 0x3a, 0x00, 0x97, 0x67, 0xa5, 0x99, 0xd6,
	0x94, 0x33, 0xcb, 0x04, 0x70, 0x17, 0x7d, 0x62, 0x85, 0x27, 0x9e, 0x2e,
	0x80, 0x42, 0x84, 0xac, 0x07, 0xee, 0x3e, 0x01, 0xac, 0x70, 0x6c, 0xcb,
	0x33, 0x71, 0x82, 0x13, 0x50, 0x58, 0xe5, 0x3a, 0x60, 0x56, 0xd7, 0xa1,
	0x9c, 0x76, 0xb6, 0x73, 0x1f, 0xc2, 0xd4, 0x7d, 0x75, 0xeb, 0x5b, 0x07,
	0x00, 0x0f, 0x80, 0xd9, 0x60, 0xb9, 0xb8, 0x3f, 0xc0, 0x9a, 0xb5, 0x3b,
	0x44, 0x71, 0xa6, 0x67, 0xc7, 0x72, 0x97, 0xaf, 0xac, 0xdf, 0x5e, 0x45,
	0xee, 0xce, 0x5a, 0xb4, 0xf7, 0x2b, 0xab, 0x91, 0x6c, 0xe8, 0x8b, 0x00,
	0x28, 0x7a, 0x91, 0x2f, 0x77, 0x99, 0x4a, 0x7f, 0x23, 0xb1, 0x37, 0x7a,
	0x06, 0x00, 0xd9, 0xbd, 0x8d, 0x53, 0xbe, 0x3f, 0x98, 0x38, 0xdb, 0xda,
	0xaf, 0x84, 0x46, 0xac, 0x5d, 0xa9, 0x7a, 0x86, 0x40, 0x4b, 0x02, 0x30,
	0x4b, 0x3c, 0x2d, 0xb3, 0xfc, 0x5b, 0x73, 0xbc, 0xa7, 0x1f, 0x26, 0x96,
	0x9e, 0x6d, 0xa2, 0xe3, 0x47, 0x61, 0xe5, 0xa7, 0x42, 0x05, 0x3c, 0x03,
	0x5a, 0x94, 0x25, 0xcf, 0x33, 0xc0, 0xfb, 0x9c, 0x59, 0x9a, 0x53, 0xac,
	0xce, 0xfd, 0xfc, 0x28, 0xea, 0x7f, 0xf4, 0x10, 0x64, 0xf3, 0x9b, 0x9b,
	0xa7, 0xb8, 0xc9, 0xcc, 0x91, 0x81, 0xb1, 0x7a, 0x68, 0x4a, 0x00, 0xc2,
	0xef, 0x86, 0x06, 0x8a, 0xa1, 0xe2, 0xfa, 0x97, 0xa8, 0x86, 0xdf, 0x7b,
	0xca, 0x80, 0xcb, 0x16, 0x36, 0x37, 0x17, 0x73, 0x67, 0x4e, 0xf2, 0x14,
	0x34, 0x56, 0x8f, 0x92, 0xf2, 0x7c, 0x40, 0x6f, 0x7d, 0x0d, 0x10, 0xf9,
	0x5d, 0x8e, 0x75, 0x27, 0x18, 0x8c, 0x56, 0x6f, 0x01, 0x98, 0xf9, 0x87,
	0xdb, 0xce, 0x1c, 0x81, 0x69, 0xbf, 0x58, 0x8b, 0x9b, 0xdb, 0x81, 0x7a,
	0x91, 0xc9, 0xa1, 0x6c, 0x51, 0x03, 0x77, 0xce, 0xb8, 0x5c, 0x7f, 0xe5,
	0x8a, 0xcf, 0xc6, 0x1e, 0x02, 0x78, 0x38, 0x9e, 0xfe, 0xde, 0x1c, 0x83,
	0x5d, 0x38, 0x55, 0x0b, 0x87, 0x5f, 0x67, 0xfb, 0x1d, 0x3e, 0x68, 0x2b,
	0x61, 0xbe, 0x84, 0x6b, 0x7f, 0xeb, 0x50, 0x6b, 0x97, 0x63, 0x1e, 0xfd,
	0x8b, 0x5b, 0x3c, 0xee, 0xcf, 0xf3, 0x7f, 0x04, 0x18, 0x00, 0xe0, 0x6e,
	0xdd, 0x63, 0x24, 0x57, 0x80, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45,
	0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}
