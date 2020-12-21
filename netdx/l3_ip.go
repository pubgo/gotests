package netdx

import (
	"net"
	"time"
	_ "unsafe"
)

//go:linkname net.DialIP DialIP
func DialIP(network string, laddr, raddr *net.IPAddr) (conn *net.IPConn, err error) {
	conn, err = net.DialIP(network, laddr, raddr)

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

	return

}

func ipDx() {
	checkGCFlags()
}
