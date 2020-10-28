
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/pubgo/catdog/example/hello/proto \
   --proto_path=${GOPATH}/src:. \
   --go-micro_out=. \
   --go_out=. \
   --grpc-gateway_out=. \
   --nrpc_out=. \
   proto/*

clear:
	rm -rf proto/*.go