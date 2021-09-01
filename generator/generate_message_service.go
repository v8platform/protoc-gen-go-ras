package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

type messageServiceGenerator struct {
	*Generator
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func (m messageServiceGenerator) GenerateFileContent() {
	if len(m.file.Services) == 0 {
		return
	}

	for _, service := range m.file.Services {
		if GetIsRequestServiceExtension(service.Desc.Options()) {
			m.genService(service)
		}
	}
}

func (m messageServiceGenerator) genService(service *protogen.Service) {

	m.g.Unskip()

	m.genImpl(service)
	m.genConstructor(service)
	m.genDefinition(service)

	for _, method := range service.Methods {
		m.genMethodHandler(service, method)
	}
}

func (m messageServiceGenerator) genImpl(service *protogen.Service) {

	m.g.P("type ", m.getServiceImpl(service), " interface {")
	for _, method := range service.Methods {
		m.g.P(method.GoName, "(*", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error)")
	}
	m.g.P()
	m.g.P("}")

	m.AddImpl(m.getServiceImpl(service), protogen.GoImportPath(m.file.GoPackageName))

}

func (m messageServiceGenerator) genConstructor(service *protogen.Service) {

	endpointImpl := "EndpointServiceImpl"
	serviceName := m.getServiceName(service)

	m.g.P("func New", serviceName, "(endpointService ", endpointImpl, ") ", m.getServiceImpl(service), "{")
	m.g.P("return &", serviceName, "{")
	m.g.P("endpointService,")
	m.g.P("}")
	m.g.P("}")

}

func (m messageServiceGenerator) genDefinition(service *protogen.Service) {
	serviceName := m.getServiceName(service)
	endpointImpl := "EndpointServiceImpl"

	m.g.P("// ", serviceName, " is the endpoint message service for RAS service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.Annotate(serviceName, service.Location)
	m.g.P("type ", serviceName, " struct {")
	m.g.P("e ", endpointImpl, "")
	m.g.P("}")
	m.g.P()
}

func (m messageServiceGenerator) genMethodHandler(service *protogen.Service, method *protogen.Method) {

	endpointRequest := "EndpointRequest"

	m.g.P("func (x *", m.getServiceName(service), ") ", method.GoName, "(req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
	m.g.P()
	m.g.P("var resp ", method.Output.GoIdent)
	m.g.P()
	m.g.P("anyRequest, err := ", anypbPackage.Ident("New"), "(req)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("anyRespond, err := ", anypbPackage.Ident("New"), "(&resp)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("endpointRequest := &", endpointRequest, "{")
	m.g.P("Request: anyRequest,")
	m.g.P("Respond: anyRespond,")
	m.g.P("}")
	m.g.P("response, err := x.e.Request(endpointRequest)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P()
	m.g.P("if err := ", anypbPackage.Ident("UnmarshalTo"),
		"(response, &resp,", protoPackage.Ident("UnmarshalOptions"), "{}); err != nil {")
	m.g.P("return nil, err")
	m.g.P("}")
	m.g.P("return &resp, nil")
	m.g.P("}")
	m.g.P()

}

func (m messageServiceGenerator) getServiceName(service *protogen.Service) string {
	return service.GoName
}

func (m messageServiceGenerator) getServiceImpl(service *protogen.Service) string {
	return service.GoName + "Impl"
}

func (m messageServiceGenerator) getClientOptionsName(service *protogen.Service) string {
	return service.GoName + "Options"
}

func (m messageServiceGenerator) getClientOptionName(service *protogen.Service) string {
	return service.GoName + "Option"
}
