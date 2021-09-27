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

	protocolv1 := m.ObjectNamed(".ras.protocol.v1.EndpointOpen").GoImportPath

	m.g.P("func (x *", unexport(m.getServiceName(service)), ") ", method.GoName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ", opts... interface{}) (*", method.Output.GoIdent, ", error) {")
	m.g.P()
	m.g.P("reply, err := x.cc.Invoke(ctx, true, req, ", m.getMethodHandlerName(method), ", opts...)")
	m.g.P("if err != nil { return nil, err }")
	m.g.P("return reply.(*", method.Output.GoIdent, "), nil")
	m.g.P("}")
	m.g.P()
	m.g.P("func ", m.getMethodHandlerName(method), "(ctx ", ctxPackage.Ident("Context"), ", channel Channel, endpoint Endpoint, req interface{}, interceptor Interceptor) (interface{}, error) {")
	m.g.P()
	m.g.P("if interceptor == nil {")
	m.g.P("reply := new(", method.Output.GoIdent, ")")
	if isEmptyPb(method.Output.Desc) {
		m.g.P("return reply, ", protocolv1.Ident("EndpointChannelRequest"), "(ctx, channel, endpoint, req.(*", method.Input.GoIdent, "), nil)")
	} else {
		m.g.P("return reply, ", protocolv1.Ident("EndpointChannelRequest"), "(ctx, channel, endpoint, req.(*", method.Input.GoIdent, "), reply)")
	}
	m.g.P("}")

	m.g.P("info := &RequestInfo {")
	m.g.P("Method : \"", method.GoName, "\",")
	m.g.P("FullMethod : \"/", method.Desc.Parent().FullName(), "/", method.GoName, "\",")
	m.g.P("}")
	m.g.P("")
	m.g.P("handler := func (ctx ", ctxPackage.Ident("Context"), ", cc Channel, endpoint Endpoint, req interface{}) (interface{}, error) {")
	m.g.P("reply := new(", method.Output.GoIdent, ")")
	if isEmptyPb(method.Output.Desc) {
		m.g.P("return reply, ", protocolv1.Ident("EndpointChannelRequest"), "(ctx, cc, endpoint, req.(*", method.Input.GoIdent, "), nil)")
	} else {
		m.g.P("return reply, ", protocolv1.Ident("EndpointChannelRequest"), "(ctx, cc, endpoint, req.(*", method.Input.GoIdent, "), reply)")
	}
	m.g.P("}")
	m.g.P("return interceptor(ctx, channel, endpoint, info, req, handler)")
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
