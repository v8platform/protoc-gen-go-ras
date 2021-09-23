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
		m.g.P(method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ", opts... interface{}) (*", method.Output.GoIdent, ", error)")
	}
	m.g.P()
	m.g.P("}")

	m.AddImpl(m.getServiceImpl(service), protogen.GoImportPath(m.file.GoPackageName))

}

func (m messageServiceGenerator) genConstructor(service *protogen.Service) {

	serviceName := m.getServiceName(service)

	m.g.P("func New", serviceName, "(client ", getClientImp(), ") ", m.getServiceImpl(service), "{")
	m.g.P("return &", unexport(serviceName), "{")
	m.g.P("client,")
	m.g.P("}")
	m.g.P("}")

}

func (m messageServiceGenerator) genDefinition(service *protogen.Service) {
	serviceName := m.getServiceName(service)

	m.g.P("// ", serviceName, " is the endpoint message service for RAS service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.Annotate(serviceName, service.Location)
	m.g.P("type ", unexport(serviceName), " struct {")
	m.g.P("cc ", getClientImp(), "")
	m.g.P("}")
	m.g.P()
}

func (m messageServiceGenerator) genMethodHandler(service *protogen.Service, method *protogen.Method) {

	m.g.P("func (x *", unexport(m.getServiceName(service)), ") ", method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ", opts... interface{}) (*", method.Output.GoIdent, ", error) {")
	m.g.P()
	m.g.P("endpoint, err := x.cc.GetEndpoint(ctx)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P("return ", m.getMethodHandlerName(method), "(ctx, x.cc.Request, endpoint, req, opts...)")
	m.g.P("}")
	m.g.P()
	m.g.P("func ", m.getMethodHandlerName(method), "(ctx ", ctxPackage.Ident("Context"), ", cc Request, endpoint Endpoint, req *", method.Input.GoIdent, ", opts... interface{}) (*", method.Output.GoIdent, ", error) {")
	m.g.P()
	m.g.P("resp := new(", method.Output.GoIdent, ")")
	if isEmptyPb(method.Output.Desc) {
		m.g.P("if err := cc(ctx, ", m.ObjectNamed(".ras.protocol.v1.EndpointOpen").GoImportPath.Ident("EndpointRequestHandler"), "(endpoint, req, nil), opts...); err != nil {")
	} else {
		m.g.P("if err := cc(ctx, ", m.ObjectNamed(".ras.protocol.v1.EndpointOpen").GoImportPath.Ident("EndpointRequestHandler"), "(endpoint, req, resp), opts...); err != nil {")
	}
	m.g.P("return nil, err")
	m.g.P("}")
	m.g.P("return resp, nil")
	m.g.P("}")
	m.g.P()

}

func (m messageServiceGenerator) getServiceName(service *protogen.Service) string {
	return service.GoName
}

func (m messageServiceGenerator) getServiceImpl(service *protogen.Service) string {
	return service.GoName
}

func (m messageServiceGenerator) getClientOptionsName(service *protogen.Service) string {
	return service.GoName + "Options"
}

func (m messageServiceGenerator) getClientOptionName(service *protogen.Service) string {
	return service.GoName + "Option"
}

func (m messageServiceGenerator) getMethodHandlerName(method *protogen.Method) interface{} {
	return method.GoName + "Handler"
}
