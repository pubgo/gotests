package netdx

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

var (
	verbose   bool
	writer, _ = ioutil.TempFile("/tmp", "*_dx.log")
)

const (
	red   = 31
	green = 32
)

func turnOnVerbose()       { verbose = true }
func turnOffVerbose()      { verbose = false }
func VerboseEnabled() bool { return verbose }

func printf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	stdWrite(fmt.Sprintf("\033[0;%dm √ %s\033[0m", green, s))
}

func errorf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	stdWrite(fmt.Sprintf("\033[0;%dm × %s\033[0m", red, s))
}

func stdWrite(s string) {
	_, _ = writer.Write([]byte(s))
	_, _ = syscall.Write(1, []byte(s))
}

func L3Printf(format string, a ...interface{}) {
	layerPrintf(L3, false, format, a...)
}

func L3Errorf(format string, a ...interface{}) {
	layerPrintf(L3, true, format, a...)
}

func L4Printf(format string, a ...interface{}) {
	layerPrintf(L4, false, format, a...)
}

func L4Errorf(format string, a ...interface{}) {
	layerPrintf(L4, true, format, a...)
}

func L7Printf(format string, a ...interface{}) {
	layerPrintf(L7, false, format, a...)
}

func L7Errorf(format string, a ...interface{}) {
	layerPrintf(L7, true, format, a...)
}

func layerPrintf(layer string, isErr bool, format string, a ...interface{}) {
	var enabled bool

	switch layer {
	case L3:
		enabled = L3Enabled()
	case L4:
		enabled = L4Enabled()
	case L7:
		enabled = L7Enabled()
	}

	if !enabled {
		return
	}

	if isErr {
		errorf(format, a...)
	} else {
		printf(format, a...)
	}
}
