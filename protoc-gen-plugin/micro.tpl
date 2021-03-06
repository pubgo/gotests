{% package main %}

{% import gen "github.com/pubgo/xprotogen/gen" %}
{% import "strings" %}

{% func Header(pkg string, fd *gen.FileDescriptor) %}
    // Code generated by protoc-gen-micro. DO NOT EDIT.
    {%- if fd.GetOptions().GetDeprecated() -%}
        // {%s fd.GetName() %} is a deprecated file.
    {%- else -%}
        // source: {%s fd.GetName() %}
    {%- endif -%}

    package {%s pkg %}

    import (
    	fmt "fmt"
    	math "math"
    	context "context"

        client "github.com/asim/go-micro/v3/client"
    	server "github.com/asim/go-micro/v3/server"
    	_ "github.com/gogo/protobuf/gogoproto"
    	proto "github.com/golang/protobuf/proto"
    	_ "google.golang.org/genproto/googleapis/api/annotations"
    )

    // Reference imports to suppress errors if they are not otherwise used.
    var _ = proto.Marshal
    var _ = fmt.Errorf
    var _ = math.Inf

    // This is a compile-time assertion to ensure that this generated file
    // is compatible with the proto package it is being compiled against.
    // A compilation error at this line likely means your copy of the
    // proto package needs to be updated.
    const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

    // Reference imports to suppress errors if they are not otherwise used.
    var _ context.Context
    var _ client.Option
    var _ server.Option
{% endfunc %}

{% func serviceClient(srv string,m *gen.Method) %}
    func (c *{%-s gen.UnExport(srv) -%}Service){%-s clientMethod(srv,m) -%}{
        {% if !m.GetServerStreaming() && !m.GetClientStreaming() %}
            req := c.c.NewRequest(c.name, "{%s srv %}.{%s m.GetName() %}", in)
            out := new({%s m.GetOutputType() %})
            err := c.c.Call(ctx, req, out, opts...)
            if err != nil {
                return nil, err
            }
            return out, nil
        {% else %}
            req := c.c.NewRequest(c.name, "{%s srv %}.{%s m.GetName() %}", &{%s m.GetInputType() %}{})
            stream, err := c.c.Stream(ctx, req, opts...)
            if err != nil { return nil, err }
            {% if !m.GetClientStreaming() %}
                if err := stream.Send(in); err != nil { return nil, err }
            {% endif %}

            return &{%s gen.UnExport(srv)+m.GetName() %}{stream}, nil
        {% endif %}
    }


    // Stream auxiliary types and methods.
    type {%s srv %}_{%s m.GetName() %}Service interface {
        Context() context.Context
        SendMsg(interface{}) error
        RecvMsg(interface{}) error
        Close() error
        {% if m.GetClientStreaming() %}
            Send(*{%s m.GetInputType() %}) error
        {% endif %}

        {% if m.GetServerStreaming() %}
            Recv(*{%s m.GetOutputType() %}) error
        {% endif %}
    }

    type {%s srv+m.GetName() %} struct {
        stream client.Stream
    }


    func (x *{%s srv+m.GetName() %}) Close() error {
    	return x.stream.Close()
    }

    func (x *{%s srv+m.GetName() %}) Context() context.Context {
    	return x.stream.Context()
    }

    func (x *{%s srv+m.GetName() %}) SendMsg(m interface{}) error {
    	return x.stream.Send(m)
    }

    func (x *{%s srv+m.GetName() %}) RecvMsg(m interface{}) error {
    	return x.stream.Recv(m)
    }

    func (x *{%s srv+m.GetName() %}) Send(m *Message) error {
    	return x.stream.Send(m)
    }


    {% if m.GetClientStreaming() %}
        func (x *{%s srv+m.GetName() %}) Send(m *{%s m.GetInputType() %}) error{
            return x.stream.Send(m)
        }
    {% endif %}

    {% if m.GetServerStreaming() %}
        func (x *{%s srv+m.GetName() %}) Recv() (*{%s m.GetOutputType() %}, error) {
            m := new({%s m.GetOutputType() %})
            err := x.stream.Recv(m)
            if err != nil {
                return nil, err
            }
                return m, nil
        }
    {% endif %}


{% endfunc %}

{% stripspace %}
{%- func clientMethod(srv string,m *gen.Method) -%}
    {%- code
        reqArg := ", in *" + m.GetInputType()
        if m.GetClientStreaming() {
                reqArg = ""
        }

        respName := "*" + m.GetOutputType()
        if m.GetServerStreaming() || m.GetClientStreaming() {
            respName = srv + "_" + m.GetName() + "Service"
        }
    -%}
    {%s m.GetName() %}(ctx context.Context{%s reqArg %}, opts ...client.CallOption) ({%-s respName -%}, error)
{%- endfunc -%}
{% endstripspace %}


