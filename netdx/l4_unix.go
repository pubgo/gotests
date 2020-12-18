package netdx

import (
	"context"
	"net"
	"reflect"
	"syscall"
	"time"
	"unsafe"
)

//go:linkname dialUnix net.(*sysDialer).dialUnix
func dialUnix(sd *sysDialer, ctx context.Context, laddr, raddr *net.UnixAddr) (conn *net.UnixConn, err error) {
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

	// WTF
	_type := symbolTable["go.itab.*net.UDPAddr,net.sockaddr"]

	fd, err := unixSocket(ctx, sd.network,
		_type, reflect.ValueOf(laddr).Pointer(),
		_type, reflect.ValueOf(raddr).Pointer(), "dial", sd.Dialer.Control)

	if err != nil {
		return nil, err
	}
	return newUnixConn(fd), nil
}

func unixDx() {
	checkGCFlags()

	// func (sd *sysDialer) dialUnix(ctx context.Context, laddr, raddr *UnixAddr) (*UnixConn, error) {
	//	fd, err := unixSocket(ctx, sd.network, laddr, raddr, "dial", sd.Dialer.Control)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return newUnixConn(fd), nil
	//}
}

//go:linkname unixSocket net.unixSocket
func unixSocket(ctx context.Context, net string,
	_ uintptr, _ uintptr,
	_ uintptr, _ uintptr,
	mode string, ctrlFn func(string, string, syscall.RawConn) error) (fd unsafe.Pointer, err error)

//go:linkname newUnixConn net.newUnixConn
func newUnixConn(fd unsafe.Pointer) *net.UnixConn
