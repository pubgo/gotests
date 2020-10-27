package grpcserver

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	uIntOpt := grpc.UnaryInterceptor(xrequestid.UnaryServerInterceptor())
	sIntOpt := grpc.StreamInterceptor(xrequestid.StreamServerInterceptor())
	grpc.NewServer(uIntOpt, sIntOpt)
}

// https://github.com/mercari/go-grpc-interceptor

func init() {
	srv := grpc.NewServer()
	l, err := net.Listen("tcp", s.Host)
	if err != nil {
		panic(err)
	}
	go func() {
		_ = srv.Serve(l)
	}()

	srv.GracefulStop()
}

func SignalContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		signal.Stop(signals)
		close(signals)
		cancel()
	}()

	return ctx
}

func init() {
	gopts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.UnknownServiceHandler(g.handler),
	}

	if creds := g.getCredentials(); creds != nil {
		gopts = append(gopts, grpc.Creds(creds))
	}

	if opts := g.getGrpcOptions(); opts != nil {
		gopts = append(gopts, opts...)
	}

}

// https://github.com/grpc-ecosystem/go-grpc-middleware
func init() {
	myServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_auth.StreamServerInterceptor(myAuthFunction),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_auth.UnaryServerInterceptor(myAuthFunction),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	pb.RegisterSmsServer(s, &server{})
	reflection.Register(s)

	if s.backendIsInsecure {
		opt = append(opt, grpc.WithInsecure())
	} else {
		opt = append(opt, grpc.WithTransportCredentials(credentials.NewTLS(s.backendTLS)))
	}

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		return metadata.NewOutgoingContext(ctx, md.Copy()), backendConn, nil
	}
	grpcServer := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()), // needed for proxy to function.
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
		/*grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
		),*/ //middleware should be a config setting or 3rd party middleware plugins like for caddyhttp
	)

	// gRPC-Web compatibility layer with CORS configured to accept on every
	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithCorsForRegisteredEndpointsOnly(false))
	wrappedGrpc.ServeHTTP(w, r)

}

// http://doc.oschina.net/grpc
// https://learnku.com/docs/the-way-to-go
// https://grpc.io/docs/languages/go/quickstart/
// https://github.com/improbable-eng/grpc-web/tree/master/go/grpcweb
