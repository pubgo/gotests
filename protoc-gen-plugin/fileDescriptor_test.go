package protoc_gen_plugin

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	descpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pubgo/xerror"
	"io/ioutil"
	"testing"
)

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func TestName(t *testing.T) {
	var fileDescriptor_ecf0878b123623e2 = []byte{
		// 401 bytes of a gzipped FileDescriptorProto
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
		0x14, 0x66, 0x6e, 0x7b, 0x6f, 0xec, 0x09, 0x22, 0xcc, 0x5d, 0xdc, 0xb4, 0x14, 0x53, 0x82, 0x42,
		0x09, 0x98, 0x60, 0xc4, 0x4d, 0x77, 0x2d, 0xee, 0x85, 0xb1, 0x14, 0x5c, 0xc9, 0xb4, 0x1d, 0xd3,
		0x81, 0x24, 0x33, 0x4d, 0x26, 0xad, 0x6b, 0x5f, 0xc1, 0x77, 0xf0, 0x55, 0xc4, 0xb5, 0xfb, 0xac,
		0x5c, 0xe5, 0x29, 0x24, 0x33, 0x53, 0x5a, 0xa4, 0x2e, 0x04, 0xdd, 0x9c, 0x9c, 0x9f, 0xef, 0xfb,
		0x4e, 0xbe, 0xc3, 0xc0, 0x13, 0x59, 0x0a, 0x25, 0x62, 0x2a, 0x79, 0xa4, 0x33, 0x0c, 0x3b, 0x96,
		0x65, 0xe2, 0x28, 0xca, 0x6c, 0x3b, 0x7a, 0x91, 0x72, 0xb5, 0xab, 0xd7, 0xd1, 0x46, 0xe4, 0x71,
		0x2a, 0x52, 0x11, 0x6b, 0xc8, 0xba, 0xfe, 0xa8, 0x2b, 0xc3, 0xec, 0x32, 0x43, 0x1d, 0x8d, 0x53,
		0x21, 0xd2, 0x8c, 0x75, 0x62, 0x31, 0x2d, 0x0a, 0xa1, 0xa8, 0xe2, 0xa2, 0xa8, 0xcc, 0x34, 0x08,
		0xc1, 0x59, 0xb2, 0x4a, 0x11, 0xb6, 0xc7, 0x3e, 0xdc, 0xf2, 0x42, 0xd6, 0xca, 0x43, 0x13, 0x34,
		0x1d, 0x2c, 0x06, 0x6d, 0xe3, 0x9b, 0x06, 0x31, 0x9f, 0xe0, 0x03, 0xb8, 0x1d, 0x76, 0x2e, 0xf9,
		0x1b, 0xaa, 0x28, 0x7e, 0x06, 0xce, 0x81, 0x95, 0x15, 0x17, 0x85, 0x65, 0x40, 0xdb, 0xf8, 0x77,
		0x73, 0xc9, 0x57, 0xac, 0x24, 0xa7, 0x11, 0x0e, 0x01, 0xaa, 0xf2, 0xb0, 0xb2, 0xc0, 0x9b, 0x33,
		0xf0, 0x9d, 0xee, 0x92, 0x8b, 0x69, 0xf0, 0x15, 0xc1, 0x63, 0xbb, 0xe1, 0x6d, 0xad, 0x64, 0xad,
		0xf0, 0x18, 0xfa, 0x1b, 0xb1, 0x65, 0x7a, 0xc1, 0xed, 0xe2, 0x51, 0xdb, 0xf8, 0xba, 0x26, 0x3a,
		0xe2, 0x21, 0xf4, 0xf2, 0x2a, 0xb5, 0xa2, 0x4e, 0xdb, 0xf8, 0x5d, 0x49, 0xba, 0x80, 0x9f, 0x83,
		0x53, 0x88, 0xe3, 0x92, 0xe7, 0xcc, 0xeb, 0x4d, 0xd0, 0xb4, 0xb7, 0x70, 0xdb, 0xc6, 0x3f, 0xb5,
		0xc8, 0x29, 0xc1, 0xaf, 0xa1, 0xbf, 0xa5, 0x8a, 0x7a, 0xfd, 0x09, 0x9a, 0xba, 0xc9, 0x43, 0x74,
		0x3e, 0x73, 0x74, 0x61, 0xd5, 0x2c, 0xee, 0x80, 0x44, 0xc7, 0xe4, 0x1b, 0x32, 0x67, 0x9b, 0x4b,
		0x8e, 0xdf, 0x83, 0x63, 0xff, 0x1f, 0xdf, 0xff, 0xce, 0x27, 0x6c, 0x3f, 0x1a, 0x5e, 0x11, 0x35,
		0xee, 0x82, 0xa7, 0x9f, 0x7f, 0xfc, 0xfc, 0x72, 0xe3, 0x05, 0xf7, 0xf1, 0xe1, 0x65, 0xcc, 0x3e,
		0xd1, 0x5c, 0x66, 0x2c, 0xb6, 0x87, 0x9b, 0xa1, 0x10, 0x53, 0x70, 0xad, 0x74, 0xc7, 0xfb, 0x6b,
		0xf9, 0x40, 0xcb, 0x8f, 0x83, 0x87, 0x2b, 0xf2, 0x8a, 0x55, 0x6a, 0x86, 0xc2, 0xe4, 0x3b, 0x82,
		0x81, 0x65, 0xad, 0x92, 0x7f, 0xe4, 0x25, 0xf9, 0xbf, 0x5e, 0x92, 0x3f, 0x78, 0x59, 0xdf, 0xe9,
		0x27, 0xfd, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x90, 0xb0, 0xc6, 0x3e, 0x03, 0x00,
		0x00,
	}

	dt, err := GzipDecode(fileDescriptor_ecf0878b123623e2)
	xerror.Exit(err)
	//fmt.Println(string(dt))

	var fs descpb.FileDescriptorProto
	xerror.Exit(proto.Unmarshal(dt, &fs))
	fmt.Println(fs.GetName())
	fmt.Println(fs.GetOptions())
	fmt.Println(fs.GetService())
	fmt.Println(fs.String())
	dt, err = json.MarshalIndent(fs, "", "  ")
	xerror.Exit(err)
	fmt.Println(string(dt))
}
