package main

import (
	"bytes"
	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

func temp(svc *ServiceSpec) string {
	f := jen.NewFile("a")
	f.ImportName("b.c/d", "d")

	f.Func().Id("main").Params().Block(
		jen.Qual("b.c/d", "E").Call(),
	)

	return f.GoString()
}

const tmplService = `
{{$root := .}}
type {{.ServiceName}}Interface interface {
	{{- range $_, $m := .MethodList}}
		{{$m.MethodName}}(in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}}(srv *rpc.Server, x {{.ServiceName}}Interface) error {
	if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
		return err
	}
	return nil
}

type {{.ServiceName}}Client struct {
	*rpc.Client
}
var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)
func Dial{{.ServiceName}}(network, address string) (*{{.ServiceName}}Client, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{Client: c}, nil
}
{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}}) error {
	return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`

func main() {
	var request plugin.CodeGeneratorRequest
	var response plugin.CodeGeneratorResponse

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: reading input: %v", err)
	}
	if err := proto.Unmarshal(data, &request); err != nil {
		log.Fatalf("error: parsing input proto: %v", err)
	}

	if len(request.GetFileToGenerate()) == 0 {
		log.Fatal("error: no files to generate")
	}

	for _, name := range request.GetFileToGenerate() {
		var fd *descriptor.FileDescriptorProto
		for _, fd = range request.GetProtoFile() {
			if name == fd.GetName() {
				break
			}
		}

		if fd == nil {
			log.Fatalf("could not find the .proto file for %s", name)
		}

		spec := &ServiceSpec{
			ServiceName: generator.CamelCase(fd.GetName()),
		}

		for _, m := range fd.GetService()[0].GetMethod() {
			spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
				MethodName:     generator.CamelCase(m.GetName()),
				InputTypeName:  m.GetInputType(),
				OutputTypeName: m.GetOutputType(),
			})
		}

		var buf bytes.Buffer
		t := template.Must(template.New("").Parse(tmplService))
		err := t.Execute(&buf, spec)
		if err != nil {
			log.Fatal(err)
		}

		response.File = append(response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String("proto/api.pb.hello.go"),
			Content: proto.String(buf.String()),
		})
	}

	data, err = proto.Marshal(&response)
	if err != nil {
		log.Println("failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		log.Println("failed to write output proto")
	}
}
