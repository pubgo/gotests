

type Proto/api.protoInterface interface {
		Version(in *.helloworld.TestReq, out *.helloworld.TestApiOutput) error
		VersionTest(in *.helloworld.TestReq, out *.helloworld.TestApiOutput) error
}

func RegisterProto/api.proto(srv *rpc.Server, x Proto/api.protoInterface) error {
	if err := srv.RegisterName("Proto/api.proto", x); err != nil {
		return err
	}
	return nil
}

type Proto/api.protoClient struct {
	*rpc.Client
}
var _ Proto/api.protoInterface = (*Proto/api.protoClient)(nil)
func DialProto/api.proto(network, address string) (*Proto/api.protoClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Proto/api.protoClient{Client: c}, nil
}

func (p *Proto/api.protoClient) Version(in *.helloworld.TestReq, out *.helloworld.TestApiOutput) error {
	return p.Client.Call("Proto/api.proto.Version", in, out)
}

func (p *Proto/api.protoClient) VersionTest(in *.helloworld.TestReq, out *.helloworld.TestApiOutput) error {
	return p.Client.Call("Proto/api.proto.VersionTest", in, out)
}

