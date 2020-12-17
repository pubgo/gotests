package main

import (
	"time"

	"github.com/go-echarts/statsview"
)

func main() {
	go func() {
		mgr := statsview.New()

		// Start() runs a HTTP server at `localhost:18066` by default.
		mgr.Start()

		// Stop() will shutdown the http server gracefully
		// mgr.Stop()
	}()

	// busy working....
	time.Sleep(time.Minute)
}
