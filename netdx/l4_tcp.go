package netdx

import (
	"context"
	"net"
	"time"
	_ "unsafe"
)

//go:linkname dialTCP net.(*sysDialer).dialTCP
func dialTCP(sysDialer *sysDialer, ctx context.Context, laddr, raddr *net.TCPAddr) (conn *net.TCPConn, err error) {
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
	return doDialTCP(sysDialer, ctx, laddr, raddr)
}

func tpcDx() {
	checkGCFlags()

	// func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
	//	if testHookDialTCP != nil {
	//		return testHookDialTCP(ctx, sd.network, laddr, raddr)
	//	}
	//	return sd.doDialTCP(ctx, laddr, raddr)
	//}
}

//go:linkname doDialTCP net.(*sysDialer).doDialTCP
func doDialTCP(sysDialer *sysDialer, ctx context.Context, laddr, raddr *net.TCPAddr) (*net.TCPConn, error)
