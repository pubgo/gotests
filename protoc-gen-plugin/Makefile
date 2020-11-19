
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: clear gen
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/pubgo/catdog/example/hello/proto \
   --proto_path=${GOPATH}/src:. \
   --go_out=plugins=grpc:. \
   --hello_out=. \
   proto/*

clear:
	rm -rf proto/*.go


gen:
	cd protoc-gen-hello && go install .