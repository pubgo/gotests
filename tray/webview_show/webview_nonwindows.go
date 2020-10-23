//+build !windows

package webview_show

/*
#cgo linux pkg-config: webkit2gtk-4.0
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa -framework Webkit

#include "webview.h"
*/
import "C"

func RunWebview(title string, width, height int) {
	C.configureAppWindow(C.CString(title), C.int(width), C.int(height))
}

func ShowWebview(url string) {
	C.showAppWindow(C.CString(url))
}
