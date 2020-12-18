package netdx

import (
	"context"
	"net"
	"reflect"
	"syscall"
	"time"
	"unsafe"
)

//go:linkname dialIP1 net.(*sysDialer).dialIP
func dialIP1(sd *sysDialer, ctx context.Context, laddr, raddr *net.IPAddr) (conn *net.IPConn, err error) {
	id := generateID()

	defer func() {
		if err != nil {
			L3Errorf("[IP-DX  ] id: %d, dialIP, laddr: %+v, raddr: %+v, err: %v \n",
				id, laddr, raddr, err)
			return
		}

		storeConn(id, conn)

		L3Printf("[IP-DX  ] id: %d, dialIP, laddr: %+v, raddr: %+v, OK \n", id, laddr, raddr)
	}()

	L3Printf("[IP-DX  ] id: %d, dialIP, laddr: %+v, raddr: %+v \n", id, laddr, raddr)

	timer := time.AfterFunc(time.Second, func() {
		L3Errorf("[IP-DX  ] id: %d, dialIP, laddr: %+v, raddr: %+v, err: connection is stuck \n",
			id, laddr, raddr)
	})
	defer timer.Stop()

	network, proto, err := parseNetwork(ctx, sd.network, true)
	if err != nil {
		return nil, err
	}

	switch network {
	case "ip", "ip4", "ip6":
	default:
		return nil, net.UnknownNetworkError(sd.network)
	}

	// WTF
	_type := symbolTable["go.itab.*net.IPAddr,net.sockaddr"]

	fd, err := internetSocket(ctx, sd.network,
		_type, reflect.ValueOf(laddr).Pointer(),
		_type, reflect.ValueOf(raddr).Pointer(),
		syscall.SOCK_RAW, proto, "dial", sd.Dialer.Control)

	if err != nil {
		return nil, err
	}

	return newIPConn(fd), nil
}

func ipDx() {
	checkGCFlags()
}

//go:linkname newIPConn net.newIPConn
func newIPConn(fd unsafe.Pointer) *net.IPConn

//go:linkname parseNetwork net.parseNetwork
func parseNetwork(ctx context.Context, network string, needsProto bool) (afnet string, proto int, err error)
