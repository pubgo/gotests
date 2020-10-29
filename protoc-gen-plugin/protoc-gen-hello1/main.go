package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/pubgo/xerror"
	plugin "google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	contextPkgPath = "context"
	clientPkgPath  = "github.com/asim/go-micro/v3/client"
	serverPkgPath  = "github.com/asim/go-micro/v3/server"
)

const tmplService = `
package {{.Package}}

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
	defer xerror.Resp(func(err xerror.XErr) {
		log.Println(err.Println())
	})

	var request plugin.CodeGeneratorRequest
	var response plugin.CodeGeneratorResponse

	data := xerror.PanicBytes(ioutil.ReadAll(os.Stdin))
	xerror.Panic(proto.Unmarshal(data, &request))

	if len(request.GetFileToGenerate()) == 0 {
		return
	}

	parameter := request.GetParameter()
	param := make(map[string]string)
	for _, p := range strings.Split(parameter, ",") {
		if i := strings.Index(p, "="); i < 0 {
			param[p] = ""
		} else {
			param[p[0:i]] = p[i+1:]
		}
	}

	for _, param := range strings.Split(request.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}
		switch param {
		case "":
			// Ignore.
		case "plugins":
			for _, plugin := range strings.Split(value, ";") {
				switch plugin {
				case "prometheus":
					pluginPrometheus = true
				default:
					log.Fatalf("invalid plugin: %s", plugin)
				}
			}
		case "paths":
			if value == "source_relative" {
				pathsSourceRelative = true
			} else if value == "import" {
				pathsSourceRelative = false
			} else {
				log.Fatalf(`unknown path type %q: want "import" or "source_relative"`, value)
			}
		}
	}

	//importMap := make(map[string]string)
	//for k, v := range param {
	//	switch k {
	//	case "import_prefix":
	//		importPrefix = v
	//	case "import_path":
	//		packageImportPath = v
	//	case "paths":
	//		switch v {
	//		case "import":
	//			pathType = pathTypeImport
	//		case "source_relative":
	//			pathType = pathTypeSourceRelative
	//		default:
	//			log.Fatal(fmt.Sprintf(`Unknown path type %q: want "import" or "source_relative".`, v))
	//		}
	//	case "plugins":
	//		pluginList = v
	//	default:
	//		if len(k) > 0 && k[0] == 'M' {
	//			importMap[k[1:]] = v
	//		}
	//	}
	//}

	for _, name := range request.GetFileToGenerate() {
		var fd *descriptor.FileDescriptorProto
		for _, fd = range request.GetProtoFile() {
			if name == fd.GetName() {
				break
			}
		}

		if fd == nil {
			xerror.Panic(xerror.Fmt("could not find the .proto file for %s", name))
			return
		}

		pkg, ok := goPackageName(fd)
		if !ok {
			log.Println("请设置package")
		}

		g := jen.NewFile(pkg)
		g.ImportAlias("net/rpc", "rpc1")

		for _, ss := range fd.GetService() {
			//spec := &ServiceSpec{
			//	Package:     fd.GetPackage(),
			//	ServiceName: generator.CamelCase(ss.GetName()),
			//}

			// service
			srvName := generator.CamelCase(ss.GetName())
			g.Comment("// test")
			g.Comment("/* ssss */")
			g.Type().Id(srvName).Id("interface").BlockFunc(func(group *jen.Group) {
				for _, m := range ss.GetMethod() {
					mthName := generator.CamelCase(m.GetName())
					group.Id(mthName).Params(
						jen.Id("in *"+getTypeName(pkg, m.GetInputType())),
						jen.Id("out *"+getTypeName(pkg, m.GetOutputType())),
					).Error()
				}
			})

			// method
			//srv *rpc.Server, x TestApiInterface
			g.Func().Id("Register"+srvName).Params(
				jen.Id("srv *").Qual("net/rpc", "Server"),
				jen.Id("x "+srvName),
			).Error().BlockFunc(func(group *jen.Group) {
				group.Return().Nil()
			})

			//for _, m := range ss.GetMethod() {
			//	spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			//		MethodName:     generator.CamelCase(m.GetName()),
			//		InputTypeName:  m.GetInputType(),
			//		OutputTypeName: m.GetOutputType(),
			//	})
			//}

			//origMethName := m.GetName()
			//methName := generator.CamelCase(origMethName)
			//reqArg := ", in *" + g.typeName(m.GetInputType())
			//if m.GetClientStreaming() {
			//	reqArg = ""
			//}
			//respName := "*" + g.typeName(m.GetOutputType())
			//if m.GetServerStreaming() || m.GetClientStreaming() {
			//	respName = servName + "_" + generator.CamelCase(origMethName) + "Service"
			//}

			//var (
			//	contextPkg = generator.RegisterUniquePackageName("context", nil)
			//	clientPkg  = generator.RegisterUniquePackageName("client", nil)
			//	serverPkg  = generator.RegisterUniquePackageName("server", nil)
			//)
			//
			//fmt.Printf("%s(ctx %s.Context%s, opts ...%s.CallOption) (%s, error)", methName, contextPkg, reqArg, clientPkg, respName)

			//var buf bytes.Buffer
			//t := template.Must(template.New("").Parse(tmplService))
			//xerror.Panic(t.Execute(&buf, spec))
		}

		name := fd.GetName()
		if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
			name = name[:len(name)-len(ext)]
		}

		response.File = append(response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name + ".pb.hello.go"),
			Content: proto.String(g.GoString()),
		})
	}

	data = xerror.PanicBytes(proto.Marshal(&response))
	xerror.ExitErr(os.Stdout.Write(data))
}

// baseName returns the last path element of the name, with the last dotted suffix removed.
func baseName(name string) string {
	// First, find the last element
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	// Now drop the suffix
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[0:i]
	}
	return name
}

// getGoPackage returns the file's go_package option.
// If it containts a semicolon, only the part before it is returned.
func getGoPackage(fd *descriptor.FileDescriptorProto) string {
	pkg := fd.GetOptions().GetGoPackage()
	if strings.Contains(pkg, ";") {
		parts := strings.Split(pkg, ";")
		if len(parts) > 2 {
			log.Fatalf(
				"protoc-gen-nrpc: go_package '%s' contains more than 1 ';'",
				pkg)
		}
		pkg = parts[1]
	}

	return pkg
}

// goPackageOption interprets the file's go_package option.
// If there is no go_package, it returns ("", "", false).
// If there's a simple name, it returns ("", pkg, true).
// If the option implies an import path, it returns (impPath, pkg, true).
func goPackageOption(d *descriptor.FileDescriptorProto) (impPath, pkg string, ok bool) {
	pkg = getGoPackage(d)
	if pkg == "" {
		return
	}
	ok = true
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(pkg, "/")
	if slash < 0 {
		return
	}
	impPath, pkg = pkg, pkg[slash+1:]
	// A semicolon-delimited suffix overrides the package name.
	sc := strings.IndexByte(impPath, ';')
	if sc < 0 {
		return
	}
	impPath, pkg = impPath[:sc], impPath[sc+1:]
	return
}

// goPackageName returns the Go package name to use in the
// generated Go file.  The result explicit reports whether the name
// came from an option go_package statement.  If explicit is false,
// the name was derived from the protocol buffer's package statement
// or the input file name.
func goPackageName(d *descriptor.FileDescriptorProto) (name string, explicit bool) {
	// Does the file have a "go_package" option?
	if _, pkg, ok := goPackageOption(d); ok {
		return pkg, true
	}

	// Does the file have a package clause?
	if pkg := d.GetPackage(); pkg != "" {
		return pkg, false
	}
	// Use the file base name.
	return baseName(d.GetName()), false
}

// goFileName returns the output name for the generated Go file.
func goFileName(d *descriptor.FileDescriptorProto) string {
	name := *d.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += ".nrpc.go"

	if pathsSourceRelative {
		return name
	}

	// Does the file have a "go_package" option?
	// If it does, it may override the filename.
	if impPath, _, ok := goPackageOption(d); ok && impPath != "" {
		// Replace the existing dirname with the declared import path.
		_, name = path.Split(name)
		name = path.Join(impPath, name)
		return name
	}

	return name
}

var pluginPrometheus bool
var pathsSourceRelative bool

// splitMessageTypeName split a message type into (package name, type name)
func splitMessageTypeName(name string) (string, string) {
	if len(name) == 0 {
		log.Fatal("Empty message type")
	}
	if name[0] != '.' {
		log.Fatalf("Expect message type name to start with '.', but it is '%s'", name)
	}
	lastDot := strings.LastIndex(name, ".")
	return name[1:lastDot], name[lastDot+1:]
}

func splitTypePath(name string) []string {
	if len(name) == 0 {
		log.Fatal("Empty message type")
	}
	if name[0] != '.' {
		log.Fatalf("Expect message type name to start with '.', but it is '%s'", name)
	}
	return strings.Split(name[1:], ".")
}

//func pkgSubject(fd *descriptor.FileDescriptorProto) string {
//	if options := fd.GetOptions(); options != nil {
//		e := proto.GetExtension(options, nrpc.E_PackageSubject)
//		if value, ok := e.(string); ok {
//			return value
//		}
//	}
//	return fd.GetPackage()
//}

func getTypeName(pkgName, mthName string) string {
	mthName = strings.Trim(mthName, ".")
	mthName = strings.TrimLeft(mthName, pkgName)
	mthName = strings.Trim(mthName, ".")
	return mthName
}

//func getGoType(pbType string) (string, string) {
//	if !strings.Contains(pbType, ".") {
//		return "", pbType
//	}
//	fd, _ := lookupMessageType(pbType)
//	name := strings.TrimPrefix(pbType, "."+fd.GetPackage()+".")
//	name = strings.Replace(name, ".", "_", -1)
//	return getGoPackage(fd), name
//}

//func getPkgImportName(goPkg string) string {
//	if goPkg == getGoPackage(currentFile) {
//		return ""
//	}
//	replacer := strings.NewReplacer(".", "_", "/", "_", "-", "_")
//	return replacer.Replace(goPkg)
//}

func GetPkg(pkg, s string) string {
	s = strings.TrimPrefix(s, ".")
	s = strings.TrimPrefix(s, pkg)
	s = strings.TrimPrefix(s, ".")
	return s
}

//func GetExtraImports(fd *descriptor.FileDescriptorProto) []string {
//	// check all the types used and imports packages from where they come
//	var imports = make(map[string]string)
//	for _, sd := range fd.GetService() {
//		for _, md := range sd.GetMethod() {
//			goPkg, _ := getGoType(md.GetInputType())
//			pkgImportName := getPkgImportName(goPkg)
//			if pkgImportName != "" {
//				imports[pkgImportName] = goPkg
//			}
//			goPkg, _ = getGoType(getResultType(md))
//			pkgImportName = getPkgImportName(goPkg)
//			if pkgImportName != "" {
//				imports[pkgImportName] = goPkg
//			}
//		}
//	}
//	var result []string
//	for importName, goPkg := range imports {
//		result = append(result,
//			fmt.Sprintf("%s \"%s\"",
//				importName,
//				goPkg,
//			),
//		)
//	}
//	return result
//}
//
//func GetPkgSubjectParams(fd *descriptor.FileDescriptorProto) []string {
//	if opts := fd.GetOptions(); opts != nil {
//		e := proto.GetExtension(opts, nrpc.E_PackageSubjectParams)
//		value := e.([]string)
//		return value
//	}
//	return nil
//}

//func GetServiceSubject(sd *descriptor.ServiceDescriptorProto) string {
//	if opts := sd.GetOptions(); opts != nil {
//		s := proto.GetExtension(opts, nrpc.E_ServiceSubject)
//		if value, ok := s.(string); ok && s != "" {
//			return value
//		}
//	}
//	if opts := currentFile.GetOptions(); opts != nil {
//		s := proto.GetExtension(opts, nrpc.E_ServiceSubjectRule)
//		switch s.(nrpc.SubjectRule) {
//		case nrpc.SubjectRule_COPY:
//			return sd.GetName()
//		case nrpc.SubjectRule_TOLOWER:
//			return strings.ToLower(sd.GetName())
//		}
//	}
//	return sd.GetName()
//}

func fileIsProto3(file *descriptor.FileDescriptorProto) bool {
	return file.GetSyntax() == "proto3"
}

// A GoImportPath is the import path of a Go package. e.g., "google.golang.org/genproto/protobuf".
type GoImportPath string

func (p GoImportPath) String() string { return strconv.Quote(string(p)) }

// And now lots of helper functions.

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// CamelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// CamelCaseSlice is like CamelCase, but the argument is a slice of strings to
// be joined with "_".
func CamelCaseSlice(elem []string) string { return CamelCase(strings.Join(elem, "_")) }

// dottedSlice turns a sliced name into a dotted name.
func dottedSlice(elem []string) string { return strings.Join(elem, ".") }

// Is this field optional?
func isOptional(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_OPTIONAL
}

// Is this field required?
func isRequired(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REQUIRED
}

// Is this field repeated?
func isRepeated(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

// Is this field a scalar numeric type?
func isScalar(field *descriptor.FieldDescriptorProto) bool {
	if field.Type == nil {
		return false
	}
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_ENUM,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
}

// baseName returns the last path element of the name, with the last dotted suffix removed.
func baseName(name string) string {
	// First, find the last element
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	// Now drop the suffix
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[0:i]
	}
	return name
}

// badToUnderscore is the mapping function used to generate Go names from package names,
// which can be dotted in the input .proto file.  It replaces non-identifier characters such as
// dot or dash with underscore.
func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}

func goFileName1(pathType pathType) string {
	name := *d.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += ".pb.go"

	if pathType == pathTypeSourceRelative {
		return name
	}

	// Does the file have a "go_package" option?
	// If it does, it may override the filename.
	if impPath, _, ok := d.goPackageOption(); ok && impPath != "" {
		// Replace the existing dirname with the declared import path.
		_, name = path.Split(name)
		name = path.Join(string(impPath), name)
		return name
	}

	return name
}

func goPackageOption1(fd *descriptor.FileDescriptorProto) (impPath GoImportPath, pkg GoPackageName, ok bool) {
	opt := fd.GetOptions().GetGoPackage()
	if opt == "" {
		return "", "", false
	}
	// A semicolon-delimited suffix delimits the import path and package name.
	sc := strings.Index(opt, ";")
	if sc >= 0 {
		return GoImportPath(opt[:sc]), cleanPackageName(opt[sc+1:]), true
	}
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(opt, "/")
	if slash >= 0 {
		return GoImportPath(opt), cleanPackageName(opt[slash+1:]), true
	}
	return "", cleanPackageName(opt), true
}

type GoPackageName string

func cleanPackageName(name string) GoPackageName {
	name = strings.Map(badToUnderscore, name)
	// Identifier must not be keyword or predeclared identifier: insert _.
	if isGoKeyword[name] {
		name = "_" + name
	}
	// Identifier must not begin with digit: insert _.
	if r, _ := utf8.DecodeRuneInString(name); unicode.IsDigit(r) {
		name = "_" + name
	}
	return GoPackageName(name)
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

var isGoPredeclaredIdentifier = map[string]bool{
	"append":     true,
	"bool":       true,
	"byte":       true,
	"cap":        true,
	"close":      true,
	"complex":    true,
	"complex128": true,
	"complex64":  true,
	"copy":       true,
	"delete":     true,
	"error":      true,
	"false":      true,
	"float32":    true,
	"float64":    true,
	"imag":       true,
	"int":        true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"int8":       true,
	"iota":       true,
	"len":        true,
	"make":       true,
	"new":        true,
	"nil":        true,
	"panic":      true,
	"print":      true,
	"println":    true,
	"real":       true,
	"recover":    true,
	"rune":       true,
	"string":     true,
	"true":       true,
	"uint":       true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uint8":      true,
	"uintptr":    true,
}
