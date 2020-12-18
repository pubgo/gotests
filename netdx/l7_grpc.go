package netdx

import (
	"context"
	"sync"
	"time"
	_ "unsafe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

//go:linkname channelIsOn google.golang.org/grpc/internal/channelz.IsOn
func channelIsOn() bool {
	return true
}

//go:linkname RegisterChannel google.golang.org/grpc/internal/channelz.RegisterChannel
func RegisterChannel(_ uintptr, cc **grpc.ClientConn, _ int64, _ string) int64 {
	c := &gConn{generateID(), *cc}
	id := db.Put(c)

	go func() {
		ticker := time.NewTicker(time.Second / 2)
		defer ticker.Stop()

		var i int
		for range ticker.C {
			state := c.GetState()

			if state == connectivity.Shutdown {
				break
			}
			if state == connectivity.Ready {
				continue
			}

			L7Errorf("[GRPC-DX] id: %d, the connectivity.State of ClientConn is %s, target: %s\n",
				c.rid, state, c.Target())

			// Three strikes law
			// https://en.wikipedia.org/wiki/Three-strikes_law
			if i++; i >= 3 {
				break
			}
		}
	}()

	L7Printf("[GRPC-DX] id: %d, create a ClientConn, target: %s \n", c.rid, c.Target())

	return id
}

//go:linkname RegisterSubChannel google.golang.org/grpc/internal/channelz.RegisterSubChannel
func RegisterSubChannel() int64 {
	return 0
}

//go:linkname RegisterNormalSocket google.golang.org/grpc/internal/channelz.RegisterNormalSocket
func RegisterNormalSocket() int64 {
	return 0
}

//go:linkname RegisterListenSocket google.golang.org/grpc/internal/channelz.RegisterListenSocket
func RegisterListenSocket() int64 {
	return 0
}

//go:linkname RegisterServer google.golang.org/grpc/internal/channelz.RegisterServer
func RegisterServer() int64 {
	return 0
}

//go:linkname AddTraceEvent google.golang.org/grpc/internal/channelz.AddTraceEvent
func AddTraceEvent() {
	return
}

//go:linkname RemoveEntry google.golang.org/grpc/internal/channelz.RemoveEntry
func RemoveEntry(id int64) {
	if c := db.Get(id); c != nil {
		L7Errorf("[GRPC-DX] id: %d, tears down the ClientConn, target: %s \n", c.rid, c.Target())
	}
}

//go:linkname GrpcInvoke google.golang.org/grpc.invoke
func GrpcInvoke(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) (err error) {
	c := db.GetByConn(cc)
	if c == nil {
		return
	}

	now := time.Now()

	defer func() {
		if err != nil {
			L7Errorf("[GRPC-DX] id: %d, invoke, method: %s, target: %s, err: %v \n",
				c.rid, method, cc.Target(), err)
			return
		}

		if VerboseEnabled() {
			L7Printf("[GRPC-DX] id: %d, invoke, method: %s, target: %s, reply: %+v, cost: %v\n",
				c.rid, method, cc.Target(), reply, time.Since(now))
		}
	}()

	timer := time.AfterFunc(time.Second, func() {
		L7Errorf(
			"[GRPC-DX] id: %d, invoke, method: %s, target: %s, err: connection is stuck \n",
			c.rid, method, cc.Target())
	})
	defer timer.Stop()

	if VerboseEnabled() {
		L7Printf("[GRPC-DX] id: %d, invoke, method: %s, target: %s, req: %+v\n",
			c.rid, method, cc.Target(), req)
	}

	cs, err := newClientStream(ctx, unaryStreamDesc, cc, method, opts...)
	if err != nil {
		return err
	}

	if VerboseEnabled() {
		L7Printf("[GRPC-DX] id: %d, invoke, method: %s, target: %s, cost: %v \n",
			c.rid, method, cc.Target(), time.Since(now))

		now = time.Now()
	}

	if err := cs.SendMsg(req); err != nil {
		return err
	}

	return cs.RecvMsg(reply)
}

func grpcDx() {
	checkGCFlags()

	// AS NEEDED
	if _, ok := symbolTable["google.golang.org/grpc/internal/channelz.IsOn"]; !ok {
		return
	}
	if _, ok := symbolTable["google.golang.org/grpc/internal/channelz.RegisterChannel"]; !ok {
		return
	}

	//func invoke(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error {
	//	cs, err := newClientStream(ctx, unaryStreamDesc, cc, method, opts...)
	//	if err != nil {
	//		return err
	//	}
	//	if err := cs.SendMsg(req); err != nil {
	//		return err
	//	}
	//	return cs.RecvMsg(reply)
	//}
}

//go:linkname unaryStreamDesc google.golang.org/grpc.unaryStreamDesc
var unaryStreamDesc *grpc.StreamDesc

//go:linkname newClientStream google.golang.org/grpc.newClientStream
func newClientStream(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (_ grpc.ClientStream, err error)

var db = newDB()

type gDB struct {
	mu    sync.Mutex
	id    int64
	store map[int64]*gConn
}

type gConn struct {
	rid int
	*grpc.ClientConn
}

func newDB() *gDB {
	return &gDB{
		store: make(map[int64]*gConn),
	}
}

func (g *gDB) Get(id int64) *gConn {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.store[id]
}

func (g *gDB) GetByConn(cc *grpc.ClientConn) *gConn {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, c := range g.store {
		if c.ClientConn == cc {
			return c
		}
	}
	return nil
}

func (g *gDB) Put(c *gConn) (id int64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	// In the worst case, it would get stuck in an endless loop.
	// But, that's not going to happen.
	for {
		g.id++
		if _, ok := g.store[g.id]; !ok && g.id != 0 {
			break
		}
	}
	g.store[g.id] = c
	return g.id
}

func (g *gDB) Remove(id int64) {
	g.mu.Lock()
	delete(g.store, id)
	g.mu.Unlock()
}
