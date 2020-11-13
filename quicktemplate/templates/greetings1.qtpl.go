// Code generated by qtc from "greetings1.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Greetings greets up to 42 names.
// It also greets John differently comparing to others.

//line templates/greetings1.qtpl:3
package templates

//line templates/greetings1.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/greetings1.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/greetings1.qtpl:3
func StreamGreetings(qw422016 *qt422016.Writer, names []string) {
//line templates/greetings1.qtpl:3
	qw422016.N().S(`
    `)
//line templates/greetings1.qtpl:4
	if len(names) == 0 {
//line templates/greetings1.qtpl:4
		qw422016.N().S(`
        Nobody to greet :(
        `)
//line templates/greetings1.qtpl:6
		return
//line templates/greetings1.qtpl:7
	}
//line templates/greetings1.qtpl:7
	qw422016.N().S(`

    `)
//line templates/greetings1.qtpl:9
	for i, name := range names {
//line templates/greetings1.qtpl:9
		qw422016.N().S(`
        `)
//line templates/greetings1.qtpl:10
		if i == 42 {
//line templates/greetings1.qtpl:10
			qw422016.N().S(`
            I'm tired to greet so many people...
            `)
//line templates/greetings1.qtpl:12
			break
//line templates/greetings1.qtpl:13
		} else if name == "John" {
//line templates/greetings1.qtpl:13
			qw422016.N().S(`
            `)
//line templates/greetings1.qtpl:14
			streamsayHi(qw422016, "Mr. "+name)
//line templates/greetings1.qtpl:14
			qw422016.N().S(`
            `)
//line templates/greetings1.qtpl:15
			continue
//line templates/greetings1.qtpl:16
		} else {
//line templates/greetings1.qtpl:16
			qw422016.N().S(`
            `)
//line templates/greetings1.qtpl:17
			StreamHello(qw422016, name)
//line templates/greetings1.qtpl:17
			qw422016.N().S(`
        `)
//line templates/greetings1.qtpl:18
		}
//line templates/greetings1.qtpl:18
		qw422016.N().S(`
    `)
//line templates/greetings1.qtpl:19
	}
//line templates/greetings1.qtpl:19
	qw422016.N().S(`
`)
//line templates/greetings1.qtpl:20
}

//line templates/greetings1.qtpl:20
func WriteGreetings(qq422016 qtio422016.Writer, names []string) {
//line templates/greetings1.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/greetings1.qtpl:20
	StreamGreetings(qw422016, names)
//line templates/greetings1.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line templates/greetings1.qtpl:20
}

//line templates/greetings1.qtpl:20
func Greetings(names []string) string {
//line templates/greetings1.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/greetings1.qtpl:20
	WriteGreetings(qb422016, names)
//line templates/greetings1.qtpl:20
	qs422016 := string(qb422016.B)
//line templates/greetings1.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/greetings1.qtpl:20
	return qs422016
//line templates/greetings1.qtpl:20
}