{% stripspace %}
{%- func serverMethod(srv string,m *gen.Method) -%}
    {%- code
        var reqArgs []string
        if !m.GetClientStreaming() {
            reqArgs = append(reqArgs, "*"+m.GetInputType())
        }
        if m.GetServerStreaming() || m.GetClientStreaming() {
            reqArgs = append(reqArgs, srv+"_"+m.GetName()+"Stream")
        }
        if !m.GetClientStreaming() && !m.GetServerStreaming() {
            reqArgs = append(reqArgs, "*"+m.GetOutputType())
        }
    -%}
    {%s m.GetName() %}(context.Context, {%s strings.Join(reqArgs, ", ") %}) error
{%- endfunc -%}
{% endstripspace %}


{%- func serviceServer(srv string,m *gen.Method) -%}
    {%- code
        methName := m.GetName()
        serveType := srv + "Handler"
        inType := m.GetInputType()
        outType := m.GetOutputType()
        unexport:=gen.UnExport
        streamType := unexport(srv) + methName + "Stream"
    -%}

    {% if !m.GetServerStreaming() && !m.GetClientStreaming() %}
        func (h *{%s unexport(srv)%}Handler){%s methName %}(ctx context.Context, in *{%s inType%}, out *{%s outType %}) error {
    		return h.{%s serveType %}.{%s methName %}(ctx, in, out)
        }
        {% return %}
    {% endif %}

    func (h *{%s unexport(srv)%}Handler){%s methName %}(ctx context.Context, stream server.Stream) error {
        {% if !m.GetClientStreaming() %}
            m := new({%s inType %})
            if err := stream.Recv(m); err != nil { return err }
            return h.{%s serveType %}.{%s methName %}(ctx, m, &{%s streamType %}{stream})
        {% else %}
            return h.{%s serveType %}.{%s methName %}(ctx, &{%s streamType %}{stream})
        {% endif %}
    }


    type {%s srv+"_"+methName %}Stream interface {
        Context() context.Context
        SendMsg(interface{}) error
        RecvMsg(interface{}) error
        Close() error
        {% if m.GetServerStreaming() %}
            Send(*{%s outType %}) error
        {% endif %}

        {% if m.GetClientStreaming() %}
            Recv() (*{%s inType %}, error)
        {% endif %}
    }

    type {%s streamType %} struct {
        stream server.Stream
    }


    func (x *{%s streamType %}) Close() error {
    	return x.stream.Close()
    }

    func (x *{%s streamType %}) Context() context.Context {
    	return x.stream.Context()
    }

    func (x *{%s streamType %}) SendMsg(m interface{}) error {
    	return x.stream.Send(m)
    }

    func (x *{%s streamType %}) RecvMsg(m interface{}) error {
    	return x.stream.Recv(m)
    }

    {% if m.GetServerStreaming() %}
        func (x *{%s streamType %}) Send(*{%s outType %}) error {
            return x.stream.Send(m)
        }
    {% endif %}

    {% if m.GetClientStreaming() %}
        func (x *{%s streamType %}) Recv() (*{%s inType %}, error) {
            m := new({%s inType %})
            if err := x.stream.Recv(m); err != nil { return nil, err }
            return m, nil
        }
    {% endif %}

{%- endfunc -%}

{% func Service(ss *gen.Service) %}
    {% code
        service:=ss.GetName()
        service1:=gen.UnExport(ss.GetName())
    %}

    // Client API for @service service
    type {%s service %}Service interface {
        {%- for _, m := range ss.GetMethod() -%}
            {%s clientMethod(service,m) %}
        {%- endfor -%}
    }

    type {%s service1 %}Service struct {
        c    client.Client
        name string
    }

    func New{%s service %}Service(name string, c client.Client) {%s service %}Service {
        return &{%s service1 %}Service{
            c:    c,
            name: name,
        }
    }

    {%- for _, m := range ss.GetMethod() -%}
        {%= serviceClient(service,m) %}
    {%- endfor -%}

    // Server API for {%s service %} service
    type {%s service %}Handler interface {
        {%- for _, m := range ss.GetMethod() -%}
            {%= serverMethod(service,m) %}
        {%- endfor -%}
    }

    func Register{%s service %}Handler(s server.Server, hdlr {%s service %}Handler, opts ...server.HandlerOption) error {
        type {%s service1 %} interface {
            {%- for _, m := range ss.GetMethod() -%}
                {% if !m.GetServerStreaming() && !m.GetClientStreaming() %}
                    {%s m.GetName() %}(ctx context.Context,in *{%s m.GetInputType() %},out *{%s m.GetOutputType() %}) error
                    {% continue %}
                {% endif %}
                {%s m.GetName() %}(ctx context.Context,stream server.Stream)error
            {%- endfor -%}
        }
        type {%s service %} struct {
            {%s service1 %}
        }
        h := &{%s service1 %}Handler{hdlr}
        return s.Handle(s.NewHandler(&{%s service %}{h}, opts...))
    }

    type {%s service1 %}Handler struct {
        {%s service %}Handler
    }

    {% for _, m := range ss.GetMethod() %}
        {%s= serviceServer(service,m) %}
    {% endfor %}


    {% for _, m := range ss.GetMethod() %}
        {% code
            method , path := m.GetHttpMethod()
        %}
        // {%s method %}
        // {%s path %}
    {% endfor %}
{% endfunc %}