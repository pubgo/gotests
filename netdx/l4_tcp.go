package netdx

import (
	"net"
	"time"
	_ "unsafe"
)

//go:linkname net.Listen Listen
func Listen(network, address string) (lt net.Listener, err error) {
	lt, err = net.Listen(network, address)

	id := generateID()

	L4Printf("[TCP-DX ] id: %d, dialTCP, laddr:, raddr: \n", id)
	return
}

//go:linkname net.DialTCP DialTCP
func DialTCP(network string, laddr, raddr *net.TCPAddr) (conn *net.TCPConn, err error) {
	conn, err = net.DialTCP(network, laddr, raddr)

	id := generateID()

	defer func() {
		if err != nil {
			L4Errorf("[TCP-DX ] id: %d, dialTCP, laddr: %+v, raddr: %+v, err: %v \n",
				id, laddr, raddr, err)
			return
		}

		storeConn(id, conn)

		L4Printf("[TCP-DX ] id: %d, dialTCP, laddr: %+v, raddr: %+v, OK \n", id, laddr, raddr)
	}()

	timer := time.AfterFunc(time.Second, func() {
		L4Errorf("[TCP-DX ] id: %d, dialTCP, laddr: %+v, raddr: %+v, err: connection is stuck \n",
			id, laddr, raddr)
	})
	defer timer.Stop()

	L4Printf("[TCP-DX ] id: %d, dialTCP, laddr: %+v, raddr: %+v \n", id, laddr, raddr)
	return
}

func tpcDx() {
	checkGCFlags()
}
