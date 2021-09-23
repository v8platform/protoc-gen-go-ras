package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

type proxyMethod struct {
	*protogen.Method
	proxyName string
}

type rasServiceGenerator struct {
	*Generator
	gen  *protogen.Plugin
	file *protogen.File
	g    *protogen.GeneratedFile

	idxRequests []*protogen.Service
	idxMethods  []proxyMethod
}

// GenerateFileContent generates the gRPC service definitions, excluding the package statement.
func (m rasServiceGenerator) GenerateFileContent() {
	if len(m.file.Services) == 0 {
		return
	}

	for _, service := range m.file.Services {
		if GetIsRASServiceExtension(service.Desc.Options()) {
			m.genService(service)
		}
	}
}

func (m *rasServiceGenerator) init() {

	for _, file := range m.gen.Files {
		if len(file.Services) == 0 {
			continue
		}

		for _, service := range file.Services {
			if GetIsRequestServiceExtension(service.Desc.Options()) {
				m.idxRequests = append(m.idxRequests, service)
			}
		}
	}

	for _, request := range m.idxRequests {
		for _, method := range request.Methods {
			ext := GetClientMethodExtension(method.Desc.Options())
			name := method.GoName
			if len(ext.ProxyName) > 0 {
				name = ext.ProxyName
			}

			m.idxMethods = append(m.idxMethods, proxyMethod{
				Method:    method,
				proxyName: name,
			})

		}
	}

}

func (m rasServiceGenerator) getServiceName(service *protogen.Service) string {
	return service.GoName
}

func (m rasServiceGenerator) getServiceImpl(service *protogen.Service) string {
	return service.GoName + "Impl"
}

func (m rasServiceGenerator) genService(service *protogen.Service) {
	//
	// m.g.Unskip()
	//
	// m.genImpl(service)
	// m.genConstructor(service)
	// m.genDefinition(service)
	//
	// for _, method := range m.idxMethods {
	// 	m.genProxyMethod(service, method)
	// }

}

func (m rasServiceGenerator) genProxyMethod(service *protogen.Service, method proxyMethod) {

	serviceName := unexport(m.getServiceName(service))

	m.g.P("// ", method.proxyName, " is proxy method to request ", m.getServiceName(method.Parent), ".", method.GoName)
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.P("func (x *", serviceName, ") ", method.proxyName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
	m.g.P("return x.", m.getServiceName(method.Parent), ".", method.GoName, "(ctx, req)")
	m.g.P("}")
	m.g.P()
}

func (m rasServiceGenerator) genImpl(service *protogen.Service) {

	m.g.P("type ", m.getServiceImpl(service), " interface {")
	for _, method := range m.idxMethods {
		m.g.P("// ", method.proxyName, " proxy request ", m.getServiceName(method.Parent), ".", method.GoName)
		m.g.P(method.proxyName, "(ctx ", ctxPackage.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error)")
	}
	m.g.P()
	m.g.P("}")

	m.AddImpl(m.getServiceImpl(service), protogen.GoImportPath(m.file.GoPackageName))

}

func (m rasServiceGenerator) genConstructor(service *protogen.Service) {

	endpointImpl := "EndpointServiceImpl"
	serviceName := unexport(m.getServiceName(service))

	m.g.P("func New", m.getServiceName(service), "(endpointService ", endpointImpl, ") ", m.getServiceImpl(service), "{")
	m.g.P("return &", serviceName, "{")
	for _, service := range m.idxRequests {
		m.g.P("// proxy service ", m.getServiceName(service))
		m.g.P(m.getServiceName(service), ":", "New", m.getServiceName(service), "(endpointService),")
		m.g.P()
	}
	m.g.P("}")
	m.g.P("}")

}

func (m rasServiceGenerator) genDefinition(service *protogen.Service) {
	serviceName := unexport(m.getServiceName(service))

	m.g.P("// ", serviceName, " is the endpoint message service for RAS service.")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		m.g.P("//")
		m.g.P(deprecationComment)
	}
	m.g.Annotate(serviceName, service.Location)
	m.g.P("type ", serviceName, " struct {")
	for _, service := range m.idxRequests {
		m.g.P("// Request service ", m.getServiceName(service))
		m.g.P(m.getServiceName(service), " ", m.getServiceImpl(service))
		m.g.P()
	}
	m.g.P("}")
	m.g.P()
}
