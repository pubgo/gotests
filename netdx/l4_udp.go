package netdx

import (
	"net"
	"time"
	_ "unsafe"
)

//go:linkname net.DialUDP DialUDP
func DialUDP(network string, laddr, raddr *net.UDPAddr) (conn *net.UDPConn, err error) {
	id := generateID()

	defer func() {
		if err != nil {
			L4Errorf("[UDP-DX ] id: %d, dialUDP, laddr: %+v, raddr: %+v, err: %v \n",
				id, laddr, raddr, err)
			return
		}

		storeConn(id, conn)

		L4Printf("[UDP-DX ] id: %d, dialUDP, laddr: %+v, raddr: %+v, OK \n", id, laddr, raddr)
	}()

	L4Printf("[UDP-DX ] id: %d, dialUDP, laddr: %+v, raddr: %+v \n", id, laddr, raddr)

	timer := time.AfterFunc(time.Second, func() {
		L4Errorf("[UDP-DX ] id: %d, dialUDP, laddr: %+v, raddr: %+v, err: connection is stuck \n",
			id, laddr, raddr)
	})
	defer timer.Stop()

	return
}

func udpDx() {
	checkGCFlags()
}
