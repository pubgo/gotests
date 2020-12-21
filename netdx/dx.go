package netdx

import (
	"math/rand"
	"time"
	_ "unsafe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/**
 * [√] OSI Layer 7 - application layer
 * [x] OSI Layer 6 - presentation layer
 * [x] OSI Layer 5 - session layer
 * [√] OSI Layer 4 - transport layer
 * [√] OSI Layer 3 - network layer
 * [x] OSI Layer 2 - data link layer
 * [x] OSI Layer 1 - physical layer
 *
 * NOTE: go run -gcflags="all=-N -l" main.go
 */
func AllLayerDx(verbose bool) {
	if verbose {
		turnOnVerbose()
	}

	connDx()
	L3Dx(false)
	L4Dx(false)
	L7Dx(false)
}

// OSI Layer 3 - network layer
// [√] IP
func L3Dx(verbose bool) {
	if verbose {
		turnOnVerbose()
	}

	connDx()
	turnOnL3()
	ipDx()
}

// OSI Layer 4 - transport layer
// [√] TCP
// [√] UDP
// [√] UNIX
func L4Dx(verbose bool) {
	if verbose {
		turnOnVerbose()
	}

	connDx()
	turnOnL4()
	tpcDx()
	udpDx()
	unixDx()
}

// OSI Layer 7 - application layer
// [√] HTTP/HTTPS
// [√] GRPC
func L7Dx(verbose bool) {
	if verbose {
		turnOnVerbose()
	}

	turnOnL7()
	httpDx()
	grpcDx()
}

// Just for testing 😁
func x666() {}

func checkGCFlags() {

	help := func() {
		errorf("[ERROR] Please add \"all=-N -l\" to your gcflags \n")
		errorf("E.x. go run -gcflags=\"all=-N -l\" main.go \n")
		//os.Exit(1)
	}

	help()

	x666()
}

// Random ID
func generateID() int {
	n := rand.Intn(999999)
	return n%899999 + 100000
}
