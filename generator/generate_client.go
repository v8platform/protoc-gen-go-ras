package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

type clientGenerator struct {
	*Generator
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func (m clientGenerator) GenerateFileContent() {
	if len(m.file.Services) == 0 {
		return
	}

	for _, service := range m.file.Services {
		if GetIsClientExtension(service.Desc.Options()) {
			m.genService(service)
		}
	}
}

func (m clientGenerator) genService(service *protogen.Service) {

	m.g.Unskip()

	m.genClientImpl(service)
	m.genClientConstructor(service)
	m.genClientDefinition(service)
	m.genDetectSupportedVersion(service)

	for _, method := range service.Methods {
		m.genMethodHandler(service, method)
	}
}

func (m clientGenerator) genHeader(packageName string) {
	m.g.P("// Code generated by protoc-gen-go-ras. DO NOT EDIT.")
	m.g.P()
	m.g.P("package ", packageName)
	m.g.P()
}
func (m clientGenerator) genClientImpl(service *protogen.Service) {

	m.g.P("type ", m.getClientServiceImp(service), " interface {")
	for _, method := range service.Methods {
		m.g.P("", method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error)")
	}
	m.g.P("}")

	m.AddImpl(m.getClientServiceImp(service), m.file.GoImportPath)

	m.g.P("type ", m.getClientImp(), " interface {")
	m.g.P("// Методы для блокировки соединения sync.Mutex")
	m.g.P("// берем из sync.Locker ")
	m.g.P(syncPackage.Ident("Locker"))
	m.g.P("// Методы для записи и чтения из соединение")
	m.g.P("// берем из io.ReadWriter")
	m.g.P(ioPackage.Ident("ReadWriter"))
	m.g.P("}")

	m.AddImpl(m.getClientImp(), m.file.GoImportPath)

}

func (m clientGenerator) genClientConstructor(service *protogen.Service) {
	serviceName := m.getClientName(service)

	m.g.P("func New", serviceName, "(client ", m.getClientImp(), ") ", m.getClientServiceImp(service), "{")
	m.g.P("return &", unexport(serviceName), "{")
	m.g.P("client: client,")
	m.g.P("}")
	m.g.P("}")

}

func (m clientGenerator) genClientDefinition(service *protogen.Service) {
	serviceName := m.getClientName(service)

	m.g.P("// ", serviceName, " is the client for RAS service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.Annotate(serviceName, service.Location)
	m.g.P("type ", unexport(serviceName), " struct {")
	m.g.P("client ", m.getClientImp(), "")
	m.g.P("}")
	m.g.P()
}

func (m clientGenerator) genDetectSupportedVersion(service *protogen.Service) {
	serviceName := m.getClientName(service)

	m.g.Annotate(serviceName, service.Location)
	m.g.P("var serviceVersions = []string{\"3.0\", \"4.0\", \"5.0\", \"6.0\", \"7.0\", \"8.0\", \"9.0\", \"10.0\"}")
	m.g.P()
	m.g.P("var re = ", regexpPackage.Ident("MustCompile"), "(`(?m)supported=(.*?)]`)")
	m.g.P()
	m.g.P("// DetectSupportedVersion func helpers detect supported version in EndpointFailureAck")
	m.g.P("func DetectSupportedVersion(err error) string {")
	m.g.P()
	m.g.P("fail, ok := err.(*", m.ObjectNamed("ras.protocol.v1.EndpointFailureAck").GoIdent, ")")
	m.g.P("if !ok { return \"\" }")
	m.g.P()
	m.g.P("if fail.Cause == nil { return \"\" }")
	m.g.P()
	m.g.P("matches := re.FindAllString(fail.Cause.Message, -1)")
	m.g.P()
	m.g.P("if len(matches) == 0 { return \"\" }")
	m.g.P()
	m.g.P("supported := matches[0]")
	m.g.P("for i := len(serviceVersions) - 1; i >= 0; i-- {")
	m.g.P("version := serviceVersions[i]")
	m.g.P("if ", stringsPackage.Ident("Contains"), "(supported, version) { return version }")
	m.g.P("}")
	m.g.P()
	m.g.P("return \"\"")
	m.g.P("}")
	m.g.P()
}

func (m clientGenerator) genMethodHandler(service *protogen.Service, method *protogen.Method) {

	ext := GetClientMethodExtension(method.Desc.Options())
	if ext.NewEndpointFunc {
		m.genNewEndpointFunc(service, method)
		return
	}

	m.g.P("func (x *", unexport(m.getClientName(service)), ") ", method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
	m.g.P()
	m.g.P("x.client.Lock()")
	m.g.P("defer x.client.Unlock()")
	m.g.P()
	m.g.P("// Check context ")
	m.g.P("select {")
	m.g.P("case <-ctx.Done():")
	m.g.P("return nil, ctx.Err()")
	m.g.P("default:")
	m.g.P("}")
	m.g.P("")
	if ext.NoPacketPack {
		m.g.P("if err := req.", m.formatFuncName, "(x.client, 0 ); err != nil { return nil, err }")
	} else {
		m.g.P("packet, err := ", method.Input.GoIdent.GoImportPath.Ident("NewPacket"), "(req)")
		m.g.P("if err != nil { return nil, err }")
		m.g.P("if _, err := packet.WriteTo(x.client); err != nil { return nil, err }")
	}
	if isEmptyPb(method.Output.Desc) {
		m.g.P("return new(", method.Output.GoIdent, "), nil")
		m.g.P("}")
		return
	}
	m.g.P("ackPacket, err := ", method.Input.GoIdent.GoImportPath.Ident("NewPacket"), "(x.client)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P("resp := new(", method.Output.GoIdent, ")")
	m.g.P("return resp, ackPacket.Unpack(resp)")
	m.g.P("}")
	m.g.P()
}

func (m clientGenerator) genNewEndpointFunc(service *protogen.Service, method *protogen.Method) {

	m.g.P("func (x *", unexport(m.getClientName(service)), ") ", method.GoName, "(_ ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
	m.g.P("return &", method.Output.GoIdent, "{")
	m.g.P("Service: req.GetService(),")
	m.g.P("Version: ", castPackage.Ident("ToInt32"), "(", castPackage.Ident("ToFloat32"), "(req.GetVersion())),")
	m.g.P("Id: req.GetEndpointId(),")
	m.g.P("Format: ", codecVersion, "(),")
	m.g.P("}, nil")
	m.g.P("}")
	m.g.P()

}

func (m clientGenerator) getClientName(service *protogen.Service) string {
	return service.GoName
}

func (m clientGenerator) getClientImp() string {
	return "ClientImpl"
}

func (m clientGenerator) getClientServiceImp(service *protogen.Service) string {
	return service.GoName + "Impl"
}

func (m clientGenerator) getClientOptionsName(service *protogen.Service) string {
	return service.GoName + "Options"
}

func (m clientGenerator) getClientOptionName(service *protogen.Service) string {
	return service.GoName + "Option"
}

func (m clientGenerator) getRemoteMockClientName(service *protogen.Service) string {
	return service.GoName + "RemoteMockClient"
}

func (m clientGenerator) getMockServerBaseInterfaceName(service *protogen.Service) string {
	return service.GoName + "Server"
}

func (m clientGenerator) getMockServiceDescriptorName(service *protogen.Service) string {
	return "_" + service.GoName + "_MockServiceDesc"
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }
