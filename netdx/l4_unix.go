package netdx

import (
	"net"
	"time"
	_ "unsafe"
)

//go:linkname net.DialUnix DialUnix
func DialUnix(network string, laddr, raddr *net.UnixAddr) (conn *net.UnixConn, err error) {
	conn, err = net.DialUnix(network, laddr, raddr)

	id := generateID()

	defer func() {
		if err != nil {
			L4Errorf("[UNIX-DX] id: %d, dialUnix, laddr: %+v, raddr: %+v, err: %v \n",
				id, laddr, raddr, err)
			return
		}
		storeConn(id, conn)

		L4Printf("[UNIX-DX] id: %d, dialUnix, laddr: %+v, raddr: %+v, OK \n", id, laddr, raddr)
	}()

	timer := time.AfterFunc(time.Second, func() {
		L4Errorf("[UNIX-DX] id: %d, dialUnix, laddr: %+v, raddr: %+v, err: connection is stuck \n",
			id, laddr, raddr)
	})
	defer timer.Stop()

	L4Printf("[UNIX-DX] id: %d, dialUnix, laddr: %+v, raddr: %+v \n", id, laddr, raddr)

	return
}

func unixDx() {
	checkGCFlags()
}
