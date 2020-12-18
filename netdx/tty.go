package netdx

import (
	"context"
	"errors"
	"github.com/gowins/gotty/backend/localcommand"
	"github.com/gowins/gotty/server"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"time"
)

// You can see the network diagnostic information on your web browser
func ShowTTY() error {
	checkGCFlags()

	// Only supports macos
	if runtime.GOOS != "darwin" {
		return errors.New("only supports macos")
	}

	// CLI: tail -100f /tmp/*_dx.log
	backendOptions := &localcommand.Options{}
	factory, err := localcommand.NewFactory("tail", []string{"-100f", writer.Name()}, backendOptions)
	if err != nil {
		return err
	}

	// Create a new instance of Server
	appOptions := &server.Options{
		Address: "0.0.0.0",
		Port:    "6996",
	}
	srv, err := server.New(factory, appOptions)
	if err != nil {
		return err
	}

	go func() {
		// Open the URL(http://127.0.0.1:6996) on your web browser
		time.AfterFunc(time.Second/2, func() {
			cmd := exec.Command("open", "-a", "Safari", "http://127.0.0.1:6996")
			cmd.Stdout = ioutil.Discard
			_ = cmd.Run()
		})

		// Run the http server
		err := srv.Run(context.Background())
		log.Printf("[netdx] ShowTTY, err: %v \n", err)
	}()

	// Waiting for the browser to open
	time.Sleep(time.Second * 3)

	return nil
}
