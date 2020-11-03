package main

import (
	"github.com/valyala/fasttemplate"
	"io"
	"net/url"
)

func Template(template string, m map[string]interface{}) string {
	return fasttemplate.ExecuteString(template, "${", "}", m)
}

func main() {
	Template(
		"http://{{host}}/?foo={{bar}}{{bar}}&q={{query}}&baz={{baz}}",
		map[string]interface{}{
			"host": "google.com",     // string - convenient
			"bar":  []byte("foobar"), // byte slice - the fastest

			// TagFunc - flexible value. TagFunc is called only if the given
			// tag exists in the template.
			"query": fasttemplate.TagFunc(func(w io.Writer, tag string) (int, error) {
				return w.Write([]byte(url.QueryEscape(tag + "=world")))
			}),
		},
	)
}
