// This file is generated by gorazor 1.2.2
// DON'T modified manually
// Should edit source file and re-generate: tpl/helper/header.gohtml

package helper

import (
	"io"
	"strings"
)

// Header generates tpl/helper/header.gohtml
func Header() string {
	var _b strings.Builder
	RenderHeader(&_b)
	return _b.String()
}

// RenderHeader render tpl/helper/header.gohtml
func RenderHeader(_buffer io.StringWriter) {
	// Line: 1
	_buffer.WriteString("<div>Welcome</div>")

}
