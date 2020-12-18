package netdx

import (
	"context"
	"net"
	"reflect"
	"syscall"
	"time"
	"unsafe"
	_ "unsafe"
)

//go:linkname dialUDP net.(*sysDialer).dialUDP
func dialUDP(sd *sysDialer, ctx context.Context, laddr, raddr *net.UDPAddr) (conn *net.UDPConn, err error) {
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

	// WTF
	_type := symbolTable["go.itab.*net.UDPAddr,net.sockaddr"]

	fd, err := internetSocket(ctx, sd.network,
		_type, reflect.ValueOf(laddr).Pointer(),
		_type, reflect.ValueOf(raddr).Pointer(),
		syscall.SOCK_DGRAM, 0, "dial", sd.Dialer.Control)

	if err != nil {
		return nil, err
	}

	return newUDPConn(fd), nil
}
func udpDx() {
	checkGCFlags()

	// func (sd *sysDialer) dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error) {
	//	fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_DGRAM, 0, "dial", sd.Dialer.Control)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return newUDPConn(fd), nil
	//}
}

//go:linkname internetSocket net.internetSocket
func internetSocket(ctx context.Context, net string,
	_ uintptr, _ uintptr,
	_ uintptr, _ uintptr,
	sotype, proto int, mode string, ctrlFn func(string, string, syscall.RawConn) error) (fd unsafe.Pointer, err error)

//go:linkname newUDPConn net.newUDPConn
func newUDPConn(fd unsafe.Pointer) *net.UDPConn
