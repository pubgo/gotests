// This file is generated by gorazor 1.2.2
// DON'T modified manually
// Should edit source file and re-generate: examples/tpl/helper/footer.gohtml

package helper

import (
	"io"
	"strings"
)

// Footer generates examples/tpl/helper/footer.gohtml
func Footer() string {
	var _b strings.Builder
	RenderFooter(&_b)
	return _b.String()
}

// RenderFooter render examples/tpl/helper/footer.gohtml
func RenderFooter(_buffer io.StringWriter) {
	// Line: 1
	_buffer.WriteString("<div>copyright 2014</div>")

}
