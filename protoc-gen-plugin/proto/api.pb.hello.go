package helloworld

import rpc1 "net/rpc"

// test
/* ssss */
type TestApi interface {
	Version(in *TestReq, out *TestApiOutput) error
	VersionTest(in *TestReq, out *TestApiOutput) error
}

func RegisterTestApi(srv *rpc1.Server, x TestApi) error {
	return nil
}

// test
/* ssss */
type TestApiV2 interface {
	Version(in *TestReq, out *TestApiOutput) error
	VersionTest(in *TestReq, out *TestApiOutput) error
}

func RegisterTestApiV2(srv *rpc1.Server, x TestApiV2) error {
	return nil
}
