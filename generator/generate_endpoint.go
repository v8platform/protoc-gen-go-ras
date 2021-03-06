package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type endpointGenerator struct {
	*Generator
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func (m endpointGenerator) GenerateFileContent() {
	if len(m.file.Services) == 0 {
		return
	}

	for _, service := range m.file.Services {
		if GetIsEndpointExtension(service.Desc.Options()) {
			m.genService(service)
		}
	}
}

func (m endpointGenerator) genService(service *protogen.Service) {

	m.g.Unskip()

	m.genImpl(service)
	m.genConstructor(service)
	m.genDefinition(service)

	for _, method := range service.Methods {
		m.genMethodHandler(service, method)
	}
}

func (m endpointGenerator) genHeader(packageName string) {
	m.g.P("// Code generated by protoc-gen-go-ras. DO NOT EDIT.")
	m.g.P()
	m.g.P("package ", packageName)
	m.g.P()
}
func (m endpointGenerator) genImpl(service *protogen.Service) {

	m.g.P("type ", m.getEndpointImp(service), " interface {")
	for _, method := range service.Methods {
		m.g.P(method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error)")
	}

	m.g.P()
	m.g.P("}")

	m.AddImpl(m.getEndpointImp(service), m.file.GoImportPath)

}

func (m endpointGenerator) genConstructor(service *protogen.Service) {
	clientServiceImpl := m.GetImpl("ClientServiceImpl")
	endpointImpl := m.GetImpl("EndpointImpl")

	serviceName := m.getEndpointName(service)

	m.g.P("func New", serviceName, "(clientService ", clientServiceImpl, ", endpoint ", endpointImpl, ") ", m.getEndpointImp(service), "{")
	m.g.P("return &", unexport(serviceName), "{")
	m.g.P("endpoint,")
	m.g.P("clientService,")
	m.g.P("}")
	m.g.P("}")
	m.g.P()

}

func (m endpointGenerator) genDefinition(service *protogen.Service) {
	serviceName := m.getEndpointName(service)
	clientServiceImpl := m.GetImpl("ClientServiceImpl")
	endpointImpl := m.GetImpl("EndpointImpl")

	m.g.P("// ", serviceName, " is the endpoint service for RAS service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.Annotate(serviceName, service.Location)
	m.g.P("type ", unexport(serviceName), " struct {")
	m.g.P("", endpointImpl, "")
	m.g.P("client ", clientServiceImpl, "")
	m.g.P("}")
	m.g.P()
}

func isEmptyPb(m protoreflect.MessageDescriptor) bool {
	return m.FullName() == "google.protobuf.Empty"
}
func (m endpointGenerator) genMethodHandler(service *protogen.Service, method *protogen.Method) {

	endpointMessageParser := m.GetImpl("EndpointMessageParser")

	m.g.P("func (x *", unexport(m.getEndpointName(service)), ") ", method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
	m.g.P("message, err := ", anypbPackage.Ident("UnmarshalNew"),
		"(req.GetRequest(),", protoPackage.Ident("UnmarshalOptions"), "{})")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("reqMessage, err := x.NewMessage(message)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("respMessage, err := x.client.EndpointMessage(ctx, reqMessage) ")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("respProtoMessage, err := ", anypbPackage.Ident("UnmarshalNew"),
		"(req.GetRespond(),", protoPackage.Ident("UnmarshalOptions"), "{})")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("if _, ok := respProtoMessage.(*", emptypbPackage.Ident("Empty"), "); ok {")
	m.g.P("if err := x.UnpackMessage(respMessage, nil); err != nil { return nil, err }")
	m.g.P("return ", anypbPackage.Ident("New"), "(respProtoMessage)")
	m.g.P("}")
	m.g.P()
	m.g.P("messageParser, ok := respProtoMessage.(", endpointMessageParser, ")")
	m.g.P("if !ok { return nil, ", fmtErrorf, "(\"not parser interface\") }")
	m.g.P("if err := x.UnpackMessage(respMessage, messageParser); err != nil { return nil, err }")
	m.g.P("return ", anypbPackage.Ident("New"), "(respProtoMessage)")
	m.g.P("}")

}

func (m endpointGenerator) getEndpointName(service *protogen.Service) string {
	return service.GoName
}

func (m endpointGenerator) getEndpointImp(service *protogen.Service) string {
	return service.GoName + "Impl"
}
