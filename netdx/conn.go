package netdx

import (
	"io"
	"net"
	"reflect"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

const (
	IP   = "ip"
	TCP  = "tcp"
	UDP  = "udp"
	UNIX = "unix"
)

//go:linkname CloseConn net.(*conn).Close
func CloseConn(c *conn) error {
	if !ok(unsafe.Pointer(c)) {
		return syscall.EINVAL
	}

	cInfo := getConnInfo(c)
	cInfo.logErrorf("%s id: %d, %s, laddr: %+v, raddr: %+v \n",
		cInfo.logPrefix, cInfo.id, "Close", cInfo.laddr, cInfo.raddr)

	removeConn(c)

	err := iClose(unsafe.Pointer(&(c.fd.pfd)))
	if err != nil {
		return errFn(c, "close", err)
	}
	return nil
}

//go:linkname WriteConn net.(*conn).Write
func WriteConn(c *conn, b []byte) (int, error) {
	if !ok(unsafe.Pointer(c)) {
		return 0, syscall.EINVAL
	}

	var (
		cInfo = getConnInfo(c)
		now   = time.Now()
	)

	if VerboseEnabled() {
		cInfo.logPrintf("%s id: %d, %s, laddr: %+v, raddr: %+v, cost: %d, []rune: %v, msg: %v \n",
			cInfo.logPrefix, cInfo.id, "Write", cInfo.laddr, cInfo.raddr,
			time.Since(now), []rune(string(b)), string(b))
	}

	timer := time.AfterFunc(time.Second, func() {
		cInfo.logPrintf("%s id: %d, %s, laddr: %+v, raddr: %+v, err: connection is stuck \n",
			cInfo.logPrefix, cInfo.id, "Write", cInfo.laddr, cInfo.raddr)
	})
	defer timer.Stop()

	n, err := iWrite(unsafe.Pointer(&(c.fd.pfd)), b)
	if err != nil {
		err = errFn(c, "write", err)
	}
	return n, err
}

//go:linkname ReadConn net.(*conn).Read
func ReadConn(c *conn, b []byte) (int, error) {
	if !ok(unsafe.Pointer(c)) {
		return 0, syscall.EINVAL
	}

	n, err := iRead(unsafe.Pointer(&(c.fd.pfd)), b)

	if VerboseEnabled() {
		cInfo := getConnInfo(c)
		cInfo.logPrintf("%s id: %d, %s, laddr: %+v, raddr: %+v, n: %d, []rune: %v, msg: %v \n",
			cInfo.logPrefix, cInfo.id, "Read", cInfo.laddr, cInfo.raddr, n, []rune(string(b[:n])), string(b[:n]))
	}

	if err != nil && err != io.EOF {
		err = errFn(c, "read", err)
	}
	return n, err
}

var errFn = func(c *conn, op string, err error) *net.OpError {
	cInfo := getConnInfo(c)
	cInfo.logErrorf("%s id: %d, %s, laddr: %+v, raddr: %+v\n",
		cInfo.logPrefix, cInfo.id, op, cInfo.laddr, cInfo.raddr)

	return &net.OpError{Op: op, Net: cInfo.net, Source: cInfo.laddr, Addr: cInfo.raddr, Err: err}
}

func connDx() {
	checkGCFlags()
}

type connInfo struct {
	layer string
	net   string
	id    int
	laddr net.Addr
	raddr net.Addr

	logPrefix string
	logPrintf func(format string, a ...interface{})
	logErrorf func(format string, a ...interface{})
}

var (
	allConn = sync.Map{}
)

type iConn struct {
	id int
	cc interface{}
}

func storeConn(id int, cc interface{}) {
	ptr := reflect.ValueOf(cc).Pointer()
	allConn.Store(ptr, &iConn{id, cc})
}

func getConn(cc interface{}) (c interface{}, id int, ok bool) {
	ptr := reflect.ValueOf(cc).Pointer()
	if c, ok := allConn.Load(ptr); ok && c != nil {
		return c.(*iConn).cc, c.(*iConn).id, true
	}
	return nil, 0, false
}

func removeConn(cc interface{}) {
	ptr := reflect.ValueOf(cc).Pointer()
	allConn.Delete(ptr)
}

func getConnInfo(c *conn) *connInfo {

	cc, id, _ := getConn(c)
	info := &connInfo{}

	switch cc := cc.(type) {
	case *net.TCPConn:
		info.laddr = cc.LocalAddr()
		info.raddr = cc.RemoteAddr()
		info.net = TCP
		info.logPrefix = "[TCP-DX ]"
		info.layer = L4
	case *net.UDPConn:
		info.laddr = cc.LocalAddr()
		info.raddr = cc.RemoteAddr()
		info.net = UDP
		info.logPrefix = "[UDP-DX ]"
		info.layer = L4
	case *net.UnixConn:
		info.laddr = cc.LocalAddr()
		info.raddr = cc.RemoteAddr()
		info.net = UNIX
		info.logPrefix = "[UNIX-DX]"
		info.layer = L4
	case *net.IPConn:
		info.laddr = cc.LocalAddr()
		info.raddr = cc.RemoteAddr()
		info.net = IP
		info.logPrefix = "[IP-DX  ]"
		info.layer = L3
	default:
		info.logPrintf = func(string, ...interface{}) {}
		info.logErrorf = func(string, ...interface{}) {}
	}

	switch info.layer {
	case L3:
		info.logPrintf = L3Printf
		info.logErrorf = L3Errorf
	case L4:
		info.logPrintf = L4Printf
		info.logErrorf = L4Errorf
	}

	info.id = id

	return info
}

//go:linkname sysDialer net.sysDialer
type sysDialer struct {
	net.Dialer
	network, address string // nolint: unused
}

type conn struct {
	fd *netFD
}

type netFD struct {
	pfd int // poll.FD
}

//go:linkname ok  net.(*conn).ok
func ok(_ unsafe.Pointer) bool

//go:linkname iClose  internal/poll.(*FD).Close
func iClose(_ unsafe.Pointer) error

//go:linkname iRead  internal/poll.(*FD).Read
func iRead(_ unsafe.Pointer, b []byte) (int, error)

//go:linkname iWrite  internal/poll.(*FD).Write
func iWrite(_ unsafe.Pointer, b []byte) (int, error)
