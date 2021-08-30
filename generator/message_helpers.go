package generator

import "google.golang.org/protobuf/compiler/protogen"

func (gen *Generator) generateMessageHelpers(g *protogen.GeneratedFile, m *protogen.Message) {

	// endpointMessageDataTypeIdent := gen.EnumNamed("ENDPOINT_DATA_TYPE_MESSAGE").GoIdent
	// messageDataIdent := gen.enumToObject["ENDPOINT_DATA_TYPE_MESSAGE"].GoIdent
	// _ = gen.enumToObject["ENDPOINT_DATA_TYPE_VOID_MESSAGE"].GoIdent
	// _ = gen.enumToObject["ENDPOINT_DATA_TYPE_EXCEPTION"].GoIdent
	//
	// messageDataEnumIdent := gen.EnumNamed("ENDPOINT_DATA_TYPE_MESSAGE").GoIdent
	// voidMessageEnumIdent := gen.EnumNamed("ENDPOINT_DATA_TYPE_VOID_MESSAGE").GoIdent
	// exceptionMessagEnumeIdent := gen.EnumNamed("ENDPOINT_DATA_TYPE_EXCEPTION").GoIdent
	//
	// MessageTypeIdent := gen.EnumNamed("GET_AGENT_ADMINS_REQUEST").Parent.GoIdent
	//
	// messageName := m.GoIdent.GoName
	// messageFormatterName := messageName + "Formatter"
	// messageParserName := messageName + "Parser"
	//
	// fields := getMessageFields(m)
	// dataField := fields.FindByNumber(5)
	//
	// g.Annotate(m.GoIdent.GoName, m.Location)
	// g.P("func (x *", m.GoIdent, ") Unpack(into ", messageParserName, ") error {")
	// g.P("switch x.GetType() {" )
	// g.P("case ",messageDataEnumIdent,":")
	// g.P("return x.GetMessage().Unpack(into)")
	// g.P("case ",voidMessageEnumIdent,":")
	// g.P("return nil ")
	// g.P("case ",exceptionMessagEnumeIdent,":")
	// g.P("return x.GetFailure()")
	// g.P("default:")
	// g.P("return ", fmtErrorf, "(\"unknown message type <%s>\", x.GetType())")
	// g.P("}")
	// g.P("}")
	// g.P()
	//
	// g.Unskip()
}
